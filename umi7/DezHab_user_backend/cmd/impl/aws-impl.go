package impl

import (
	"context"
	"fmt"
	"github.com/go-cmd/cmd"
	awsCmd "gitlab.com/umi7/DezHab_user_backend/cmd"
	"gitlab.com/umi7/DezHab_user_backend/common/logger"
	"gitlab.com/umi7/DezHab_user_backend/config"
	"gitlab.com/umi7/DezHab_user_backend/dto"
	"gitlab.com/umi7/DezHab_user_backend/utils/constants"
)

var Command awsCmd.Command

type AWSCommandImpl struct {
	AWSConfig dto.AWSConfig
}

func init() {
	Command = AWSCommandImpl{
		AWSConfig: dto.AWSConfig{
			ID:     config.AppConfig.Aws.ID,
			Token:  config.AppConfig.Aws.Token,
			Region: config.AppConfig.Aws.Region,
		},
	}
}

func init() {
}

// Rename used to rename one folder name
func (awsCommand AWSCommandImpl) Rename(ctx context.Context, bucket, userId, source,
	target string) (string, error) {
	s3BucketPrefix := constants.S3Prefix + bucket + constants.PathSeparator + userId + constants.PathSeparator
	sourceFolderName := s3BucketPrefix + source
	destinationFolderName := s3BucketPrefix + target
	c := cmd.NewCmd("aws", "s3", "--recursive", "mv", sourceFolderName, destinationFolderName)
	s := <-c.Start()
	err := s.Error
	if err != nil {
		logger.Error(ctx, fmt.Sprintf("Error occurred while renaming the folder name: %v", err))
		return "", err
	}
	// Below command is folder rename
	folder := cmd.NewCmd("aws", "s3", "mv", sourceFolderName, destinationFolderName)
	folderS := <-folder.Start()
	folderErr := folderS.Error
	if folderErr != nil {
		logger.Error(ctx, fmt.Sprintf("Error occurred while renaming the folder name: %v", err))
		return "", err
	}
	response := "Renamed successfully."
	return response, nil
}

// Copy used to Copy one folder to another folder name
func (awsCommand AWSCommandImpl) Copy(ctx context.Context, bucket, userId, source,
	destination string, filesArray []string) (string, error) {
	s3BucketPrefix := constants.S3Prefix + bucket + constants.PathSeparator + userId + constants.PathSeparator
	sourceFolderName := s3BucketPrefix + source
	destinationFolderName := s3BucketPrefix + destination
	for _, value := range filesArray {
		sourceFileName := sourceFolderName + value
		c := cmd.NewCmd("aws", "s3", "cp", sourceFileName, destinationFolderName)
		s := <-c.Start()
		err := s.Error
		if err != nil {
			logger.Error(ctx, fmt.Sprintf("Error occurred while copying the folder name: %v", err))
		}
	}
	response := "Copied files successfully."
	return response, nil
}

// Move used to Move files from one folder to other Folder
func (awsCommand AWSCommandImpl) Move(ctx context.Context, bucket, organizationId,
	source, destination string, filesArray []string) (string, error) {
	s3BucketPrefix := constants.S3Prefix + bucket + constants.PathSeparator
	sourceFolderName := s3BucketPrefix + source
	destinationFolderName := s3BucketPrefix + destination

	for _, value := range filesArray {
		sourceFileName := sourceFolderName + value
		c := cmd.NewCmd("aws", "s3", "mv", sourceFileName, destinationFolderName, "--recursive")
		s := <-c.Start()
		err := s.Error
		if err != nil {
			logger.Error(ctx, fmt.Sprintf("Error occurred while copying the folder name: %v", err))
		}
	}
	response := "Moved files successfully."
	return response, nil
}

// ZipFolder used to Zip folder
func (awsCommand AWSCommandImpl) ZipFolder(ctx context.Context, bucket, organizationId,
	source, destination string, filesArray []string) (string, error) {
	s3BucketPrefix := "s3://" + bucket + "/" + "organizations/"
	sourceFolderName := s3BucketPrefix + organizationId + "/" + constants.AWSFolderName + source
	destinationFolderName := s3BucketPrefix + organizationId + "/" + constants.AWSFolderName + destination
	fmt.Println("AAAAAAAAAAa", destinationFolderName)
	c := cmd.NewCmd("aws", "s3", "sync", sourceFolderName, destinationFolderName, "-- recursive")
	s := <-c.Start()
	err := s.Error
	if err != nil {
		logger.Error(ctx, fmt.Sprintf("Error occurred while zip folder: %v", err))
	}
	response := "Copied files successfully."
	return response, nil
}
