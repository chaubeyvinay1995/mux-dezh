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

// GetOrganizationInfo used to create the organization info detail
var GetOrganizationInfo = func(ctx context.Context, req *http.Request, request common.Validator) handler.Response {
	getOrganizationInfoRequest, ok := request.(*dto.GetOrganizationInfoRequest)
	if !ok {
		return api.HandleError(ctx, errors.New(TypeErr), errorHandler.ServerErr)
	}

	organizationInfo, err := impl.OrganizationInfo.Get(ctx, getOrganizationInfoRequest.ID)
	if err != nil {
		return api.HandleMerchantError(ctx, err)
	}
	// now check the authorization of the request
	requestType := utils.GetRequestType(req)
	if requestType {
		// now return the private response if request header with authorization has bearer token
		return handler.Response{
			Code: http.StatusOK,
			Payload: dto.GetOrganizationPrivateInfoResponse{
				Response: dto.GetOrganizationPrivateResponse{
					Id: organizationInfo.ID, UserID: organizationInfo.UserId, EmailId: organizationInfo.EmailId,
					EntityType: organizationInfo.EntityType, ContactPersonName: organizationInfo.ContactPersonName,
					FirstName: organizationInfo.FirstName, MiddleName: organizationInfo.MiddleName,
					LastName: organizationInfo.LastName, State: organizationInfo.State,
					City: organizationInfo.City, Country: organizationInfo.Country, Gender: organizationInfo.Gender,
					Address1: organizationInfo.Address1, Address2: organizationInfo.Address2,
					PrimaryContactNumber:   organizationInfo.PrimaryContactNumber,
					SecondaryContactNumber: organizationInfo.SecondaryContactNumber,
					WebsiteUrl:             organizationInfo.WebsiteUrl, FacebookUrl: organizationInfo.FacebookUrl,
					About: organizationInfo.About, PreferredContactInTime: organizationInfo.PreferredContactInTime,
					PreferredContactOutTime: organizationInfo.PreferredContactOutTime,
					ProfessionalExperience:  organizationInfo.ProfessionalExperience,
				},
			},
		}
	} else {
		// now return the public response if request header has no authorization has bearer token
		return handler.Response{
			Code: http.StatusOK,
			Payload: dto.GetOrganizationPublicInfoResponse{
				Response: dto.GetOrganizationPublicResponse{
					CompositeScore: organizationInfo.CompositeScore, About: organizationInfo.About,
					ProfessionalExperience: organizationInfo.ProfessionalExperience,
				},
			},
		}
	}

}

// CreateOrganizationInfo used to create the organization info detail
var CreateOrganizationInfo = func(ctx context.Context, req *http.Request, request common.Validator) handler.Response {
	params, ok := request.(*dto.CreateOrganizationInfoRequest)
	if !ok {
		return api.HandleError(ctx, errors.New(TypeErr), errorHandler.ServerErr)
	}
	inTime := utils.DatetimeObject(params.PreferredContactInTime)
	outTime := utils.DatetimeObject(params.PreferredContactOutTime)
	organizationId, err := impl.OrganizationInfo.Create(ctx, params.UserID, params.EntityType, params.FirstName, params.MiddleName,
		params.LastName, params.ContactPersonName, params.Gender, params.City, params.State, params.Country,
		params.Address1, params.Address2, params.PrimaryContactNumber, params.SecondaryContactNumber,
		params.EmailId, params.WebsiteUrl, params.FacebookUrl, params.About, params.ProfessionalExperience,
		inTime, outTime)
	if err != nil {
		return api.HandleMerchantError(ctx, err)
	}

	// now return the response
	return handler.Response{
		Code: http.StatusOK,
		Payload: dto.CreateOrganizationInfoResponse{
			Response: dto.CreateOrganizationResponse{
				UserId:  organizationId,
				Message: constants.OrganizationCreated,
			},
		},
	}
}

// UpdateOrganizationInfo controller is used to update the organization.
var UpdateOrganizationInfo = func(ctx context.Context, req *http.Request, request common.Validator) handler.Response {
	params, ok := request.(*dto.UpdateOrganizationInfoRequest)
	if !ok {
		return api.HandleError(ctx, errors.New(TypeErr), errorHandler.ServerErr)
	}

	err := impl.OrganizationInfo.Update(ctx, params.Id, params.UserID, params.EntityType, params.FirstName,
		params.MiddleName, params.LastName, params.ContactPersonName, params.Gender, params.City, params.State,
		params.Country, params.Address1, params.Address2, params.PrimaryContactNumber, params.SecondaryContactNumber,
		params.WebsiteUrl, params.FacebookUrl, params.About, params.ProfessionalExperience)
	if err != nil {
		return api.HandleMerchantError(ctx, err)
	}

	// now return the response
	return handler.Response{
		Code: http.StatusOK,
		Payload: dto.UpdateOrganizationInfoResponse{
			Response: dto.UpdateOrganizationResponse{
				Id:      params.Id,
				Message: constants.OrganizationUpdated,
			},
		},
	}
}
