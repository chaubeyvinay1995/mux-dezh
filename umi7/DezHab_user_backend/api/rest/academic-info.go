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

// GetAcademic controller used to get the Academic detail
var GetAcademic = func(ctx context.Context, req *http.Request, request common.Validator) handler.Response{
	params, ok := request.(*dto.GetAcademicRequest)
	if !ok {
		return api.HandleError(ctx, errors.New(TypeErr), errorHandler.ServerErr)
	}

	academicInfo, err := impl.AcademicInfo.Get(ctx, params.OrganizationId)
	if err != nil {
		return api.HandleMerchantError(ctx, err)
	}

	// now check the authorization of the request
	requestType := utils.GetRequestType(req)
	if requestType {
		// return the public
		return handler.Response{
			Code:    http.StatusOK,
			Payload: dto.GetAcademicDetailPublicResponse{
				Response: academicInfo,
			},
		}
	}else{
		// return the private response
		var responseArray []dto.AcademicPrivateResponse
		for _, value := range academicInfo{
			data := dto.AcademicPrivateResponse{}
			data.OrganizationId = value.OrganizationId
			data.Id = value.ID
			data.Degree = value.Degree
			responseArray = append(responseArray, data)
			//responseArray = append(responseArray, projectStruct)
		}
		return handler.Response{
			Code:    http.StatusOK,
			Payload: dto.GetAcademicDetailPrivateResponse{
				Response: responseArray,
			},
		}
	}
}

// CreateAcademic controller used to create the academic detail.
var CreateAcademic = func(ctx context.Context, req *http.Request, request common.Validator) handler.Response{
	params, ok := request.(*dto.CreateAcademicRequest)
	if !ok {
		return api.HandleError(ctx, errors.New(TypeErr), errorHandler.ServerErr)
	}

	academicId, err := impl.AcademicInfo.Create(ctx, params.OrganizationId, params.Degree, params.InstituteName,
		params.CertificateLink)
	if err != nil {
		return api.HandleMerchantError(ctx, err)
	}

	// return the response
	return handler.Response{
		Code: http.StatusOK,
		Payload: dto.CreateAcademicResponse{
			Response: dto.CreateAcademicData{
				AcademicId: academicId,
				Message:        constants.AcademicCreated,
			},
		},
	}
}


// UpdateAcademic is used to update the Academic.
var UpdateAcademic = func(ctx context.Context, req *http.Request, request common.Validator) handler.Response {
	params, ok := request.(*dto.UpdateAcademicDetailRequest)
	if !ok {
		return api.HandleError(ctx, errors.New(TypeErr), errorHandler.ServerErr)
	}

	err := impl.AcademicInfo.Update(ctx, params.Id, params.Degree, params.InstituteName, params.CertificateLink)
	if err != nil {
		return api.HandleMerchantError(ctx, err)
	}

	// now return the response
	return handler.Response{
		Code: http.StatusOK,
		Payload: dto.UpdateAcademicResponse{
			Response: dto.UpdateAcademicData{
				AcademicId: params.Id,
				Message: constants.AcademicUpdated,
			},
		},
	}

}

// DeleteAcademic is used to delete the Project.
var DeleteAcademic = func( ctx context.Context, req *http.Request, request common.Validator) handler.Response{
	params, ok := request.(*dto.DeleteAcademicRequest)
	if !ok {
		return api.HandleError(ctx, errors.New(TypeErr), errorHandler.ServerErr)
	}

	err := impl.AcademicInfo.Delete(ctx, params.Id)
	if err != nil {
		return api.HandleMerchantError(ctx, err)
	}

	// now return the response
	return handler.Response{
		Code: http.StatusOK,
		Payload: dto.DeleteAcademicResponse{
			Response: dto.DeleteAcademicData{
				Message: constants.AcademicDeleted,
			},
		},
	}
}

