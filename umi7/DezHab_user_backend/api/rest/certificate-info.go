package rest

import (
	"context"
	"errors"
	"gitlab.com/umi7/DezHab_user_backend/api"
	utils "gitlab.com/umi7/DezHab_user_backend/api"
	"gitlab.com/umi7/DezHab_user_backend/api/handler"
	"gitlab.com/umi7/DezHab_user_backend/common"
	"gitlab.com/umi7/DezHab_user_backend/dto"
	errorHandler "gitlab.com/umi7/DezHab_user_backend/error-handler"
	"gitlab.com/umi7/DezHab_user_backend/service/impl"
	"gitlab.com/umi7/DezHab_user_backend/utils/constants"
	"net/http"
)

// GetCertificate controller used to get the Certificate detail
var GetCertificate = func(ctx context.Context, req *http.Request, request common.Validator) handler.Response {
	params, ok := request.(*dto.GetCertificateRequest)
	if !ok {
		return api.HandleError(ctx, errors.New(TypeErr), errorHandler.ServerErr)
	}

	certificateInfo, err := impl.CertificateInfo.Get(ctx, params.UserId)
	if err != nil {
		return api.HandleMerchantError(ctx, err)
	}

	// now check the authorization of the request
	requestType := utils.GetRequestType(req)
	// return the private response data
	if requestType {
		return handler.Response{
			Code: http.StatusOK,
			Payload: dto.GetCertificatePublicDetailResponse{
				Response: certificateInfo,
			},
		}
	} else {
		// return the public response data
		var responseArray []dto.CertificatePrivateResponse
		for _, value := range certificateInfo {
			data := dto.CertificatePrivateResponse{}
			data.CertificateId = value.ID
			data.CourseName = value.CourseName
			responseArray = append(responseArray, data)
			//responseArray = append(responseArray, projectStruct)
		}
		return handler.Response{
			Code: http.StatusOK,
			Payload: dto.GetCertificatePrivateDetailResponse{
				Response: responseArray,
			},
		}
	}
}

// CreateCertificate controller used to create the Certificate detail
var CreateCertificate = func(ctx context.Context, req *http.Request, request common.Validator) handler.Response {
	params, ok := request.(*dto.CreateCertificateRequest)
	if !ok {
		return api.HandleError(ctx, errors.New(TypeErr), errorHandler.ServerErr)
	}

	certificateId, err := impl.CertificateInfo.Create(ctx, params.OrganizationId, params.CourseName,
		params.Institute, params.Url)
	if err != nil {
		return api.HandleMerchantError(ctx, err)
	}

	// return the response
	return handler.Response{
		Code: http.StatusOK,
		Payload: dto.CreateCertificateResponse{
			Response: dto.CreateCertificateData{
				CertificateId: certificateId,
				Message:       constants.CertificateCreated,
			},
		},
	}
}

// UpdateCertificate controller is used to update the Certificate by the certificate id.
var UpdateCertificate = func(ctx context.Context, req *http.Request, request common.Validator) handler.Response {
	params, ok := request.(*dto.UpdateCertificateRequest)
	if !ok {
		return api.HandleError(ctx, errors.New(TypeErr), errorHandler.ServerErr)
	}

	err := impl.CertificateInfo.Update(ctx, params.Id, params.CourseName, params.Institute, params.Url)
	if err != nil {
		return api.HandleMerchantError(ctx, err)
	}

	// return the response
	return handler.Response{
		Code: http.StatusOK,
		Payload: dto.UpdateCertificateResponse{
			Response: dto.UpdateCertificateData{
				CertificateId: params.Id,
				Message:       constants.CertificateUpdated,
			},
		},
	}
}

// DeleteCertificate is used to delete the Project.
var DeleteCertificate = func(ctx context.Context, req *http.Request, request common.Validator) handler.Response {
	params, ok := request.(*dto.DeleteCertificateRequest)
	if !ok {
		return api.HandleError(ctx, errors.New(TypeErr), errorHandler.ServerErr)
	}

	err := impl.CertificateInfo.Delete(ctx, params.Id)
	if err != nil {
		return api.HandleMerchantError(ctx, err)
	}

	// now return the response
	return handler.Response{
		Code: http.StatusOK,
		Payload: dto.DeleteCertificateResponse{
			Response: dto.DeleteCertificateData{
				Message: constants.CertificateDeleted,
			},
		},
	}
}
