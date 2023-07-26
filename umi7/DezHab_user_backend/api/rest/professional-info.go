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

// GetProfessionalInfo is used to get Professional info on the basis of ID

var GetProfessionalInfo = func(ctx context.Context, req *http.Request, request common.Validator) handler.Response {
	getPersonalInfoRequest, ok := request.(*dto.GetUserProfessionalInfoRequest)
	if !ok {
		return api.HandleError(ctx, errors.New(TypeErr), errorHandler.ServerErr)
	}
	professionalInfo, err := impl.ProfessionalInfo.Get(ctx, getPersonalInfoRequest.ID)
	if err != nil {
		return api.HandleMerchantError(ctx, err)
	}
	// get the user detail
	personalInfo, personalInfoError := utils.GetUserDetail(getPersonalInfoRequest.ID)
	if personalInfoError != nil {
		return api.HandleMerchantError(ctx, personalInfoError)
	}
	// now check the authorization of the request
	requestType := utils.GetRequestType(req)
	// if authorization then return all data i.e private data
	if requestType {
		publicResponseArray := utils.GetProfessionalInfoPrivateResponseData(professionalInfo)
		// return the private response means return all the data
		return handler.Response{
			Code: http.StatusOK,
			Payload: dto.GetProfessionalInfoPrivateResponse{
				CompositeScore: personalInfo.CompositeScore,
				Response:       publicResponseArray,
			},
		}

	}
	publicResponseArray := utils.GetProfessionalInfoPublicResponseData(professionalInfo)
	return handler.Response{
		Code: http.StatusOK,
		Payload: dto.GetProfessionalInfoPublicResponse{
			CompositeScore: personalInfo.CompositeScore,
			Response:       publicResponseArray,
		},
	}
}

// DeleteProfessionalInfo controller is used to delete the professional info.
var DeleteProfessionalInfo = func(ctx context.Context, req *http.Request, request common.Validator) handler.Response {
	params, ok := request.(*dto.DeleteProfessionalInfoRequest)
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
	err := impl.ProfessionalInfo.Delete(ctx, params.Id)
	if err != nil {
		return api.HandleMerchantError(ctx, err)
	}
	// now return the response
	return handler.Response{
		Code: http.StatusOK,
		Payload: dto.DeleteProfessionalInfoResponse{
			Response: dto.DeleteProfessionalInfoMessage{
				Message: constants.ProfessionalInfoDeleted,
			},
		},
	}
}

// CreateProfessionalInfo controller used to create the professional Info.
var CreateProfessionalInfo = func(ctx context.Context, req *http.Request, request common.Validator) handler.Response {
	params, ok := request.(*dto.CreateProfessionalInfoRequest)
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
	passingOutDateObject := utils.DatetimeObject(params.PassingOutDate)
	// If validate request type then create the Personal info detail.
	professionalInfoId, err := impl.ProfessionalInfo.Create(ctx, params.UserId, params.UserAlias, params.EntityType,
		params.ContactHrsFrom, params.ContactMinFrom, params.ContactMinFrom, params.ContactMinTo,
		params.ContactPersonName, params.ExpYearsCategory, params.ExpertProfileDesc, params.ServiceState,
		params.InstName, params.InstDegree, params.InstDegreeUrl, params.EdLevel, params.CertName, params.CertInst,
		params.CertUrl, params.AchievementName, params.AchievementDescription,
		params.AchievementProofUrls, params.ExpertOthConstMatSup, params.ExpertOthMechSer, params.SupplierTileLabor,
		params.ExpertOthHardware, params.SupplierOthEqp, params.SupplierOthFin, params.ProjDesc, params.ProjLoc,
		params.ProjArea, params.ProjBudget, params.ProjType, &passingOutDateObject, params.ExpertSpcRes,
		params.ExpertSpcCom, params.ExpertSpcInst, params.ExpertSpcHealth, params.ExpertSpcInfra,
		params.ExpertSpcCSR, params.ExpertArch, params.ExpertArchConceptDraw, params.ExpertArchWorkingDraw,
		params.ExpertArch3DRender, params.ExpertArchScaleModel, params.ExpertArch3DModel, params.ExpertArchVirModel,
		params.ExpertArchExec, params.ExpertStruct, params.ExpertStructWorkingDraw, params.ExpertStructExec,
		params.ExpertInterior, params.ExpertIntConceptDraw, params.ExpertIntWorkingDraw, params.ExpertIntPhyScale,
		params.ExpertInt3DModel, params.ExpertIntScaleModel, params.ExpertIntVirModel, params.ExpertIntExec,
		params.ExpertLScape, params.ExpertSiteSupervisor, params.ExpertSoilTest, params.ExpertSanctions,
		params.SupplierConstMat, params.SupplierCement, params.SupplierSand, params.SupplierBrick,
		params.SupplierPrecast, params.SupplierCoarse, params.SupplierConcrete, params.SupplierReinfbar,
		params.SupplierTile, params.SupplierWproofMat, params.SupplierShuttering, params.SupplierScaffs,
		params.ExpertMechanical, params.ExpertHeating, params.ExpertAc, params.ExpertElevator, params.ExpertElectrical,
		params.ExpertElecWiring, params.ExpertElecFix, params.ExpertOthElecSer, params.ExpertPlumbing,
		params.SupplierLabor, params.SupplierCivilLabor, params.SupplierPainter, params.SupplierCarpenter,
		params.SupplierPlumbing, params.SupplierWaterTank, params.SupplierWaterPipeAndFitting, params.SupplierWaterFix,
		params.SupplierElec, params.SupplierWiringMat, params.SupplierSwitch, params.SupplierSensorSup,
		params.SupplierEleAppSup, params.SupplierOthEle, params.SupplierPaint, params.SupplierIntPaint,
		params.SupplierExtPaint, params.SupplierIntPrimer, params.SupplierExtPrimer, params.SupplierPutty,
		params.SupplierPaintAccessory, params.SupplierOtherPaint, params.ExpertFabrication, params.ExpertGlassFab,
		params.ExpertWoodFab, params.ExpertMetalAl, params.ExpertMetalFe, params.ExpertMetalSS, params.SupplierConstEqp,
		params.SupplierJcb, params.SupplierConcMixer, params.SupplierHaulTruck, params.SupplierCrane,
		params.SupplierHouseProduct, params.SupplierFurniture, params.SupplierAppliance, params.SupplierFurnish)
	if err != nil {
		return api.HandleMerchantError(ctx, err)
	}
	return handler.Response{
		Code: http.StatusCreated,
		Payload: dto.CreateProfessionalInfoResponse{
			Response: dto.CreateProfessionalInfoData{
				Message: constants.ProfessionalInfoCreated,
				Id:      professionalInfoId,
			},
		},
	}

}

// UpdateProfessionalInfo controller used to update the professional Info.
var UpdateProfessionalInfo = func(ctx context.Context, req *http.Request, request common.Validator) handler.Response {
	params, ok := request.(*dto.UpdateProfessionalInfoRequest)
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
	passingOutDateObject := utils.DatetimeObject(params.PassingOutDate)
	// If validate request type then create the Personal info detail.
	err := impl.ProfessionalInfo.Update(ctx, params.Id, params.UserAlias, params.EntityType,
		params.ContactHrsFrom, params.ContactMinFrom, params.ContactMinFrom, params.ContactMinTo,
		params.ContactPersonName, params.ExpYearsCategory, params.ExpertProfileDesc, params.ServiceState,
		params.InstName, params.InstDegree, params.InstDegreeUrl, params.EdLevel, params.CertName, params.CertInst,
		params.CertUrl, params.AchievementName, params.AchievementDescription,
		params.AchievementProofUrls, params.ExpertOthConstMatSup, params.ExpertOthMechSer, params.SupplierTileLabor,
		params.ExpertOthHardware, params.SupplierOthEqp, params.SupplierOthFin, params.ProjDesc, params.ProjLoc,
		params.ProjArea, params.ProjBudget, params.ProjType, passingOutDateObject, params.ExpertSpcRes,
		params.ExpertSpcCom, params.ExpertSpcInst, params.ExpertSpcHealth, params.ExpertSpcInfra,
		params.ExpertSpcCSR, params.ExpertArch, params.ExpertArchConceptDraw, params.ExpertArchWorkingDraw,
		params.ExpertArch3DRender, params.ExpertArchScaleModel, params.ExpertArch3DModel, params.ExpertArchVirModel,
		params.ExpertArchExec, params.ExpertStruct, params.ExpertStructWorkingDraw, params.ExpertStructExec,
		params.ExpertInterior, params.ExpertIntConceptDraw, params.ExpertIntWorkingDraw, params.ExpertIntPhyScale,
		params.ExpertInt3DModel, params.ExpertIntScaleModel, params.ExpertIntVirModel, params.ExpertIntExec,
		params.ExpertLScape, params.ExpertSiteSupervisor, params.ExpertSoilTest, params.ExpertSanctions,
		params.SupplierConstMat, params.SupplierCement, params.SupplierSand, params.SupplierBrick,
		params.SupplierPrecast, params.SupplierCoarse, params.SupplierConcrete, params.SupplierReinfbar,
		params.SupplierTile, params.SupplierWproofMat, params.SupplierShuttering, params.SupplierScaffs,
		params.ExpertMechanical, params.ExpertHeating, params.ExpertAc, params.ExpertElevator, params.ExpertElectrical,
		params.ExpertElecWiring, params.ExpertElecFix, params.ExpertOthElecSer, params.ExpertPlumbing,
		params.SupplierLabor, params.SupplierCivilLabor, params.SupplierPainter, params.SupplierCarpenter,
		params.SupplierPlumbing, params.SupplierWaterTank, params.SupplierWaterPipeAndFitting, params.SupplierWaterFix,
		params.SupplierElec, params.SupplierWiringMat, params.SupplierSwitch, params.SupplierSensorSup,
		params.SupplierEleAppSup, params.SupplierOthEle, params.SupplierPaint, params.SupplierIntPaint,
		params.SupplierExtPaint, params.SupplierIntPrimer, params.SupplierExtPrimer, params.SupplierPutty,
		params.SupplierPaintAccessory, params.SupplierOtherPaint, params.ExpertFabrication, params.ExpertGlassFab,
		params.ExpertWoodFab, params.ExpertMetalAl, params.ExpertMetalFe, params.ExpertMetalSS, params.SupplierConstEqp,
		params.SupplierJcb, params.SupplierConcMixer, params.SupplierHaulTruck, params.SupplierCrane,
		params.SupplierHouseProduct, params.SupplierFurniture, params.SupplierAppliance, params.SupplierFurnish)

	if err != nil {
		return api.HandleMerchantError(ctx, err)
	}

	return handler.Response{
		Code: http.StatusOK,
		Payload: dto.UpdateProfessionalInfoResponse{
			Response: dto.UpdateProfessionalInfoData{
				Message: constants.ProfessionalInfoUpdated,
			},
		},
	}

}
