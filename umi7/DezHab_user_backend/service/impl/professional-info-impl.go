package impl

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"gitlab.com/umi7/DezHab_user_backend/common/logger"
	"gitlab.com/umi7/DezHab_user_backend/dao/models"
	"gitlab.com/umi7/DezHab_user_backend/dao/repos/impl"
	"gitlab.com/umi7/DezHab_user_backend/service"
	"time"
)

var ProfessionalInfo service.ProfessionalInfo

type ProfessionalInfoImpl struct{}

func init() {
	ProfessionalInfo = ProfessionalInfoImpl{}
}

func (p ProfessionalInfoImpl) Get(ctx context.Context, id string) (professionalInfo []models.UserProfessionalInfo, err error) {
	professionalInfo, err = impl.ProfessionalInfo.Search(ctx, id)
	if err != nil {
		logger.Error(ctx, err)
		return
	}
	return

}

func (p ProfessionalInfoImpl) Create(ctx context.Context, userId, userAlias, entityType, contactHrsFrom,
	contactMinFrom, contactHrsTo, contactMinTo, contactPersonName, expYearsCategory, expertProfileDesc, serviceState,
	instName, instDegree, instDegreeUrl, edLevel, certName, certInst, certUrl, achievementName, achievementDescription,
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
	supplierFurniture, supplierAppliance, supplierFurnish bool) (id string, err error) {
	id = uuid.New().String()
	professionalInfo := models.UserProfessionalInfo{
		ID: id, UserId: userId, UserAlias: userAlias, EntityType: entityType, ContactHrsFrom: contactHrsFrom,
		ContactMinFrom: contactMinFrom, ContactHrsTo: contactHrsTo, ContactMinTo: contactMinTo,
		ContactPersonName: contactPersonName, ExpYearsCategory: expYearsCategory, ExpertProfileDesc: expertProfileDesc,
		ServiceState: serviceState, InstName: instName, InstDegree: instDegree, InstDegreeUrl: instDegreeUrl,
		EdLevel: edLevel, CertName: certName, CertInst: certInst, CertUrl: certUrl, AchievementName: achievementName,
		AchievementDescription: achievementDescription, AchievementProofUrls: achievementProofUrls,
		ExpertOthConstMatSup: expertOthConstMatSup, ExpertOthMechSer: expertOthMechSer,
		SupplierTileLabor: supplierTileLabor, ExpertOthHardware: expertOthHardware, SupplierOthEqp: supplierOthEqp,
		SupplierOthFin: supplierOthFin, ProjDesc: projDesc, ProjLoc: projLoc, ProjArea: projArea,
		ProjBudget: projBudget, ProjType: projType, PassingOutDate: passingOutDate, ExpertSpcRes: expertSpcRes,
		ExpertSpcCom: expertSpcCom, ExpertSpcInst: expertSpcInst, ExpertSpcHealth: expertSpcHealth,
		ExpertSpcInfra: expertSpcInfra, ExpertSpcCSR: expertSpcCSR, ExpertArch: expertArch,
		ExpertArchConceptDraw: expertArchConceptDraw, ExpertArchWorkingDraw: expertArchWorkingDraw,
		ExpertArch3DRender: expertArch3DRender, ExpertArchScaleModel: expertArchScaleModel,
		ExpertArch3DModel: expertArch3DModel, ExpertArchVirModel: expertArchVirModel, ExpertArchExec: expertArchExec,
		ExpertStruct: expertStruct, ExpertStructWorkingDraw: expertStructWorkingDraw,
		ExpertStructExec: expertStructExec, ExpertInterior: expertInterior, ExpertIntConceptDraw: expertIntConceptDraw,
		ExpertIntWorkingDraw: expertIntWorkingDraw, ExpertIntPhyScale: expertIntPhyScale,
		ExpertInt3DModel: expertInt3DModel, ExpertIntScaleModel: expertIntScaleModel,
		ExpertIntVirModel: expertIntVirModel, ExpertIntExec: expertIntExec, ExpertLScape: expertLScape,
		ExpertSiteSupervisor: expertSiteSupervisor, ExpertSoilTest: expertSoilTest, ExpertSanctions: expertSanctions,
		SupplierConstMat: supplierConstMat, SupplierCement: supplierCement, SupplierSand: supplierSand,
		SupplierBrick: supplierBrick, SupplierPrecast: supplierPrecast, SupplierCoarse: supplierCoarse,
		SupplierConcrete: supplierConcrete, SupplierReinfbar: supplierReinfbar, SupplierTile: supplierTile,
		SupplierWproofMat: supplierWproofMat, SupplierShuttering: supplierShuttering, SupplierScaffs: supplierScaffs,
		ExpertMechanical: expertMechanical, ExpertHeating: expertHeating, ExpertAc: expertAc,
		ExpertElevator: expertElevator, ExpertElectrical: expertElectrical, ExpertElecWiring: expertElecWiring,
		ExpertElecFix: expertElecFix, ExpertOthElecSer: expertOthElecSer, ExpertPlumbing: expertPlumbing,
		SupplierLabor: supplierLabor, SupplierCivilLabor: supplierCivilLabor, SupplierPainter: supplierPainter,
		SupplierCarpenter: supplierCarpenter, SupplierPlumbing: supplierPlumbing,
		SupplierWaterTank: supplierWaterTank, SupplierWaterPipeAndFitting: supplierWaterPipeAndFitting,
		SupplierWaterFix: supplierWaterFix, SupplierElec: supplierElec, SupplierWiringMat: supplierWiringMat,
		SupplierSwitch: supplierSwitch, SupplierSensorSup: supplierSensorSup, SupplierEleAppSup: supplierEleAppSup,
		SupplierOthEle: supplierOthEle, SupplierPaint: supplierPaint, SupplierIntPaint: supplierIntPaint,
		SupplierExtPaint: supplierExtPaint, SupplierIntPrimer: supplierIntPrimer, SupplierExtPrimer: supplierExtPrimer,
		SupplierPutty: supplierPutty, SupplierPaintAccessory: supplierPaintAccessory,
		SupplierOtherPaint: supplierOtherPaint, ExpertFabrication: expertFabrication, ExpertGlassFab: expertGlassFab,
		ExpertWoodFab: expertWoodFab, ExpertMetalAl: expertMetalAl, ExpertMetalFe: expertMetalFe,
		ExpertMetalSS: expertMetalSS, SupplierConstEqp: supplierConstEqp, SupplierJcb: supplierJcb,
		SupplierConcMixer: supplierConcMixer, SupplierHaulTruck: supplierHaulTruck, SupplierCrane: supplierCrane,
		SupplierHouseProduct: supplierHouseProduct, SupplierFurniture: supplierFurniture,
		SupplierAppliance: supplierAppliance, SupplierFurnish: supplierFurnish,
	}
	err = impl.ProfessionalInfo.Create(ctx, professionalInfo)
	if err != nil {
		logger.Error(ctx, fmt.Sprintf("Error occurred while creating entry on professional info : %v", err))
	}
	return

}

func (p ProfessionalInfoImpl) Update(ctx context.Context, id, userAlias, entityType, contactHrsFrom, contactMinFrom,
	contactHrsTo, contactMinTo, contactPersonName, expYearsCategory, expertProfileDesc, serviceState, instName,
	instDegree, instDegreeUrl, edLevel, certName, certInst, certUrl, achievementName, achievementDescription,
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
	supplierFurniture, supplierAppliance, supplierFurnish bool) (err error) {
	professionalInfo := models.UserProfessionalInfo{
		ID: id, UserAlias: userAlias, EntityType: entityType, ContactHrsFrom: contactHrsFrom,
		ContactMinFrom: contactMinFrom, ContactHrsTo: contactHrsTo, ContactMinTo: contactMinTo,
		ContactPersonName: contactPersonName, ExpYearsCategory: expYearsCategory, ExpertProfileDesc: expertProfileDesc,
		ServiceState: serviceState, InstName: instName, InstDegree: instDegree, InstDegreeUrl: instDegreeUrl,
		EdLevel: edLevel, CertName: certName, CertInst: certInst, CertUrl: certUrl, AchievementName: achievementName,
		AchievementDescription: achievementDescription, AchievementProofUrls: achievementProofUrls,
		ExpertOthConstMatSup: expertOthConstMatSup, ExpertOthMechSer: expertOthMechSer,
		SupplierTileLabor: supplierTileLabor, ExpertOthHardware: expertOthHardware, SupplierOthEqp: supplierOthEqp,
		SupplierOthFin: supplierOthFin, ProjDesc: projDesc, ProjLoc: projLoc, ProjArea: projArea,
		ProjBudget: projBudget, ProjType: projType, PassingOutDate: &passingOutDate, ExpertSpcRes: expertSpcRes,
		ExpertSpcCom: expertSpcCom, ExpertSpcInst: expertSpcInst, ExpertSpcHealth: expertSpcHealth,
		ExpertSpcInfra: expertSpcInfra, ExpertSpcCSR: expertSpcCSR, ExpertArch: expertArch,
		ExpertArchConceptDraw: expertArchConceptDraw, ExpertArchWorkingDraw: expertArchWorkingDraw,
		ExpertArch3DRender: expertArch3DRender, ExpertArchScaleModel: expertArchScaleModel,
		ExpertArch3DModel: expertArch3DModel, ExpertArchVirModel: expertArchVirModel, ExpertArchExec: expertArchExec,
		ExpertStruct: expertStruct, ExpertStructWorkingDraw: expertStructWorkingDraw,
		ExpertStructExec: expertStructExec, ExpertInterior: expertInterior, ExpertIntConceptDraw: expertIntConceptDraw,
		ExpertIntWorkingDraw: expertIntWorkingDraw, ExpertIntPhyScale: expertIntPhyScale,
		ExpertInt3DModel: expertInt3DModel, ExpertIntScaleModel: expertIntScaleModel,
		ExpertIntVirModel: expertIntVirModel, ExpertIntExec: expertIntExec, ExpertLScape: expertLScape,
		ExpertSiteSupervisor: expertSiteSupervisor, ExpertSoilTest: expertSoilTest, ExpertSanctions: expertSanctions,
		SupplierConstMat: supplierConstMat, SupplierCement: supplierCement, SupplierSand: supplierSand,
		SupplierBrick: supplierBrick, SupplierPrecast: supplierPrecast, SupplierCoarse: supplierCoarse,
		SupplierConcrete: supplierConcrete, SupplierReinfbar: supplierReinfbar, SupplierTile: supplierTile,
		SupplierWproofMat: supplierWproofMat, SupplierShuttering: supplierShuttering, SupplierScaffs: supplierScaffs,
		ExpertMechanical: expertMechanical, ExpertHeating: expertHeating, ExpertAc: expertAc,
		ExpertElevator: expertElevator, ExpertElectrical: expertElectrical, ExpertElecWiring: expertElecWiring,
		ExpertElecFix: expertElecFix, ExpertOthElecSer: expertOthElecSer, ExpertPlumbing: expertPlumbing,
		SupplierLabor: supplierLabor, SupplierCivilLabor: supplierCivilLabor, SupplierPainter: supplierPainter,
		SupplierCarpenter: supplierCarpenter, SupplierPlumbing: supplierPlumbing,
		SupplierWaterTank: supplierWaterTank, SupplierWaterPipeAndFitting: supplierWaterPipeAndFitting,
		SupplierWaterFix: supplierWaterFix, SupplierElec: supplierElec, SupplierWiringMat: supplierWiringMat,
		SupplierSwitch: supplierSwitch, SupplierSensorSup: supplierSensorSup, SupplierEleAppSup: supplierEleAppSup,
		SupplierOthEle: supplierOthEle, SupplierPaint: supplierPaint, SupplierIntPaint: supplierIntPaint,
		SupplierExtPaint: supplierExtPaint, SupplierIntPrimer: supplierIntPrimer, SupplierExtPrimer: supplierExtPrimer,
		SupplierPutty: supplierPutty, SupplierPaintAccessory: supplierPaintAccessory,
		SupplierOtherPaint: supplierOtherPaint, ExpertFabrication: expertFabrication, ExpertGlassFab: expertGlassFab,
		ExpertWoodFab: expertWoodFab, ExpertMetalAl: expertMetalAl, ExpertMetalFe: expertMetalFe,
		ExpertMetalSS: expertMetalSS, SupplierConstEqp: supplierConstEqp, SupplierJcb: supplierJcb,
		SupplierConcMixer: supplierConcMixer, SupplierHaulTruck: supplierHaulTruck, SupplierCrane: supplierCrane,
		SupplierHouseProduct: supplierHouseProduct, SupplierFurniture: supplierFurniture,
		SupplierAppliance: supplierAppliance, SupplierFurnish: supplierFurnish,
	}
	err = impl.ProfessionalInfo.Update(ctx, professionalInfo)
	if err != nil {
		logger.Error(ctx, fmt.Sprintf("Error occurred while creating entry on professional info : %v", err))
	}
	return

}

func (p ProfessionalInfoImpl) Delete(ctx context.Context, id string) (err error) {
	professionalInfo := models.UserProfessionalInfo{
		ID: id,
	}
	err = impl.ProfessionalInfo.Delete(ctx, professionalInfo)
	if err != nil {
		logger.Error(ctx, fmt.Sprintf("Error occurred while creating entry on professional info : %v", err))
	}
	return

}
