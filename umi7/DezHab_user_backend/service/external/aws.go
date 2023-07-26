package external

import (
	"context"
	"gitlab.com/umi7/DezHab_user_backend/dto"
)

type Aws interface {
	GetUserAWSS3Assets(ctx context.Context, bucket string,
		userId string) ([]dto.S3FileResponseStruct, error)
	UpdateAWSFolderName(ctx context.Context, organizationId, bucket string, source string,
		target string, fileName []string) (string, error)
	PushFileToS3(ctx context.Context, bucket string, folder string, fileKey string, tagsData string,
		accessPermission string) (string, string, error)
	CreateAWSFolder(ctx context.Context, userId string, bucket string, folderName string,
		folderPermission string) (string, error)
	DeleteAWSAssets(ctx context.Context, userId string, bucket string, assetsArray []string) (string, error)
}

//go:generate mockgen -source=aws.go -destination=mock-impl/aws-impl.go -package=mock_impl
