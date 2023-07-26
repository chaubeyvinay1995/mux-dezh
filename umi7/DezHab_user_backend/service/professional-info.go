package service

import (
	"context"
	"gitlab.com/umi7/DezHab_user_backend/dao/models"
	"time"
)

type ProfessionalInfo interface {
	Get(ctx context.Context, id string) ([]models.UserProfessionalInfo, error)

	Create(ctx context.Context, userId, userAlias, entityType, contactHrsFrom, contactMinFrom, contactHrsTo,
		contactMinTo, contactPersonName, expYearsCategory, expertProfileDesc, serviceState, instName, instDegree,
		instDegreeUrl, edLevel, certName, certInst, certUrl, achievementName, achievementDescription,
		achievementProofUrls, expertOthConstMatSup, expertOthMechSer, supplierTileLabor, expertOthHardware,
		supplierOthEqp, supplierOthFin, projDesc, projLoc, projArea, projBudget, projType string,
		passingOutDate *time.Time, expertSpcRes, expertSpcCom, expertSpcInst, expertSpcHealth, expertSpcInfra,
		expertSpcCSR, expertArch, expertArchConceptDraw, expertArchWorkingDraw, expertArch3DRender,
		expertArchScaleModel, expertArch3DModel, expertArchVirModel, expertArchExec, expertStruct,
		expertStructWorkingDraw, expertStructExec, expertInterior, expertIntConceptDraw, expertIntWorkingDraw,
		expertIntPhyScale, expertInt3DModel, expertIntScaleModel, expertIntVirModel, expertIntExec, expertLScape,
		expertSiteSupervisor, expertSoilTest, expertSanctions, supplierConstMat, supplierCement, supplierSand,
		supplierBrick, supplierPrecast, supplierCoarse, supplierConcrete, supplierReinfbar, supplierTile,
		supplierWproofMat, supplierShuttering, supplierScaffs, expertMechanical, expertHeating, expertAc,
		expertElevator, expertElectrical, expertElecWiring, expertElecFix, expertOthElecSer, expertPlumbing,
		supplierLabor, supplierCivilLabor, supplierPainter, supplierCarpenter, supplierPlumbing, supplierWaterTank,
		supplierWaterPipeAndFitting, supplierWaterFix, supplierElec, supplierWiringMat, supplierSwitch,
		supplierSensorSup, supplierEleAppSup, supplierOthEle, supplierPaint, supplierIntPaint, supplierExtPaint,
		supplierIntPrimer, supplierExtPrimer, supplierPutty, supplierPaintAccessory, supplierOtherPaint,
		expertFabrication, expertGlassFab, expertWoodFab, expertMetalAl, expertMetalFe, expertMetalSS,
		supplierConstEqp, supplierJcb, supplierConcMixer, supplierHaulTruck, supplierCrane, supplierHouseProduct,
		supplierFurniture, supplierAppliance, supplierFurnish bool) (id string, err error)

	Update(ctx context.Context, id, userAlias, entityType, contactHrsFrom, contactMinFrom, contactHrsTo,
		contactMinTo, contactPersonName, expYearsCategory, expertProfileDesc, serviceState, instName, instDegree,
		instDegreeUrl, edLevel, certName, certInst, certUrl, achievementName, achievementDescription,
		achievementProofUrls, expertOthConstMatSup, expertOthMechSer, supplierTileLabor, expertOthHardware,
		supplierOthEqp, supplierOthFin, projDesc, projLoc, projArea, projBudget, projType string,
		passingOutDate time.Time, expertSpcRes, expertSpcCom, expertSpcInst, expertSpcHealth, expertSpcInfra,
		expertSpcCSR, expertArch, expertArchConceptDraw, expertArchWorkingDraw, expertArch3DRender,
		expertArchScaleModel, expertArch3DModel, expertArchVirModel, expertArchExec, expertStruct,
		expertStructWorkingDraw, expertStructExec, expertInterior, expertIntConceptDraw, expertIntWorkingDraw,
		expertIntPhyScale, expertInt3DModel, expertIntScaleModel, expertIntVirModel, expertIntExec, expertLScape,
		expertSiteSupervisor, expertSoilTest, expertSanctions, supplierConstMat, supplierCement, supplierSand,
		supplierBrick, supplierPrecast, supplierCoarse, supplierConcrete, supplierReinfbar, supplierTile,
		supplierWproofMat, supplierShuttering, supplierScaffs, expertMechanical, expertHeating, expertAc,
		expertElevator, expertElectrical, expertElecWiring, expertElecFix, expertOthElecSer, expertPlumbing,
		supplierLabor, supplierCivilLabor, supplierPainter, supplierCarpenter, supplierPlumbing, supplierWaterTank,
		supplierWaterPipeAndFitting, supplierWaterFix, supplierElec, supplierWiringMat, supplierSwitch,
		supplierSensorSup, supplierEleAppSup, supplierOthEle, supplierPaint, supplierIntPaint, supplierExtPaint,
		supplierIntPrimer, supplierExtPrimer, supplierPutty, supplierPaintAccessory, supplierOtherPaint,
		expertFabrication, expertGlassFab, expertWoodFab, expertMetalAl, expertMetalFe, expertMetalSS,
		supplierConstEqp, supplierJcb, supplierConcMixer, supplierHaulTruck, supplierCrane, supplierHouseProduct,
		supplierFurniture, supplierAppliance, supplierFurnish bool) (err error)

	Delete(ctx context.Context, id string) error
}
