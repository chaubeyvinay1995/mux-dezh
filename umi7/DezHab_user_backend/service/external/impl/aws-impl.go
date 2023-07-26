package impl

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"gitlab.com/umi7/DezHab_user_backend/common/logger"
	"gitlab.com/umi7/DezHab_user_backend/config"
	"gitlab.com/umi7/DezHab_user_backend/dto"
	"gitlab.com/umi7/DezHab_user_backend/service/external"
	"gitlab.com/umi7/DezHab_user_backend/utils/constants"
	"time"

	//"log"
	//"net/url"
	"os"
	"path/filepath"
)

var Aws external.Aws

type AWSClientImpl struct {
	AWSConfig dto.AWSConfig
}

func init() {
	Aws = AWSClientImpl{
		AWSConfig: dto.AWSConfig{
			ID:     config.AppConfig.Aws.ID,
			Token:  config.AppConfig.Aws.Token,
			Region: config.AppConfig.Aws.Region,
		},
	}
}

func (awsClient AWSClientImpl) PushFileToS3(ctx context.Context, bucket string, folder string,
	fileKey string, tagsData string, accessPermission string) (string, string, error) {
	//TODO : Move all the data store calls to separate repository layer.
	file, err := os.Open(fileKey)
	if err != nil {
		logger.Error(ctx, err)
		return "", "", err
	}
	defer file.Close()
	// credentials from the shared credentials file ~/.aws/credentials.
	sess, _ := session.NewSession(&aws.Config{
		Region:      aws.String(awsClient.AWSConfig.Region),
		Credentials: credentials.NewStaticCredentials(awsClient.AWSConfig.ID, awsClient.AWSConfig.Token, ""),
	})

	uploader := s3manager.NewUploader(sess)
	logger.Info(ctx, fmt.Sprintf("Uploading to S3 : bucket : %s, key : %s, body : %v", bucket, fileKey, file))
	result, err := uploader.Upload(&s3manager.UploadInput{
		Bucket:  aws.String(bucket),
		Key:     aws.String(folder + filepath.Base(fileKey)),
		Body:    file,
		ACL:     aws.String(accessPermission),
		Tagging: aws.String(tagsData),
	})
	if err != nil {
		logger.Error(ctx, err)
		return "", "", err
	}
	return result.Location, folder + filepath.Base(fileKey), nil
}

// GetUserAWSS3Assets used to user assets.
func (awsClient AWSClientImpl) GetUserAWSS3Assets(ctx context.Context, bucket,
	userId string) (responseArray []dto.S3FileResponseStruct, err error) {
	svc := s3.New(session.New(), &aws.Config{
		Region:      aws.String(awsClient.AWSConfig.Region),
		Credentials: credentials.NewStaticCredentials(awsClient.AWSConfig.ID, awsClient.AWSConfig.Token, ""),
	})
	bucketPrefixValue := userId + constants.PathSeparator
	params := &s3.ListObjectsInput{
		Bucket: aws.String(bucket),
		Prefix: aws.String(bucketPrefixValue),
	}

	resp, err := svc.ListObjects(params)
	awsURL := constants.HttpPrefix + bucket + constants.S3Extension + awsClient.AWSConfig.Region + constants.S3Suffix
	for _, key := range resp.Contents {
		data := dto.S3FileResponseStruct{}
		data.Tag = *key.ETag
		data.Key = *key.Key
		data.URL = awsURL + "/" + data.Key
		responseArray = append(responseArray, data)
	}
	if err != nil {
		logger.Error(ctx, err)
		return responseArray, err
	}
	return responseArray, err
}

// CreateAWSFolder used to create folder on aws
func (awsClient AWSClientImpl) CreateAWSFolder(ctx context.Context, userId, bucket, folderName,
	folderPermission string) (string, error) {
	fmt.Println(folderPermission)
	svc := s3.New(session.New(), &aws.Config{
		Region:      aws.String(awsClient.AWSConfig.Region),
		Credentials: credentials.NewStaticCredentials(awsClient.AWSConfig.ID, awsClient.AWSConfig.Token, ""),
	})
	AWSFolderName := userId + constants.PathSeparator + folderName + constants.PathSeparator
	metaData := map[string]*string{"user-id": aws.String("vinay")}
	putObjectInput := &s3.PutObjectInput{
		Bucket: aws.String(bucket), Key: aws.String(AWSFolderName), ACL: aws.String(folderPermission),
		Metadata: metaData,
	}
	_, err := svc.PutObject(putObjectInput)
	if err != nil {
		logger.Error(ctx, err)
		return "", err
	}
	response := constants.FolderCreatedResponse
	return response, nil
}

// UpdateAWSFolderName using the source and target destination.
func (awsClient AWSClientImpl) UpdateAWSFolderName(ctx context.Context, userId, bucket, source,
	target string, filesName []string) (string, error) {
	svc := s3.New(session.New(), &aws.Config{
		Region:      aws.String(awsClient.AWSConfig.Region),
		Credentials: credentials.NewStaticCredentials(awsClient.AWSConfig.ID, awsClient.AWSConfig.Token, ""),
	})
	AWSFolderName := userId + constants.PathSeparator + source + constants.PathSeparator
	for _, value := range filesName {
		sourceObject := userId + constants.PathSeparator + source + constants.PathSeparator + value
		sourceName := bucket + constants.PathSeparator + sourceObject
		targetName := userId + constants.PathSeparator + target + constants.PathSeparator + value
		copyObjectInput := &s3.CopyObjectInput{
			Bucket: aws.String(bucket), CopySource: aws.String(sourceName),
			Key: aws.String(targetName),
		}
		_, err := svc.CopyObject(copyObjectInput)
		if err != nil {
			logger.Error(ctx, err)
			return err.Error(), err
		}
		time.Sleep(1)
		fmt.Println("Source object which we want to delete ----->>>", sourceObject)
		deleteObjectInput := &s3.DeleteObjectInput{
			Bucket: aws.String(bucket), Key: aws.String(sourceObject),
		}
		// delete the previous record
		_, deleteErr := svc.DeleteObject(deleteObjectInput)
		if deleteErr != nil {
			logger.Error(ctx, deleteErr)

		}
	}
	putObjectInput := &s3.PutObjectInput{
		Bucket: aws.String(bucket), Key: aws.String(AWSFolderName),
	}
	_, err := svc.PutObject(putObjectInput)
	if err != nil {
		logger.Error(ctx, err)
	}
	response := "Files moved successfully."
	return response, nil

}

// DeleteAWSAssets used to delete the assets from aws
func (awsClient AWSClientImpl) DeleteAWSAssets(ctx context.Context, userId, bucket string,
	assetsArray []string) (string, error) {
	svc := s3.New(session.New(), &aws.Config{
		Region:      aws.String(awsClient.AWSConfig.Region),
		Credentials: credentials.NewStaticCredentials(awsClient.AWSConfig.ID, awsClient.AWSConfig.Token, ""),
	})

	for _, assetValue := range assetsArray {
		updatedAssetValue := userId + constants.PathSeparator + assetValue
		deleteObjectInput := &s3.DeleteObjectInput{Bucket: aws.String(bucket), Key: aws.String(updatedAssetValue)}
		_, err := svc.DeleteObject(deleteObjectInput)
		if err != nil {
			logger.Error(ctx, err)
		}
	}
	response := constants.AWSAssetsDeleteResponse
	return response, nil
}
