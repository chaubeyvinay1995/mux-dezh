package dto

import (
	"gitlab.com/umi7/DezHab_user_backend/common"
)

// UpdateS3FolderRequest struct used to validate the UpdateS3Folder Request
type UpdateS3FolderRequest struct {
	Source    string   `json:"source" validate:"required"`
	Target    string   `json:"target" validate:"required"`
	FileNames []string `json:"fileNames" validate:"required"`
	UserId    string   `json:"userId" validate:"required"`
}

// Validate --> Call Validate method used to validate the UpdateS3Folder
func (ca UpdateS3FolderRequest) Validate() error {
	return common.Validate.Struct(ca)
}

// CreateS3FolderRequest struct used to validate the CreateS3FolderRequest Request
type CreateS3FolderRequest struct {
	FolderName       string `json:"folderName" validate:"required"`
	FolderPermission string `json:"folderPermission" validate:"required"`
	UserId           string `json:"userId" validate:"required"`
}

// Validate --> Call Validate method used to validate the CreateS3FolderRequest
func (ca CreateS3FolderRequest) Validate() error {
	return common.Validate.Struct(ca)
}

// RenameS3FolderNameRequest struct used to validate the UpdateS3FolderNameRequest Request
type RenameS3FolderNameRequest struct {
	OldFolderName string `json:"oldFolderName" validate:"required"`
	NewFolderName string `json:"newFolderName" validate:"required"`
	UserId        string `json:"userId" validate:"required"`
}

// Validate --> Call Validate method used to validate the RenameS3FolderNameRequest
func (ca RenameS3FolderNameRequest) Validate() error {
	return common.Validate.Struct(ca)
}

// CopyS3FolderNameRequest struct used to validate the CopyS3FolderNameRequest Request
type CopyS3FolderNameRequest struct {
	Source    string   `json:"source" validate:"required"`
	Target    string   `json:"target" validate:"required"`
	FileNames []string `json:"fileNames" validate:"required"`
	UserId    string   `json:"userId" validate:"required"`
}

// Validate --> Call Validate method used to validate the RenameS3FolderNameRequest
func (ca CopyS3FolderNameRequest) Validate() error {
	return common.Validate.Struct(ca)
}

// DeleteS3AssetsRequest struct used to validate the DeleteS3AssetsRequest Request
type DeleteS3AssetsRequest struct {
	Assets []string `json:"assets" validate:"required"`
	UserId string   `json:"userId" validate:"required"`
}

// Validate --> Call Validate method used to validate the RenameS3FolderNameRequest
func (ca DeleteS3AssetsRequest) Validate() error {
	return common.Validate.Struct(ca)
}

type FileUploadCertificationResponse struct {
	Response FileUploadResponse `json:"response"`
}

type FileUploadResponse struct {
	Message  string `json:"message"`
	FilePath string `json:"filePath"`
}

type AWSFileMovedResponse struct {
	Response FileMovedResponse `json:"response"`
}

type FileMovedResponse struct {
	Message string `json:"message"`
}

type CreateAWSS3Response struct {
	Response AWSS3Response `json:"response"`
}

type AWSS3Response struct {
	Message string `json:"message"`
}

type GetS3FileResponse struct {
	Response []S3FileResponseStruct `json:"response"`
}

type S3FileResponseStruct struct {
	Tag string `json:"tag"`
	Key string `json:"key"`
	URL string `json:"url"`
}
