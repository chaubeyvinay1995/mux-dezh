package impl

import (
	"context"
	"errors"
	"gitlab.com/umi7/DezHab_user_backend/dao/core/impl"
	"gitlab.com/umi7/DezHab_user_backend/dao/models"
	"gitlab.com/umi7/DezHab_user_backend/dao/repos"
	"time"
)

const (
	updateProfessionalInfoQuery = "id=?"
)

var ProfessionalInfo repos.ProfessionalInfo

type ProfessionalInfoImpl struct{}

func init() {
	ProfessionalInfo = ProfessionalInfoImpl{}
}

func (professionalInfoImpl ProfessionalInfoImpl) Search(ctx context.Context,
	userId string) (professionalInfo []models.UserProfessionalInfo, err error) {
	var professionalList []models.UserProfessionalInfo
	if err = impl.Read.Find(&professionalList, models.UserProfessionalInfo{UserId: userId},
		&professionalList); err != nil {
		return
	}
	if len(professionalList) == 0 {
		err = errors.New(ErrNoData)
		return
	}
	professionalInfo = professionalList
	return
}

func (professionalInfoImpl ProfessionalInfoImpl) Create(ctx context.Context,
	professionalInfo models.UserProfessionalInfo) (err error) {
	professionalInfo.CreatedAt = time.Now().UTC()
	professionalInfo.UpdatedAt = professionalInfo.CreatedAt
	err = impl.Write.Create(&professionalInfo)
	return
}

func (professionalInfoImpl ProfessionalInfoImpl) Delete(ctx context.Context, professionalInfo models.UserProfessionalInfo) (err error) {
	err = impl.Write.Delete(&professionalInfo)
	return
}

func (professionalInfoImpl ProfessionalInfoImpl) Update(ctx context.Context, professionalInfo models.UserProfessionalInfo) (err error) {
	professionalInfo.UpdatedAt = time.Now().UTC()
	updateValues := map[string]interface{}{
		"user_alias": professionalInfo.UserAlias, "entity_type": professionalInfo.EntityType,
		"contact_hrs_from": professionalInfo.ContactHrsFrom, "contact_min_from": professionalInfo.ContactMinFrom,
		"contact_hrs_to": professionalInfo.ContactHrsTo, "contact_min_to": professionalInfo.ContactMinTo,
		"contact_person_name": professionalInfo.ContactPersonName,
		"exp_years_category":  professionalInfo.ExpYearsCategory,
		"expert_profile_desc": professionalInfo.ExpertProfileDesc, "service_state": professionalInfo.ServiceState,
		"inst_name": professionalInfo.InstName, "inst_degree": professionalInfo.InstDegree,
		"inst_degree_url": professionalInfo.InstDegreeUrl, "ed_level": professionalInfo.EdLevel,
		"cert_name": professionalInfo.CertName, "cert_inst": professionalInfo.CertInst,
		"cert_url": professionalInfo.CertUrl, "achievement_name": professionalInfo.AchievementName,
		"achievement_description":  professionalInfo.AchievementDescription,
		"achievement_proof_urls":   professionalInfo.AchievementProofUrls,
		"expert_oth_const_mat_sup": professionalInfo.ExpertOthConstMatSup,
		"expert_oth_mech_ser":      professionalInfo.ExpertOthMechSer,
		"supplier_tile_labor":      professionalInfo.SupplierTileLabor,
		"expert_oth_hardware":      professionalInfo.ExpertOthHardware,
		"supplier_oth_eqp":         professionalInfo.SupplierOthEqp, "supplier_oth_fin": professionalInfo.SupplierOthFin,
		"proj_desc": professionalInfo.ProjDesc, "proj_loc": professionalInfo.ProjLoc,
		"proj_area": professionalInfo.ProjArea, "proj_budget": professionalInfo.ProjBudget,
		"proj_type": professionalInfo.ProjType, "passing_out_date": professionalInfo.PassingOutDate,
		"expert_spc_res": professionalInfo.ExpertSpcRes, "expert_spc_com": professionalInfo.ExpertSpcCom,
		"expert_spc_inst": professionalInfo.ExpertSpcInst, "expert_spc_health": professionalInfo.ExpertSpcHealth,
		"expert_spc_infra": professionalInfo.ExpertSpcInfra, "expert_spc_csr": professionalInfo.ExpertSpcCSR,
		"expert_arch": professionalInfo.ExpertArch, "expert_arch_concept_draw": professionalInfo.ExpertArchConceptDraw,
		"expert_arch_working_draw": professionalInfo.ExpertArchWorkingDraw,
		"expert_arch_3d_render":    professionalInfo.ExpertArch3DRender,
		"expert_arch_scale_model":  professionalInfo.ExpertArchScaleModel,
		"expert_arch_3d_model":     professionalInfo.ExpertArch3DModel,
		"expert_arch_vir_model":    professionalInfo.ExpertArchVirModel,
		"expert_arch_exec":         professionalInfo.ExpertArchExec, "expert_struct": professionalInfo.ExpertStruct,
		"expert_struct_working_draw": professionalInfo.ExpertStructWorkingDraw,
		"expert_struct_exec":         professionalInfo.ExpertStructExec,
		"expert_interior":            professionalInfo.ExpertInterior,
		"expert_int_conceptDraw":     professionalInfo.ExpertIntConceptDraw,
		"expert_int_working_draw":    professionalInfo.ExpertIntWorkingDraw,
		"expert_int_phy_scale":       professionalInfo.ExpertIntPhyScale,
		"expert_int_3d_model":        professionalInfo.ExpertInt3DModel,
		"expert_int_scale_model":     professionalInfo.ExpertIntScaleModel,
		"expert_int_vir_model":       professionalInfo.ExpertIntVirModel,
		"expert_int_exec":            professionalInfo.ExpertIntExec,
		"expert_l_scape":             professionalInfo.ExpertLScape,
		"expert_site_supervisor":     professionalInfo.ExpertSiteSupervisor,
		"expert_soil_test":           professionalInfo.ExpertSoilTest,
		"expert_sanctions":           professionalInfo.ExpertSanctions,
		"supplier_const_mat":         professionalInfo.SupplierConstMat,
		"supplier_cement":            professionalInfo.SupplierCement,
		"supplier_sand":              professionalInfo.SupplierSand,
		"supplier_brick":             professionalInfo.SupplierBrick,
		"supplier_precast":           professionalInfo.SupplierPrecast,
		"supplier_coarse":            professionalInfo.SupplierCoarse,
		"supplier_concrete":          professionalInfo.SupplierConcrete,
		"supplier_reinfbar":          professionalInfo.SupplierReinfbar, "supplier_tile": professionalInfo.SupplierTile,
		"supplier_wproof_mat": professionalInfo.SupplierWproofMat,
		"supplier_shuttering": professionalInfo.SupplierShuttering, "supplier_scaffs": professionalInfo.SupplierScaffs,
		"expert_mechanical": professionalInfo.ExpertMechanical, "expert_heating": professionalInfo.ExpertHeating,
		"expert_ac": professionalInfo.ExpertAc, "expert_elevator": professionalInfo.ExpertElevator,
		"expert_electrical": professionalInfo.ExpertElectrical, "expert_elec_wiring": professionalInfo.ExpertElecWiring,
		"expert_elec_fix": professionalInfo.ExpertElecFix, "expert_oth_elec_ser": professionalInfo.ExpertOthElecSer,
		"expert_plumbing": professionalInfo.ExpertPlumbing, "supplier_labor": professionalInfo.SupplierLabor,
		"supplier_civil_labor":            professionalInfo.SupplierCivilLabor,
		"supplier_painter":                professionalInfo.SupplierPainter,
		"supplier_carpenter":              professionalInfo.SupplierCarpenter,
		"supplier_plumbing":               professionalInfo.SupplierPlumbing,
		"supplier_water_tank":             professionalInfo.SupplierWaterTank,
		"supplier_water_pipe_and_fitting": professionalInfo.SupplierWaterPipeAndFitting,
		"supplier_water_fix":              professionalInfo.SupplierWaterFix,
		"supplier_elec":                   professionalInfo.SupplierElec,
		"supplier_wiring_mat":             professionalInfo.SupplierWiringMat,
		"supplier_switch":                 professionalInfo.SupplierSwitch,
		"supplier_sensor_sup":             professionalInfo.SupplierSensorSup,
		"supplier_ele_app_sup":            professionalInfo.SupplierEleAppSup,
		"supplier_oth_ele":                professionalInfo.SupplierOthEle,
		"supplier_paint":                  professionalInfo.SupplierPaint,
		"supplier_int_paint":              professionalInfo.SupplierIntPaint,
		"supplier_ext_paint":              professionalInfo.SupplierExtPaint,
		"supplier_int_primer":             professionalInfo.SupplierIntPrimer,
		"supplier_ext_primer":             professionalInfo.SupplierExtPrimer,
		"supplier_putty":                  professionalInfo.SupplierPutty,
		"supplier_paint_accessory":        professionalInfo.SupplierPaintAccessory,
		"supplier_other_paint":            professionalInfo.SupplierOtherPaint,
		"expert_fabrication":              professionalInfo.ExpertFabrication,
		"expert_glass_fab":                professionalInfo.ExpertGlassFab,
		"expert_wood_fab":                 professionalInfo.ExpertWoodFab,
		"expert_metal_al":                 professionalInfo.ExpertMetalAl,
		"expert_metal_fe":                 professionalInfo.ExpertMetalFe,
		"expert_metal_ss":                 professionalInfo.ExpertMetalSS,
		"supplier_const_eqp":              professionalInfo.SupplierConstEqp,
		"supplier_jcb":                    professionalInfo.SupplierJcb,
		"supplier_conc_mixer":             professionalInfo.SupplierConcMixer,
		"supplier_haul_truck":             professionalInfo.SupplierHaulTruck,
		"supplier_crane":                  professionalInfo.SupplierCrane,
		"supplier_house_product":          professionalInfo.SupplierHouseProduct,
		"supplier_furniture":              professionalInfo.SupplierFurniture,
		"supplier_appliance":              professionalInfo.SupplierAppliance,
		"supplier_furnish":                professionalInfo.SupplierFurnish,
	}
	err = impl.Write.Update(&professionalInfo, updateProfessionalInfoQuery, professionalInfo.ID, updateValues)
	return
}
