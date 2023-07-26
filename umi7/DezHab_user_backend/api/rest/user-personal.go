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

// GetPersonalInfo is used to get personal info on the basis of ID

var GetPersonalInfo = func(ctx context.Context, req *http.Request, request common.Validator) handler.Response {
	getPersonalInfoRequest, ok := request.(*dto.GetUserPersonalInfoRequest)
	if !ok {
		return api.HandleError(ctx, errors.New(TypeErr), errorHandler.ServerErr)
	}
	personalInfo, err := impl.PersonalInfo.Get(ctx, getPersonalInfoRequest.ID)
	if err != nil {
		return api.HandleMerchantError(ctx, err)
	}
	// now check the authorization of the request
	requestType := utils.GetRequestType(req)
	// if authorization then return all data
	if requestType {
		return handler.Response{
			Code: http.StatusOK,
			Payload: dto.GetPrivatePersonalInfoResponse{
				Response: dto.GetPrivatePersonalInfoData{
					Id: personalInfo.ID, UserId: personalInfo.UserId, CompositeScore: personalInfo.CompositeScore,
					FirstName: personalInfo.FirstName, MiddleName: personalInfo.MiddleName,
					LastName: personalInfo.LastName, Email: personalInfo.Email,
					PrimaryContact: personalInfo.PrimaryContact, SecondaryContact: personalInfo.SecondaryContact,
					PinCode: personalInfo.PinCode, Address: personalInfo.Address,
					ResidentState: personalInfo.ResidentState, UserAliasName: personalInfo.UserAliasName,
					RegisterEdas: personalInfo.RegisterEdas, IsBlocked: personalInfo.IsBlocked,
					IsVerified: personalInfo.IsVerified, PendingPersonalDetail: personalInfo.PendingPersonalDetail,
					PendingProfessionalDetail: personalInfo.PendingProfessionalDetail, GoogleID: personalInfo.GoogleID,
					WebProfile: personalInfo.WebProfile, LinkedIn: personalInfo.WebProfile,
					DateOfBirth: utils.DateString(personalInfo.DateOfBirth),
				},
			},
		}
	}
	// Else return the empty record
	return handler.Response{
		Code: 403,
		Payload: dto.GetPublicPersonalInfoResponse{
			Response: dto.GetPublicPersonalInfoData{
				Message: constants.NotAuthorized,
			},
		},
	}
}

// CreatePersonalInfo controller used to create the personal Info.

var CreatePersonalInfo = func(ctx context.Context, req *http.Request, request common.Validator) handler.Response {
	params, ok := request.(*dto.CreatePersonalInfoRequest)
	if !ok {
		return api.HandleError(ctx, errors.New(TypeErr), errorHandler.ServerErr)
	}
	// now check the authorization of the request
	requestType := utils.GetRequestType(req)
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
	// validate whether user id is present in table or not
	exists := utils.ValidateUserId(params.UserID)
	if exists {
		return handler.Response{
			Code: http.StatusBadRequest,
			Payload: dto.CreatePublicPersonalInfoResponse{
				Response: dto.CreatePublicPersonalInfoData{
					Message: constants.USERIDAlreadyPresent,
				},
			},
		}
	}
	dateOfBirthObject := utils.DatetimeObject(params.DateOfBirth)
	// If validate request type then create the Personal info detail.
	personalId, err := impl.PersonalInfo.Create(ctx, params.UserID, params.FirstName, params.MiddleName,
		params.LastName, params.Email, params.PrimaryContact, params.SecondaryContact, params.PinCode,
		params.Address, params.ResidentState, params.UserAliasName, params.RegisterEdas, params.GoogleID,
		params.WebProfile, params.LinkedIn, dateOfBirthObject, params.IsBlocked, params.IsVerified,
		params.PendingPersonalDetail, params.PendingProfessionalDetail)
	if err != nil {
		return api.HandleMerchantError(ctx, err)
	}
	return handler.Response{
		Code: http.StatusCreated,
		Payload: dto.CreatePublicPersonalInfoResponse{
			Response: dto.CreatePublicPersonalInfoData{
				Message: constants.PersonalInfoCreated,
				Id:      personalId,
			},
		},
	}

}

// DeletePersonalInfo controller is used to delete the PersonalInfo.
var DeletePersonalInfo = func(ctx context.Context, req *http.Request, request common.Validator) handler.Response {
	params, ok := request.(*dto.DeletePersonalInfoRequest)
	if !ok {
		return api.HandleError(ctx, errors.New(TypeErr), errorHandler.ServerErr)
	}
	// now check the authorization of the request
	requestType := utils.GetRequestType(req)
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
	err := impl.PersonalInfo.Delete(ctx, params.Id)
	if err != nil {
		return api.HandleMerchantError(ctx, err)
	}
	// now return the response
	return handler.Response{
		Code: http.StatusOK,
		Payload: dto.DeletePersonalInfoResponse{
			Response: dto.DeletePersonalInfoData{
				Message: constants.PersonalInfoDeleted,
			},
		},
	}
}

//UpdatePersonalInfo controller is used to update the Personal info detail

var UpdatePersonalInfo = func(ctx context.Context, req *http.Request, request common.Validator) handler.Response {
	params, ok := request.(*dto.UpdatePersonalInfoRequest)
	if !ok {
		return api.HandleError(ctx, errors.New(TypeErr), errorHandler.ServerErr)
	}
	// now check the authorization of the request
	requestType := utils.GetRequestType(req)
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
	dateOfBirthObject := utils.DatetimeObject(params.DateOfBirth)
	err := impl.PersonalInfo.Update(ctx, params.Id, params.FirstName, params.MiddleName, params.LastName, params.PrimaryContact,
		params.SecondaryContact, params.PinCode, params.Address, params.ResidentState, params.UserAliasName,
		params.RegisterEdas, params.GoogleID, params.WebProfile, params.LinkedIn, dateOfBirthObject, params.IsBlocked,
		params.IsVerified, params.PendingPersonalDetail, params.PendingProfessionalDetail)
	if err != nil {
		return api.HandleMerchantError(ctx, err)
	}
	// now return the response
	return handler.Response{
		Code: http.StatusOK,
		Payload: dto.UpdatePersonalInfoResponse{
			Response: dto.UpdatePersonalInfoData{
				Message: constants.PersonalInfoUpdated,
			},
		},
	}
}
