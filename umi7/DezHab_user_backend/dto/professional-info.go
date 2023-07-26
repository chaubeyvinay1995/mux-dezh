package dto

import (
	"gitlab.com/umi7/DezHab_user_backend/common"
)

// GetUserProfessionalInfoRequest used to validate the get user professional info request
type GetUserProfessionalInfoRequest struct {
	ID string `json:"id" validate:"required"`
}

// Validate function used to validate the request.
// This is the implementation of the Validation interface.

func (g GetUserProfessionalInfoRequest) Validate() error {
	return common.Validate.Struct(g)
}

// GetPublicProfessionalInfoResponse response struct
type GetPublicProfessionalInfoResponse struct {
	Response GetPublicProfessionalInfoData `json:"response"`
}

type GetPublicProfessionalInfoData struct {
	CompositeScore uint `json:"compositeScore"`
}

// GetProfessionalInfoPrivateResponse struct used to send get academic detail response.
type GetProfessionalInfoPrivateResponse struct {
	CompositeScore uint                                     `json:"compositeScore"`
	Response       []GetProfessionalInfoPrivateResponseData `json:"response"`
}

type GetProfessionalInfoPrivateResponseData struct {
	Id                          string `json:"id"`
	UserId                      string `json:"userId"`
	UserAlias                   string `json:"userAlias"`
	EntityType                  string `json:"entityType"`
	ContactHrsFrom              string `json:"contactHrsFrom"`
	ContactMinFrom              string `json:"contactMinFrom"`
	ContactHrsTo                string `json:"contactHrsTo"`
	ContactMinTo                string `json:"contactMinTo"`
	ContactPersonName           string `json:"contactPersonName"`
	ExpYearsCategory            string `json:"expYearsCategory"`
	ExpertProfileDesc           string `json:"expertProfileDesc"`
	ServiceState                string `json:"serviceState"`
	InstName                    string `json:"instName"`
	InstDegree                  string `json:"instDegree"`
	InstDegreeUrl               string `json:"instDegreeUrl"`
	EdLevel                     string `json:"edLevel"`
	CertName                    string `json:"certName"`
	CertInst                    string `json:"certInst"`
	CertUrl                     string `json:"certUrl"`
	AchievementName             string `json:"achievementName"`
	AchievementDescription      string `json:"achievementDescription"`
	AchievementProofUrls        string `json:"achievementProofUrls"`
	ExpertOthConstMatSup        string `json:"expertOthConstMatSup"`
	ExpertOthMechSer            string `json:"expertOthMechSer"`
	SupplierTileLabor           string `json:"supplierTileLabor"`
	ExpertOthHardware           string `json:"expertOthHardware"`
	SupplierOthEqp              string `json:"supplierOthEqp"`
	SupplierOthFin              string `json:"supplierOthFin"`
	ProjDesc                    string `json:"projDesc"`
	ProjLoc                     string `json:"projLoc"`
	ProjArea                    string `json:"projArea"`
	ProjBudget                  string `json:"projBudget"`
	ProjType                    string `json:"projType"`
	PassingOutDate              string `json:"passingOutDate"`
	ExpertSpcRes                bool   `json:"expertSpcRes"`
	ExpertSpcCom                bool   `json:"expertSpcCom"`
	ExpertSpcInst               bool   `json:"expertSpcInst"`
	ExpertSpcHealth             bool   `json:"expertSpcHealth"`
	ExpertSpcInfra              bool   `json:"expertSpcInfra"`
	ExpertSpcCSR                bool   `json:"expertSpcCSR"`
	ExpertArch                  bool   `json:"expertArch"`
	ExpertArchConceptDraw       bool   `json:"expertArchConceptDraw"`
	ExpertArchWorkingDraw       bool   `json:"expertArchWorkingDraw"`
	ExpertArch3DRender          bool   `json:"expertArch3DRender"`
	ExpertArchScaleModel        bool   `json:"expertArchScaleModel"`
	ExpertArch3DModel           bool   `json:"expertArch3DModel"`
	ExpertArchVirModel          bool   `json:"expertArchVirModel"`
	ExpertArchExec              bool   `json:"expertArchExec"`
	ExpertStruct                bool   `json:"expertStruct"`
	ExpertStructWorkingDraw     bool   `json:"expertStructWorkingDraw"`
	ExpertStructExec            bool   `json:"expertStructExec"`
	ExpertInterior              bool   `json:"expertInterior"`
	ExpertIntConceptDraw        bool   `json:"expertIntConceptDraw"`
	ExpertIntWorkingDraw        bool   `json:"expertIntWorkingDraw"`
	ExpertIntPhyScale           bool   `json:"expertIntPhyScale"`
	ExpertInt3DModel            bool   `json:"expertInt3DModel"`
	ExpertIntScaleModel         bool   `json:"expertIntScaleModel"`
	ExpertIntVirModel           bool   `json:"expertIntVirModel"`
	ExpertIntExec               bool   `json:"expertIntExec"`
	ExpertLScape                bool   `json:"expertLScape"`
	ExpertSiteSupervisor        bool   `json:"expertSiteSupervisor"`
	ExpertSoilTest              bool   `json:"expertSoilTest"`
	ExpertSanctions             bool   `json:"expertSanctions"`
	SupplierConstMat            bool   `json:"supplierConstMat"`
	SupplierCement              bool   `json:"supplierCement"`
	SupplierSand                bool   `json:"supplierSand"`
	SupplierBrick               bool   `json:"supplierBrick"`
	SupplierPrecast             bool   `json:"supplierPrecast"`
	SupplierCoarse              bool   `json:"supplierCoarse"`
	SupplierConcrete            bool   `json:"supplierConcrete"`
	SupplierReinfbar            bool   `json:"supplierReinfbar"`
	SupplierTile                bool   `json:"supplierTile"`
	SupplierWproofMat           bool   `json:"supplierWproofMat"`
	SupplierShuttering          bool   `json:"supplierShuttering"`
	SupplierScaffs              bool   `json:"supplierScaffs"`
	ExpertMechanical            bool   `json:"expertMechanical"`
	ExpertHeating               bool   `json:"expertHeating"`
	ExpertAc                    bool   `json:"expertAc"`
	ExpertElevator              bool   `json:"expertElevator"`
	ExpertElectrical            bool   `json:"expertElectrical"`
	ExpertElecWiring            bool   `json:"expertElecWiring"`
	ExpertElecFix               bool   `json:"expertElecFix"`
	ExpertOthElecSer            bool   `json:"expertOthElecSer"`
	ExpertPlumbing              bool   `json:"expertPlumbing"`
	SupplierLabor               bool   `json:"supplierLabor"`
	SupplierCivilLabor          bool   `json:"supplierCivilLabor"`
	SupplierPainter             bool   `json:"supplierPainter"`
	SupplierCarpenter           bool   `json:"supplierCarpenter"`
	SupplierPlumbing            bool   `json:"supplierPlumbing"`
	SupplierWaterTank           bool   `json:"supplierWaterTank"`
	SupplierWaterPipeAndFitting bool   `json:"supplierWaterPipeAndFitting"`
	SupplierWaterFix            bool   `json:"supplierWaterFix"`
	SupplierElec                bool   `json:"supplierElec"`
	SupplierWiringMat           bool   `json:"supplierWiringMat"`
	SupplierSwitch              bool   `json:"supplierSwitch"`
	SupplierSensorSup           bool   `json:"supplierSensorSup"`
	SupplierEleAppSup           bool   `json:"supplierEleAppSup"`
	SupplierOthEle              bool   `json:"supplierOthEle"`
	SupplierPaint               bool   `json:"supplierPaint"`
	SupplierIntPaint            bool   `json:"supplierIntPaint"`
	SupplierExtPaint            bool   `json:"supplierExtPaint"`
	SupplierIntPrimer           bool   `json:"supplierIntPrimer"`
	SupplierExtPrimer           bool   `json:"supplierExtPrimer"`
	SupplierPutty               bool   `json:"supplierPutty"`
	SupplierPaintAccessory      bool   `json:"supplierPaintAccessory"`
	SupplierOtherPaint          bool   `json:"supplierOtherPaint"`
	ExpertFabrication           bool   `json:"expertFabrication"`
	ExpertGlassFab              bool   `json:"expertGlassFab"`
	ExpertWoodFab               bool   `json:"expertWoodFab"`
	ExpertMetalAl               bool   `json:"expertMetalAl"`
	ExpertMetalFe               bool   `json:"expertMetalFe"`
	ExpertMetalSS               bool   `json:"expertMetalSS"`
	SupplierConstEqp            bool   `json:"supplierConstEqp"`
	SupplierJcb                 bool   `json:"supplierJcb"`
	SupplierConcMixer           bool   `json:"supplierConcMixer"`
	SupplierHaulTruck           bool   `json:"supplierHaulTruck"`
	SupplierCrane               bool   `json:"supplierCrane"`
	SupplierHouseProduct        bool   `json:"supplierHouseProduct"`
	SupplierFurniture           bool   `json:"supplierFurniture"`
	SupplierAppliance           bool   `json:"supplierAppliance"`
	SupplierFurnish             bool   `json:"supplierFurnish"`
}

// GetProfessionalInfoPublicResponse struct used to send get academic detail response.
type GetProfessionalInfoPublicResponse struct {
	CompositeScore uint                                    `json:"compositeScore"`
	Response       []GetProfessionalInfoPublicResponseData `json:"response"`
}

type GetProfessionalInfoPublicResponseData struct {
	ExpYearsCategory string `json:"expYearsCategory"`
	ExpertSpcCom     bool   `json:"expertSpcCom"`
	ExpertSpcCSR     bool   `json:"expertSpcCSR"`
	ServiceState     string `json:"serviceState"`
	ProjLoc          string `json:"projLoc"`
	ProjArea         string `json:"projArea"`
	EntityType       string `json:"entityType"`
}

// DeleteProfessionalInfoRequest struct used to validate the delete request for Professional.
type DeleteProfessionalInfoRequest struct {
	Id string `json:"id" validate:"required"`
}

// Validate function used to validate the request.
// This is the implementation of the Validation interface.

func (g DeleteProfessionalInfoRequest) Validate() error {
	return common.Validate.Struct(g)
}

type DeleteProfessionalInfoResponse struct {
	Response DeleteProfessionalInfoMessage `json:"response"`
}

type DeleteProfessionalInfoMessage struct {
	Message string `json:"message"`
}

// CreateProfessionalInfoRequest struct used to validate the Create Professional Info Request
type CreateProfessionalInfoRequest struct {
	UserId                      string `json:"userId" validate:"min=1,max=50,required"`
	UserAlias                   string `json:"userAlias" validate:"min=1,max=50,required"`
	EntityType                  string `json:"entityType" validate:"min=1,max=50,required"`
	ContactHrsFrom              string `json:"contactHrsFrom" validate:"min=1,max=50,required"`
	ContactMinFrom              string `json:"contactMinFrom" validate:"min=1,max=50,required"`
	ContactHrsTo                string `json:"contactHrsTo" validate:"min=1,max=50,required"`
	ContactMinTo                string `json:"contactMinTo" validate:"min=1,max=50,required"`
	ContactPersonName           string `json:"contactPersonName" validate:"min=1,max=50,required"`
	ExpYearsCategory            string `json:"expYearsCategory" validate:"min=1,max=50,required"`
	ExpertProfileDesc           string `json:"expertProfileDesc" validate:"min=1,max=50,required"`
	ServiceState                string `json:"serviceState" validate:"min=1,max=50,required"`
	InstName                    string `json:"instName" validate:"min=1,max=50,required"`
	InstDegree                  string `json:"instDegree" validate:"min=1,max=50,required"`
	InstDegreeUrl               string `json:"instDegreeUrl" validate:"min=1,max=50,required"`
	EdLevel                     string `json:"edLevel" validate:"min=1,max=50,required"`
	CertName                    string `json:"certName" validate:"min=1,max=50,required"`
	CertInst                    string `json:"certInst" validate:"min=1,max=50,required"`
	CertUrl                     string `json:"certUrl" validate:"min=1,max=50,required"`
	AchievementName             string `json:"achievementName" validate:"min=1,max=50,required"`
	AchievementDescription      string `json:"achievementDescription" validate:"min=1,max=50,required"`
	AchievementProofUrls        string `json:"achievementProofUrls" validate:"min=1,max=50,required"`
	ExpertOthConstMatSup        string `json:"expertOthConstMatSup" validate:"min=1,max=50,required"`
	ExpertOthMechSer            string `json:"expertOthMechSer" validate:"min=1,max=50,required"`
	SupplierTileLabor           string `json:"supplierTileLabor" validate:"min=1,max=50,required"`
	ExpertOthHardware           string `json:"expertOthHardware" validate:"min=1,max=50,required"`
	SupplierOthEqp              string `json:"supplierOthEqp" validate:"min=1,max=50,required"`
	SupplierOthFin              string `json:"supplierOthFin" validate:"min=1,max=50,required"`
	ProjDesc                    string `json:"projDesc" validate:"min=1,max=50,required"`
	ProjLoc                     string `json:"projLoc" validate:"min=1,max=50,required"`
	ProjArea                    string `json:"projArea" validate:"min=1,max=50,required"`
	ProjBudget                  string `json:"projBudget" validate:"min=1,max=50,required"`
	ProjType                    string `json:"projType" validate:"min=1,max=50,required"`
	PassingOutDate              string `json:"passingOutDate" validate:"min=1,max=50,required"`
	ExpertSpcRes                bool   `json:"expertSpcRes"`
	ExpertSpcCom                bool   `json:"expertSpcCom"`
	ExpertSpcInst               bool   `json:"expertSpcInst"`
	ExpertSpcHealth             bool   `json:"expertSpcHealth"`
	ExpertSpcInfra              bool   `json:"expertSpcInfra"`
	ExpertSpcCSR                bool   `json:"expertSpcCSR"`
	ExpertArch                  bool   `json:"expertArch"`
	ExpertArchConceptDraw       bool   `json:"expertArchConceptDraw"`
	ExpertArchWorkingDraw       bool   `json:"expertArchWorkingDraw"`
	ExpertArch3DRender          bool   `json:"expertArch3DRender"`
	ExpertArchScaleModel        bool   `json:"expertArchScaleModel"`
	ExpertArch3DModel           bool   `json:"expertArch3DModel"`
	ExpertArchVirModel          bool   `json:"expertArchVirModel"`
	ExpertArchExec              bool   `json:"expertArchExec"`
	ExpertStruct                bool   `json:"expertStruct"`
	ExpertStructWorkingDraw     bool   `json:"expertStructWorkingDraw"`
	ExpertStructExec            bool   `json:"expertStructExec"`
	ExpertInterior              bool   `json:"expertInterior"`
	ExpertIntConceptDraw        bool   `json:"expertIntConceptDraw"`
	ExpertIntWorkingDraw        bool   `json:"expertIntWorkingDraw"`
	ExpertIntPhyScale           bool   `json:"expertIntPhyScale"`
	ExpertInt3DModel            bool   `json:"expertInt3DModel"`
	ExpertIntScaleModel         bool   `json:"expertIntScaleModel"`
	ExpertIntVirModel           bool   `json:"expertIntVirModel"`
	ExpertIntExec               bool   `json:"expertIntExec"`
	ExpertLScape                bool   `json:"expertLScape"`
	ExpertSiteSupervisor        bool   `json:"expertSiteSupervisor"`
	ExpertSoilTest              bool   `json:"expertSoilTest"`
	ExpertSanctions             bool   `json:"expertSanctions"`
	SupplierConstMat            bool   `json:"supplierConstMat"`
	SupplierCement              bool   `json:"supplierCement"`
	SupplierSand                bool   `json:"supplierSand"`
	SupplierBrick               bool   `json:"supplierBrick"`
	SupplierPrecast             bool   `json:"supplierPrecast"`
	SupplierCoarse              bool   `json:"supplierCoarse"`
	SupplierConcrete            bool   `json:"supplierConcrete"`
	SupplierReinfbar            bool   `json:"supplierReinfbar"`
	SupplierTile                bool   `json:"supplierTile"`
	SupplierWproofMat           bool   `json:"supplierWproofMat"`
	SupplierShuttering          bool   `json:"supplierShuttering"`
	SupplierScaffs              bool   `json:"supplierScaffs"`
	ExpertMechanical            bool   `json:"expertMechanical"`
	ExpertHeating               bool   `json:"expertHeating"`
	ExpertAc                    bool   `json:"expertAc"`
	ExpertElevator              bool   `json:"expertElevator"`
	ExpertElectrical            bool   `json:"expertElectrical"`
	ExpertElecWiring            bool   `json:"expertElecWiring"`
	ExpertElecFix               bool   `json:"expertElecFix"`
	ExpertOthElecSer            bool   `json:"expertOthElecSer"`
	ExpertPlumbing              bool   `json:"expertPlumbing"`
	SupplierLabor               bool   `json:"supplierLabor"`
	SupplierCivilLabor          bool   `json:"supplierCivilLabor"`
	SupplierPainter             bool   `json:"supplierPainter"`
	SupplierCarpenter           bool   `json:"supplierCarpenter"`
	SupplierPlumbing            bool   `json:"supplierPlumbing"`
	SupplierWaterTank           bool   `json:"supplierWaterTank"`
	SupplierWaterPipeAndFitting bool   `json:"supplierWaterPipeAndFitting"`
	SupplierWaterFix            bool   `json:"supplierWaterFix"`
	SupplierElec                bool   `json:"supplierElec"`
	SupplierWiringMat           bool   `json:"supplierWiringMat"`
	SupplierSwitch              bool   `json:"supplierSwitch"`
	SupplierSensorSup           bool   `json:"supplierSensorSup"`
	SupplierEleAppSup           bool   `json:"supplierEleAppSup"`
	SupplierOthEle              bool   `json:"supplierOthEle"`
	SupplierPaint               bool   `json:"supplierPaint"`
	SupplierIntPaint            bool   `json:"supplierIntPaint"`
	SupplierExtPaint            bool   `json:"supplierExtPaint"`
	SupplierIntPrimer           bool   `json:"supplierIntPrimer"`
	SupplierExtPrimer           bool   `json:"supplierExtPrimer"`
	SupplierPutty               bool   `json:"supplierPutty"`
	SupplierPaintAccessory      bool   `json:"supplierPaintAccessory"`
	SupplierOtherPaint          bool   `json:"supplierOtherPaint"`
	ExpertFabrication           bool   `json:"expertFabrication"`
	ExpertGlassFab              bool   `json:"expertGlassFab"`
	ExpertWoodFab               bool   `json:"expertWoodFab"`
	ExpertMetalAl               bool   `json:"expertMetalAl"`
	ExpertMetalFe               bool   `json:"expertMetalFe"`
	ExpertMetalSS               bool   `json:"expertMetalSS"`
	SupplierConstEqp            bool   `json:"supplierConstEqp"`
	SupplierJcb                 bool   `json:"supplierJcb"`
	SupplierConcMixer           bool   `json:"supplierConcMixer"`
	SupplierHaulTruck           bool   `json:"supplierHaulTruck"`
	SupplierCrane               bool   `json:"supplierCrane"`
	SupplierHouseProduct        bool   `json:"supplierHouseProduct"`
	SupplierFurniture           bool   `json:"supplierFurniture"`
	SupplierAppliance           bool   `json:"supplierAppliance"`
	SupplierFurnish             bool   `json:"supplierFurnish"`
}

// Validate function used to validate the request.
// This is the implementation of the Validation interface.
func (g CreateProfessionalInfoRequest) Validate() error {
	return common.Validate.Struct(g)
}

// CreateProfessionalInfoResponse response struct
type CreateProfessionalInfoResponse struct {
	Response CreateProfessionalInfoData `json:"response"`
}

type CreateProfessionalInfoData struct {
	Id      string `json:"id"`
	Message string `json:"message"`
}

// UpdateProfessionalInfoRequest struct used to validate the Update Professional Info Request
type UpdateProfessionalInfoRequest struct {
	Id                          string `json:"id" validate:"min=1,max=50,required"`
	UserAlias                   string `json:"userAlias" validate:"min=1,max=50,required"`
	EntityType                  string `json:"entityType" validate:"min=1,max=50,required"`
	ContactHrsFrom              string `json:"contactHrsFrom" validate:"min=1,max=50,required"`
	ContactMinFrom              string `json:"contactMinFrom" validate:"min=1,max=50,required"`
	ContactHrsTo                string `json:"contactHrsTo" validate:"min=1,max=50,required"`
	ContactMinTo                string `json:"contactMinTo" validate:"min=1,max=50,required"`
	ContactPersonName           string `json:"contactPersonName" validate:"min=1,max=50,required"`
	ExpYearsCategory            string `json:"expYearsCategory" validate:"min=1,max=50,required"`
	ExpertProfileDesc           string `json:"expertProfileDesc" validate:"min=1,max=50,required"`
	ServiceState                string `json:"serviceState" validate:"min=1,max=50,required"`
	InstName                    string `json:"instName" validate:"min=1,max=50,required"`
	InstDegree                  string `json:"instDegree" validate:"min=1,max=50,required"`
	InstDegreeUrl               string `json:"instDegreeUrl" validate:"min=1,max=50,required"`
	EdLevel                     string `json:"edLevel" validate:"min=1,max=50,required"`
	CertName                    string `json:"certName" validate:"min=1,max=50,required"`
	CertInst                    string `json:"certInst" validate:"min=1,max=50,required"`
	CertUrl                     string `json:"certUrl" validate:"min=1,max=50,required"`
	AchievementName             string `json:"achievementName" validate:"min=1,max=50,required"`
	AchievementDescription      string `json:"achievementDescription" validate:"min=1,max=50,required"`
	AchievementProofUrls        string `json:"achievementProofUrls" validate:"min=1,max=50,required"`
	ExpertOthConstMatSup        string `json:"expertOthConstMatSup" validate:"min=1,max=50,required"`
	ExpertOthMechSer            string `json:"expertOthMechSer" validate:"min=1,max=50,required"`
	SupplierTileLabor           string `json:"supplierTileLabor" validate:"min=1,max=50,required"`
	ExpertOthHardware           string `json:"expertOthHardware" validate:"min=1,max=50,required"`
	SupplierOthEqp              string `json:"supplierOthEqp" validate:"min=1,max=50,required"`
	SupplierOthFin              string `json:"supplierOthFin" validate:"min=1,max=50,required"`
	ProjDesc                    string `json:"projDesc" validate:"min=1,max=50,required"`
	ProjLoc                     string `json:"projLoc" validate:"min=1,max=50,required"`
	ProjArea                    string `json:"projArea" validate:"min=1,max=50,required"`
	ProjBudget                  string `json:"projBudget" validate:"min=1,max=50,required"`
	ProjType                    string `json:"projType" validate:"min=1,max=50,required"`
	PassingOutDate              string `json:"passingOutDate" validate:"min=1,max=50,required"`
	ExpertSpcRes                bool   `json:"expertSpcRes"`
	ExpertSpcCom                bool   `json:"expertSpcCom"`
	ExpertSpcInst               bool   `json:"expertSpcInst"`
	ExpertSpcHealth             bool   `json:"expertSpcHealth"`
	ExpertSpcInfra              bool   `json:"expertSpcInfra"`
	ExpertSpcCSR                bool   `json:"expertSpcCSR"`
	ExpertArch                  bool   `json:"expertArch"`
	ExpertArchConceptDraw       bool   `json:"expertArchConceptDraw"`
	ExpertArchWorkingDraw       bool   `json:"expertArchWorkingDraw"`
	ExpertArch3DRender          bool   `json:"expertArch3DRender"`
	ExpertArchScaleModel        bool   `json:"expertArchScaleModel"`
	ExpertArch3DModel           bool   `json:"expertArch3DModel"`
	ExpertArchVirModel          bool   `json:"expertArchVirModel"`
	ExpertArchExec              bool   `json:"expertArchExec"`
	ExpertStruct                bool   `json:"expertStruct"`
	ExpertStructWorkingDraw     bool   `json:"expertStructWorkingDraw"`
	ExpertStructExec            bool   `json:"expertStructExec"`
	ExpertInterior              bool   `json:"expertInterior"`
	ExpertIntConceptDraw        bool   `json:"expertIntConceptDraw"`
	ExpertIntWorkingDraw        bool   `json:"expertIntWorkingDraw"`
	ExpertIntPhyScale           bool   `json:"expertIntPhyScale"`
	ExpertInt3DModel            bool   `json:"expertInt3DModel"`
	ExpertIntScaleModel         bool   `json:"expertIntScaleModel"`
	ExpertIntVirModel           bool   `json:"expertIntVirModel"`
	ExpertIntExec               bool   `json:"expertIntExec"`
	ExpertLScape                bool   `json:"expertLScape"`
	ExpertSiteSupervisor        bool   `json:"expertSiteSupervisor"`
	ExpertSoilTest              bool   `json:"expertSoilTest"`
	ExpertSanctions             bool   `json:"expertSanctions"`
	SupplierConstMat            bool   `json:"supplierConstMat"`
	SupplierCement              bool   `json:"supplierCement"`
	SupplierSand                bool   `json:"supplierSand"`
	SupplierBrick               bool   `json:"supplierBrick"`
	SupplierPrecast             bool   `json:"supplierPrecast"`
	SupplierCoarse              bool   `json:"supplierCoarse"`
	SupplierConcrete            bool   `json:"supplierConcrete"`
	SupplierReinfbar            bool   `json:"supplierReinfbar"`
	SupplierTile                bool   `json:"supplierTile"`
	SupplierWproofMat           bool   `json:"supplierWproofMat"`
	SupplierShuttering          bool   `json:"supplierShuttering"`
	SupplierScaffs              bool   `json:"supplierScaffs"`
	ExpertMechanical            bool   `json:"expertMechanical"`
	ExpertHeating               bool   `json:"expertHeating"`
	ExpertAc                    bool   `json:"expertAc"`
	ExpertElevator              bool   `json:"expertElevator"`
	ExpertElectrical            bool   `json:"expertElectrical"`
	ExpertElecWiring            bool   `json:"expertElecWiring"`
	ExpertElecFix               bool   `json:"expertElecFix"`
	ExpertOthElecSer            bool   `json:"expertOthElecSer"`
	ExpertPlumbing              bool   `json:"expertPlumbing"`
	SupplierLabor               bool   `json:"supplierLabor"`
	SupplierCivilLabor          bool   `json:"supplierCivilLabor"`
	SupplierPainter             bool   `json:"supplierPainter"`
	SupplierCarpenter           bool   `json:"supplierCarpenter"`
	SupplierPlumbing            bool   `json:"supplierPlumbing"`
	SupplierWaterTank           bool   `json:"supplierWaterTank"`
	SupplierWaterPipeAndFitting bool   `json:"supplierWaterPipeAndFitting"`
	SupplierWaterFix            bool   `json:"supplierWaterFix"`
	SupplierElec                bool   `json:"supplierElec"`
	SupplierWiringMat           bool   `json:"supplierWiringMat"`
	SupplierSwitch              bool   `json:"supplierSwitch"`
	SupplierSensorSup           bool   `json:"supplierSensorSup"`
	SupplierEleAppSup           bool   `json:"supplierEleAppSup"`
	SupplierOthEle              bool   `json:"supplierOthEle"`
	SupplierPaint               bool   `json:"supplierPaint"`
	SupplierIntPaint            bool   `json:"supplierIntPaint"`
	SupplierExtPaint            bool   `json:"supplierExtPaint"`
	SupplierIntPrimer           bool   `json:"supplierIntPrimer"`
	SupplierExtPrimer           bool   `json:"supplierExtPrimer"`
	SupplierPutty               bool   `json:"supplierPutty"`
	SupplierPaintAccessory      bool   `json:"supplierPaintAccessory"`
	SupplierOtherPaint          bool   `json:"supplierOtherPaint"`
	ExpertFabrication           bool   `json:"expertFabrication"`
	ExpertGlassFab              bool   `json:"expertGlassFab"`
	ExpertWoodFab               bool   `json:"expertWoodFab"`
	ExpertMetalAl               bool   `json:"expertMetalAl"`
	ExpertMetalFe               bool   `json:"expertMetalFe"`
	ExpertMetalSS               bool   `json:"expertMetalSS"`
	SupplierConstEqp            bool   `json:"supplierConstEqp"`
	SupplierJcb                 bool   `json:"supplierJcb"`
	SupplierConcMixer           bool   `json:"supplierConcMixer"`
	SupplierHaulTruck           bool   `json:"supplierHaulTruck"`
	SupplierCrane               bool   `json:"supplierCrane"`
	SupplierHouseProduct        bool   `json:"supplierHouseProduct"`
	SupplierFurniture           bool   `json:"supplierFurniture"`
	SupplierAppliance           bool   `json:"supplierAppliance"`
	SupplierFurnish             bool   `json:"supplierFurnish"`
}

// Validate function used to validate the request.
// This is the implementation of the Validation interface.
func (g UpdateProfessionalInfoRequest) Validate() error {
	return common.Validate.Struct(g)
}

// UpdateProfessionalInfoResponse response struct
type UpdateProfessionalInfoResponse struct {
	Response UpdateProfessionalInfoData `json:"response"`
}

type UpdateProfessionalInfoData struct {
	Message string `json:"message"`
}
