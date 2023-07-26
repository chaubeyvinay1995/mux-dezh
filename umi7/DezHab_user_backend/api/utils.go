package api

import (
	"fmt"
	"gitlab.com/umi7/DezHab_user_backend/dao/core/impl"
	"gitlab.com/umi7/DezHab_user_backend/dao/models"
	"gitlab.com/umi7/DezHab_user_backend/dto"
	"gitlab.com/umi7/DezHab_user_backend/utils/constants"
	"net/http"
	"strings"
	"time"
)

// GetRequestType used to identify the incoming request has header with Authorization with Value Bearer
var GetRequestType = func(req *http.Request) bool {
	authorizationHeader := req.Header.Get("Authorization")
	return strings.HasPrefix(authorizationHeader, "Bearer")
}

// DatetimeObject function used to convert datetime string to datetime object
var DatetimeObject = func(TimeString string) time.Time {
	dateTimeObject, err := time.Parse("2006-01-02", TimeString)
	if err != nil {
		fmt.Println(err)
	}
	return dateTimeObject
}

// DateString Function used to convert the datetime object to date string.
var DateString = func(DatetimeObject time.Time) string {
	dateString := DatetimeObject.Format("2006-01-02")
	return dateString

}

// ValidateUserId function used to validate userId is present into personal info table or not.
var ValidateUserId = func(userId string) (exists bool) {
	var userPersonalInfo models.UserPersonalInfo
	error := impl.Read.First(&userPersonalInfo, models.UserPersonalInfo{UserId: userId})
	if error != nil {
		exists = false
	} else {
		exists = true
	}
	return
}

// GetUserDetail on the basis of the id.
var GetUserDetail = func(id string) (personalInfo models.UserPersonalInfo, err error) {
	var userPersonalInfo models.UserPersonalInfo
	personalInfo = userPersonalInfo
	err = impl.Read.First(&userPersonalInfo, models.UserPersonalInfo{ID: id})
	return
}

// CheckFileExtension used to check the file extension
var CheckFileExtension = func(fileExtension string) bool {
	isSupportedExtension := false
	for _, value := range constants.SupportedFileExtensions {
		if value == fileExtension {
			isSupportedExtension = true
			return isSupportedExtension
		}
	}
	return isSupportedExtension
}

// GetProfessionalInfoPrivateResponseData used to get professional info private data
var GetProfessionalInfoPrivateResponseData = func(professionalInfo []models.UserProfessionalInfo) (privateResponseData []dto.GetProfessionalInfoPrivateResponseData) {
	for _, value := range professionalInfo {
		data := dto.GetProfessionalInfoPrivateResponseData{}
		data.Id = value.ID
		data.UserId = value.UserId
		data.UserAlias = value.UserAlias
		data.EntityType = value.EntityType
		data.ContactHrsFrom = value.ContactHrsFrom
		data.ContactMinFrom = value.ContactMinFrom
		data.ContactHrsTo = value.ContactHrsTo
		data.ContactMinTo = value.ContactMinTo
		data.ContactPersonName = value.ContactPersonName
		data.ExpYearsCategory = value.ExpYearsCategory
		data.ExpertSpcRes = value.ExpertSpcRes
		data.ExpertSpcCom = value.ExpertSpcCom
		data.ExpertSpcInst = value.ExpertSpcInst
		data.ExpertSpcHealth = value.ExpertSpcHealth
		data.ExpertSpcInfra = value.ExpertSpcInfra
		data.ExpertSpcCSR = value.ExpertSpcCSR
		data.ExpertProfileDesc = value.ExpertProfileDesc
		data.ServiceState = value.ServiceState
		data.InstName = value.InstName
		data.InstDegree = value.InstDegree
		data.InstDegreeUrl = value.InstDegreeUrl
		data.EdLevel = value.EdLevel
		data.CertName = value.CertName
		data.CertInst = value.CertInst
		data.CertUrl = value.CertUrl
		data.AchievementName = value.AchievementName
		data.AchievementDescription = value.AchievementDescription
		data.AchievementProofUrls = value.AchievementProofUrls
		data.ExpertArch = value.ExpertArch
		data.ExpertArchConceptDraw = value.ExpertArchConceptDraw
		data.ExpertArchWorkingDraw = value.ExpertArchWorkingDraw
		data.ExpertArch3DRender = value.ExpertArch3DRender
		data.ExpertArchScaleModel = value.ExpertArchScaleModel
		data.ExpertArch3DModel = value.ExpertArch3DModel
		data.ExpertArchVirModel = value.ExpertArchVirModel
		data.ExpertArchExec = value.ExpertArchExec
		data.ExpertStruct = value.ExpertStruct
		data.ExpertStructWorkingDraw = value.ExpertStructWorkingDraw
		data.ExpertStructExec = value.ExpertStructExec
		data.ExpertInterior = value.ExpertInterior
		data.ExpertIntConceptDraw = value.ExpertIntConceptDraw
		data.ExpertIntWorkingDraw = value.ExpertIntWorkingDraw
		data.ExpertIntPhyScale = value.ExpertIntPhyScale
		data.ExpertInt3DModel = value.ExpertInt3DModel
		data.ExpertIntScaleModel = value.ExpertIntScaleModel
		data.ExpertIntVirModel = value.ExpertIntVirModel
		data.ExpertIntExec = value.ExpertIntExec
		data.ExpertLScape = value.ExpertLScape
		data.ExpertSiteSupervisor = value.ExpertSiteSupervisor
		data.ExpertSoilTest = value.ExpertSoilTest
		data.ExpertSanctions = value.ExpertSanctions
		data.SupplierConstMat = value.SupplierConstMat
		data.SupplierCement = value.SupplierCement
		data.SupplierSand = value.SupplierSand
		data.SupplierBrick = value.SupplierBrick
		data.SupplierPrecast = value.SupplierPrecast
		data.SupplierCoarse = value.SupplierCoarse
		data.SupplierConcrete = value.SupplierConcrete
		data.SupplierReinfbar = value.SupplierReinfbar
		data.SupplierTile = value.SupplierTile
		data.SupplierWproofMat = value.SupplierWproofMat
		data.SupplierShuttering = value.SupplierShuttering
		data.SupplierScaffs = value.SupplierScaffs
		data.ExpertOthConstMatSup = value.ExpertOthConstMatSup
		data.ExpertMechanical = value.ExpertMechanical
		data.ExpertHeating = value.ExpertHeating
		data.ExpertAc = value.ExpertAc
		data.ExpertElevator = value.ExpertElevator
		data.ExpertOthMechSer = value.ExpertOthMechSer
		data.ExpertElectrical = value.ExpertElectrical
		data.ExpertElecWiring = value.ExpertElecWiring
		data.ExpertElecFix = value.ExpertElecFix
		data.ExpertOthElecSer = value.ExpertOthElecSer
		data.ExpertPlumbing = value.ExpertPlumbing
		data.SupplierLabor = value.SupplierLabor
		data.SupplierCivilLabor = value.SupplierCivilLabor
		data.SupplierTileLabor = value.SupplierTileLabor
		data.SupplierPainter = value.SupplierPainter
		data.SupplierCarpenter = value.SupplierCarpenter
		data.SupplierPlumbing = value.SupplierPlumbing
		data.SupplierWaterTank = value.SupplierWaterTank
		data.SupplierWaterPipeAndFitting = value.SupplierWaterPipeAndFitting
		data.SupplierWaterFix = value.SupplierWaterFix
		data.SupplierWaterFix = value.SupplierWaterFix
		data.SupplierWiringMat = value.SupplierWiringMat
		data.SupplierSwitch = value.SupplierSwitch
		data.SupplierSensorSup = value.SupplierSensorSup
		data.SupplierEleAppSup = value.SupplierEleAppSup
		data.SupplierOthEle = value.SupplierOthEle
		data.SupplierPaint = value.SupplierPaint
		data.SupplierIntPaint = value.SupplierIntPaint
		data.SupplierExtPaint = value.SupplierExtPaint
		data.SupplierIntPrimer = value.SupplierIntPrimer
		data.SupplierExtPrimer = value.SupplierExtPrimer
		data.SupplierPutty = value.SupplierPutty
		data.SupplierPaintAccessory = value.SupplierPaintAccessory
		data.SupplierOtherPaint = value.SupplierOtherPaint
		data.ExpertFabrication = value.ExpertFabrication
		data.ExpertGlassFab = value.ExpertGlassFab
		data.ExpertWoodFab = value.ExpertWoodFab
		data.ExpertMetalAl = value.ExpertMetalAl
		data.ExpertMetalFe = value.ExpertMetalFe
		data.ExpertMetalSS = value.ExpertMetalSS
		data.ExpertOthHardware = value.ExpertOthHardware
		data.SupplierConstEqp = value.SupplierConstEqp
		data.SupplierJcb = value.SupplierJcb
		data.SupplierConcMixer = value.SupplierConcMixer
		data.SupplierHaulTruck = value.SupplierHaulTruck
		data.SupplierCrane = value.SupplierCrane
		data.SupplierOthEqp = value.SupplierOthEqp
		data.SupplierHouseProduct = value.SupplierHouseProduct
		data.SupplierFurniture = value.SupplierFurniture
		data.SupplierAppliance = value.SupplierAppliance
		data.SupplierFurnish = value.SupplierFurnish
		data.SupplierOthFin = value.SupplierOthFin
		data.ProjDesc = value.ProjDesc
		data.ProjLoc = value.ProjLoc
		data.ProjArea = value.ProjArea
		data.ProjBudget = value.ProjBudget
		data.ProjType = value.ProjType
		//data.PassingOutDate = DateString(*value.PassingOutDate)
		privateResponseData = append(privateResponseData, data)
	}
	return privateResponseData
}

// GetProfessionalInfoPublicResponseData used to get professional public info data
var GetProfessionalInfoPublicResponseData = func(professionalInfo []models.UserProfessionalInfo) (publicResponseArray []dto.GetProfessionalInfoPublicResponseData) {
	for _, value := range professionalInfo {
		data := dto.GetProfessionalInfoPublicResponseData{}
		data.ExpYearsCategory = value.ExpYearsCategory
		data.ExpertSpcCom = value.ExpertSpcCom
		data.ExpertSpcCSR = value.ExpertSpcCSR
		data.ServiceState = value.ServiceState
		data.EntityType = value.EntityType
		data.ProjLoc = value.ProjLoc
		data.ProjArea = value.ProjArea
		publicResponseArray = append(publicResponseArray, data)
	}
	return publicResponseArray
}
