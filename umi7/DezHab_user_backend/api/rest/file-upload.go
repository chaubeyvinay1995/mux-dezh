package rest

import (
	"context"
	"errors"
	"gitlab.com/umi7/DezHab_user_backend/api"
	apiUtils "gitlab.com/umi7/DezHab_user_backend/api"
	"gitlab.com/umi7/DezHab_user_backend/api/handler"
	awsCmd "gitlab.com/umi7/DezHab_user_backend/cmd/impl"
	"gitlab.com/umi7/DezHab_user_backend/common"
	"gitlab.com/umi7/DezHab_user_backend/common/logger"
	"gitlab.com/umi7/DezHab_user_backend/config"
	"gitlab.com/umi7/DezHab_user_backend/dto"
	errorHandler "gitlab.com/umi7/DezHab_user_backend/error-handler"
	"gitlab.com/umi7/DezHab_user_backend/service/external/impl"
	"gitlab.com/umi7/DezHab_user_backend/utils"
	"gitlab.com/umi7/DezHab_user_backend/utils/constants"
	"io"
	"net/http"
	"os"
	"path"
	"path/filepath"
)

// FileUpload controller used to upload files to s3.
var FileUpload = func(req *http.Request) handler.Response {
	var (
		ctx = context.Background()
		err error
	)
	file, header, err := req.FormFile("file")
	if err != nil {
		logger.Error(ctx, err)
		return api.HandleError(ctx, err, errorHandler.RequestErr)
	}

	// now check the authorization of the request
	requestType := apiUtils.GetRequestType(req)
	if !requestType {
		return handler.Response{
			Code: 403,
			Payload: dto.GetPublicPersonalInfoResponse{
				Response: dto.GetPublicPersonalInfoData{
					Message: constants.NotAuthorized,
				},
			},
		}
	}

	fileKey := utils.Environment() + "_" + header.Filename
	writer, err := os.Create(fileKey)
	if err != nil {
		logger.Error(ctx, err)
		return api.HandleError(ctx, err, errorHandler.RequestErr)
	}
	defer writer.Close()
	defer os.Remove(fileKey)

	_, err = io.Copy(writer, file)
	fileAbsPath, err := filepath.Abs(filepath.Dir(writer.Name()))
	if err != nil {
		logger.Error(ctx, err)
		return api.HandleError(ctx, err, errorHandler.RequestErr)
	}
	filePath := path.Join(fileAbsPath, fileKey)
	defer os.Remove(filePath)
	// Now call function to validate the file extension
	if !apiUtils.CheckFileExtension(filepath.Ext(fileKey)) {
		return handler.Response{
			Code: 400,
			Payload: dto.GetPublicPersonalInfoResponse{
				Response: dto.GetPublicPersonalInfoData{
					Message: constants.FileTypeNotSupported,
				},
			},
		}
	}

	s3Folder := req.Form.Get("userId") + constants.PathSeparator + req.Form.Get("path") + constants.PathSeparator
	var s3FilePath string
	s3FilePath, _, err = impl.Aws.PushFileToS3(ctx, config.AppConfig.Aws.Bucket, s3Folder, filePath,
		req.Form.Get("tags"), req.Form.Get("accessPermission"))
	// now return the response
	return handler.Response{
		Code: http.StatusOK,
		Payload: dto.FileUploadCertificationResponse{
			Response: dto.FileUploadResponse{
				Message:  constants.FileUploadResponse,
				FilePath: s3FilePath,
			},
		},
	}
}

// GetS3Assets controller used to get the folder structure and their data from s3.
var GetS3Assets = func(ctx context.Context, req *http.Request, request common.Validator) handler.Response {
	params, ok := request.(*dto.GetCertificateRequest)
	if !ok {
		return api.HandleError(ctx, errors.New(TypeErr), errorHandler.ServerErr)
	}
	// now check the authorization of the request
	requestType := apiUtils.GetRequestType(req)
	if !requestType {
		return handler.Response{
			Code: 403,
			Payload: dto.GetPublicPersonalInfoResponse{
				Response: dto.GetPublicPersonalInfoData{
					Message: constants.NotAuthorized,
				},
			},
		}
	}

	responseArray, _ := impl.Aws.GetUserAWSS3Assets(ctx, config.AppConfig.Aws.Bucket, params.UserId)
	return handler.Response{
		Code: http.StatusOK,
		Payload: dto.GetS3FileResponse{
			Response: responseArray,
		},
	}
}

// CreateS3Folder controller used to create the folder
var CreateS3Folder = func(ctx context.Context, req *http.Request, request common.Validator) handler.Response {
	params, ok := request.(*dto.CreateS3FolderRequest)
	if !ok {
		return api.HandleError(ctx, errors.New(TypeErr), errorHandler.ServerErr)
	}
	// now check the authorization of the request
	requestType := apiUtils.GetRequestType(req)
	if !requestType {
		return handler.Response{
			Code: 403,
			Payload: dto.GetPublicPersonalInfoResponse{
				Response: dto.GetPublicPersonalInfoData{
					Message: constants.NotAuthorized,
				},
			},
		}
	}

	response, err := impl.Aws.CreateAWSFolder(ctx, params.UserId, config.AppConfig.Aws.Bucket,
		params.FolderName, params.FolderPermission)
	if err != nil {
		return api.HandleMerchantError(ctx, err)
	}

	return handler.Response{
		Code: http.StatusOK,
		Payload: dto.CreateAWSS3Response{
			Response: dto.AWSS3Response{
				Message: response,
			},
		},
	}
}

// MoveS3Assets controller used to move the assets from Location to other location.
var MoveS3Assets = func(ctx context.Context, req *http.Request, request common.Validator) handler.Response {
	params, ok := request.(*dto.UpdateS3FolderRequest)
	if !ok {
		return api.HandleError(ctx, errors.New(TypeErr), errorHandler.ServerErr)
	}

	requestType := apiUtils.GetRequestType(req)
	if !requestType {
		return handler.Response{
			Code: 403,
			Payload: dto.GetPublicPersonalInfoResponse{
				Response: dto.GetPublicPersonalInfoData{
					Message: constants.NotAuthorized,
				},
			},
		}
	}

	response, _ := impl.Aws.UpdateAWSFolderName(ctx, params.UserId, config.AppConfig.Aws.Bucket,
		params.Source, params.Target, params.FileNames)
	return handler.Response{
		Code: http.StatusOK,
		Payload: dto.AWSFileMovedResponse{
			Response: dto.FileMovedResponse{
				Message: response,
			},
		},
	}
}

// DeleteS3Assets controller used to delete the aws assets
var DeleteS3Assets = func(ctx context.Context, req *http.Request, request common.Validator) handler.Response {
	params, ok := request.(*dto.DeleteS3AssetsRequest)
	if !ok {
		return api.HandleError(ctx, errors.New(TypeErr), errorHandler.ServerErr)
	}
	// now check the authorization of the request
	requestType := apiUtils.GetRequestType(req)
	if !requestType {
		return handler.Response{
			Code: 403,
			Payload: dto.GetPublicPersonalInfoResponse{
				Response: dto.GetPublicPersonalInfoData{
					Message: constants.NotAuthorized,
				},
			},
		}
	}

	response, err := impl.Aws.DeleteAWSAssets(ctx, params.UserId, config.AppConfig.Aws.Bucket, params.Assets)
	if err != nil {
		return api.HandleMerchantError(ctx, err)
	}

	return handler.Response{
		Code: http.StatusOK,
		Payload: dto.CreateAWSS3Response{
			Response: dto.AWSS3Response{
				Message: response,
			},
		},
	}
}

// RenameS3Folder controller used to update the folder name
var RenameS3Folder = func(ctx context.Context, req *http.Request, request common.Validator) handler.Response {
	params, ok := request.(*dto.RenameS3FolderNameRequest)
	if !ok {
		return api.HandleError(ctx, errors.New(TypeErr), errorHandler.ServerErr)
	}

	// now check the authorization of the request
	requestType := apiUtils.GetRequestType(req)
	if !requestType {
		return handler.Response{
			Code: 403,
			Payload: dto.GetPublicPersonalInfoResponse{
				Response: dto.GetPublicPersonalInfoData{
					Message: constants.NotAuthorized,
				},
			},
		}
	}

	response, err := awsCmd.Command.Rename(ctx, config.AppConfig.Aws.Bucket, params.UserId,
		params.OldFolderName, params.NewFolderName)

	if err != nil {
		return api.HandleMerchantError(ctx, err)
	}

	return handler.Response{
		Code: http.StatusOK,
		Payload: dto.CreateAWSS3Response{
			Response: dto.AWSS3Response{
				Message: response,
			},
		},
	}
}

// CopyS3Files controller used to copy the files from one folder to another folder
var CopyS3Files = func(ctx context.Context, req *http.Request, request common.Validator) handler.Response {
	params, ok := request.(*dto.CopyS3FolderNameRequest)
	if !ok {
		return api.HandleError(ctx, errors.New(TypeErr), errorHandler.ServerErr)
	}

	response, err := awsCmd.Command.Copy(ctx, config.AppConfig.Aws.Bucket, params.UserId,
		params.Source, params.Target, params.FileNames)

	if err != nil {
		return api.HandleMerchantError(ctx, err)
	}

	return handler.Response{
		Code: http.StatusOK,
		Payload: dto.CreateAWSS3Response{
			Response: dto.AWSS3Response{
				Message: response,
			},
		},
	}
}

// MoveS3Files controller used to move the files from folder to another folder
var MoveS3Files = func(ctx context.Context, req *http.Request, request common.Validator) handler.Response {
	params, ok := request.(*dto.CopyS3FolderNameRequest)
	if !ok {
		return api.HandleError(ctx, errors.New(TypeErr), errorHandler.ServerErr)
	}

	response, err := awsCmd.Command.Move(ctx, config.AppConfig.Aws.Bucket, params.UserId,
		params.Source, params.Target, params.FileNames)

	if err != nil {
		return api.HandleMerchantError(ctx, err)
	}

	return handler.Response{
		Code: http.StatusOK,
		Payload: dto.CreateAWSS3Response{
			Response: dto.AWSS3Response{
				Message: response,
			},
		},
	}
}

// ZipS3Folder controller used to zip the aws folder.
var ZipS3Folder = func(ctx context.Context, req *http.Request, request common.Validator) handler.Response {
	params, ok := request.(*dto.CopyS3FolderNameRequest)
	if !ok {
		return api.HandleError(ctx, errors.New(TypeErr), errorHandler.ServerErr)
	}

	organizationId := "bbea4e38-1930-11ec-bf9e-27b32348df8c"
	response, err := awsCmd.Command.ZipFolder(ctx, config.AppConfig.Aws.Bucket, organizationId,
		params.Source, params.Target, params.FileNames)

	if err != nil {
		return api.HandleMerchantError(ctx, err)
	}

	return handler.Response{
		Code: http.StatusOK,
		Payload: dto.CreateAWSS3Response{
			Response: dto.AWSS3Response{
				Message: response,
			},
		},
	}
}
