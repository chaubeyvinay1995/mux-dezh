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

var GetProject = func(ctx context.Context, req *http.Request, request common.Validator) handler.Response {
	getProjectRequest, ok := request.(*dto.GetProjectRequest)
	if !ok {
		return api.HandleError(ctx, errors.New(TypeErr), errorHandler.ServerErr)
	}

	projectInfo, err := impl.ProjectInfo.Get(ctx, getProjectRequest.OrganizationId)
	if err != nil {
		return api.HandleMerchantError(ctx, err)
	}

	// now check the authorization of the request
	requestType := utils.GetRequestType(req)
	// return the public response
	if requestType{
		return handler.Response{
			Code:    http.StatusOK,
			Payload: dto.GetProjectPrivateInfoResponse{
				Response: projectInfo,
			},
		}
	} else{
		var responseArray []dto.ProjectPublicResponse
		for _, value := range projectInfo{
			data := dto.ProjectPublicResponse{}
			data.OrganizationId = value.OrganizationId
			data.Id = value.ID
			data.ProjectType = value.ProjectType
			responseArray = append(responseArray, data)
		}
		return handler.Response{
			Code: http.StatusOK,
			Payload: dto.GetProjectPublicInfoResponse{
				Response: responseArray,
			},
		}

	}
	// Else return the private response
}

// CreateProject controller used to create the project.
var CreateProject = func(ctx context.Context, req *http.Request, request common.Validator) handler.Response {
	params, ok := request.(*dto.CreateProjectRequest)
	if !ok {
		return api.HandleError(ctx, errors.New(TypeErr), errorHandler.ServerErr)
	}

	projectId, err := impl.ProjectInfo.Create(ctx, params.OrganizationId, params.ProjectArea, params.ProjectField,
		params.ProjectType, params.Status, params.Description, params.Location)
	if err != nil {
		return api.HandleMerchantError(ctx, err)
	}

	// return the response
	return handler.Response{
		Code: http.StatusOK,
		Payload: dto.CreateProjectResponse{
			Response: dto.CreateProjectData{
				OrganizationId: projectId,
				Message:        constants.ProjectCreated,
			},
		},
	}
}

// UpdateProject controller is used to update the project.
var UpdateProject = func(ctx context.Context, req *http.Request, request common.Validator) handler.Response {
	params, ok := request.(*dto.UpdateProjectRequest)
	if !ok {
		return api.HandleError(ctx, errors.New(TypeErr), errorHandler.ServerErr)
	}

	err := impl.ProjectInfo.Update(ctx, params.Id, params.ProjectArea, params.ProjectField, params.ProjectType,
		params.Status, params.Description, params.Location)
	if err != nil {
		return api.HandleMerchantError(ctx, err)
	}

	// now return the response
	return handler.Response{
		Code: http.StatusOK,
		Payload: dto.UpdateProjectResponse{
			Response: dto.UpdateProjectData{
				Id: params.Id,
				Message: constants.ProjectUpdated,
			},
		},
	}

}

// DeleteProject controller is used to delete the Project.
var DeleteProject = func( ctx context.Context, req *http.Request, request common.Validator) handler.Response{
	params, ok := request.(*dto.DeleteProjectRequest)
	if !ok {
		return api.HandleError(ctx, errors.New(TypeErr), errorHandler.ServerErr)
	}

	err := impl.ProjectInfo.Delete(ctx, params.Id)
	if err != nil {
		return api.HandleMerchantError(ctx, err)
	}

	// now return the response
	return handler.Response{
		Code: http.StatusOK,
		Payload: dto.DeleteProjectResponse{
			Response: dto.DeleteProjectData{
				Id: params.Id,
				Message: constants.ProjectDeleted,
			},
		},
	}
}
