package cmd

import (
	"context"
)

type Command interface {
	Rename(ctx context.Context, bucket string, userId string, source string,
		destination string) (string, error)
	Copy(ctx context.Context, bucket string, userId string, oldFolderName string,
		newFolderName string, filesArray []string) (string, error)
	Move(ctx context.Context, bucket string, userId string, oldFolderName string,
		newFolderName string, filesArray []string) (string, error)
	ZipFolder(ctx context.Context, bucket string, userId string, oldFolderName string,
		newFolderName string, filesArray []string) (string, error)
}
