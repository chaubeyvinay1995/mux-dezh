package models

import "time"

type UserProfessionalInfo struct {
	ID                          string `gorm:"primary_key;varchar(50)"`
	UserId                      string `sql:"type:varchar REFERENCES user_personal_infos(id)"` // F.K with UserPersonal info table
	UserAlias                   string `gorm:"varchar(80)"`
	EntityType                  string `gorm:"varchar(10)"`
	ContactHrsFrom              string `gorm:"varchar(2)"`
	ContactMinFrom              string `gorm:"varchar(2)"`
	ContactHrsTo                string `gorm:"varchar(2)"`
	ContactMinTo                string `gorm:"varchar(2)"`
	ContactPersonName           string `gorm:"varchar(80)"`
	ExpYearsCategory            string `gorm:"varchar(2)"`
	ExpertSpcRes                bool   `gorm:"type:bool;default:false"`
	ExpertSpcCom                bool   `gorm:"type:bool;default:false"`
	ExpertSpcInst               bool   `gorm:"type:bool;default:false"`
	ExpertSpcHealth             bool   `gorm:"type:bool;default:false"`
	ExpertSpcInfra              bool   `gorm:"type:bool;default:false"`
	ExpertSpcCSR                bool   `gorm:"type:bool;default:false"`
	ExpertProfileDesc           string `gorm:"varchar(500)"`
	ServiceState                string `gorm:"varchar(50)"`
	InstName                    string `gorm:"varchar(200)"`
	InstDegree                  string `gorm:"varchar(200)"`
	InstDegreeUrl               string `gorm:"varchar(1000)"`
	EdLevel                     string `gorm:"varchar(100)"`
	CertName                    string `gorm:"varchar(100)"`
	CertInst                    string `gorm:"varchar(100)"`
	CertUrl                     string `gorm:"varchar(1000)"`
	AchievementName             string `gorm:"varchar(200)"`
	AchievementDescription      string `gorm:"varchar(255)"`
	AchievementProofUrls        string `gorm:"varchar(1000)"`
	ExpertArch                  bool   `gorm:"type:bool;default:false"`
	ExpertArchConceptDraw       bool   `gorm:"type:bool;default:false"`
	ExpertArchWorkingDraw       bool   `gorm:"type:bool;default:false"`
	ExpertArch3DRender          bool   `gorm:"type:bool;default:false"`
	ExpertArchScaleModel        bool   `gorm:"type:bool;default:false"`
	ExpertArch3DModel           bool   `gorm:"type:bool;default:false"`
	ExpertArchVirModel          bool   `gorm:"type:bool;default:false"`
	ExpertArchExec              bool   `gorm:"type:bool;default:false"`
	ExpertStruct                bool   `gorm:"type:bool;default:false"`
	ExpertStructWorkingDraw     bool   `gorm:"type:bool;default:false"`
	ExpertStructExec            bool   `gorm:"type:bool;default:false"`
	ExpertInterior              bool   `gorm:"type:bool;default:false"`
	ExpertIntConceptDraw        bool   `gorm:"type:bool;default:false"`
	ExpertIntWorkingDraw        bool   `gorm:"type:bool;default:false"`
	ExpertIntPhyScale           bool   `gorm:"type:bool;default:false"`
	ExpertInt3DModel            bool   `gorm:"type:bool;default:false"`
	ExpertIntScaleModel         bool   `gorm:"type:bool;default:false"`
	ExpertIntVirModel           bool   `gorm:"type:bool;default:false"`
	ExpertIntExec               bool   `gorm:"type:bool;default:false"`
	ExpertLScape                bool   `gorm:"type:bool;default:false"`
	ExpertSiteSupervisor        bool   `gorm:"type:bool;default:false"`
	ExpertSoilTest              bool   `gorm:"type:bool;default:false"`
	ExpertSanctions             bool   `gorm:"type:bool;default:false"`
	SupplierConstMat            bool   `gorm:"type:bool;default:false"`
	SupplierCement              bool   `gorm:"type:bool;default:false"`
	SupplierSand                bool   `gorm:"type:bool;default:false"`
	SupplierBrick               bool   `gorm:"type:bool;default:false"`
	SupplierPrecast             bool   `gorm:"type:bool;default:false"`
	SupplierCoarse              bool   `gorm:"type:bool;default:false"`
	SupplierConcrete            bool   `gorm:"type:bool;default:false"`
	SupplierReinfbar            bool   `gorm:"type:bool;default:false"`
	SupplierTile                bool   `gorm:"type:bool;default:false"`
	SupplierWproofMat           bool   `gorm:"type:bool;default:false"`
	SupplierShuttering          bool   `gorm:"type:bool;default:false"`
	SupplierScaffs              bool   `gorm:"type:bool;default:false"`
	ExpertOthConstMatSup        string `gorm:"varchar(50)"`
	ExpertMechanical            bool   `gorm:"type:bool;default:false"`
	ExpertHeating               bool   `gorm:"type:bool;default:false"`
	ExpertAc                    bool   `gorm:"type:bool;default:false"`
	ExpertElevator              bool   `gorm:"type:bool;default:false"`
	ExpertOthMechSer            string `gorm:"varchar(50)"`
	ExpertElectrical            bool   `gorm:"type:bool;default:false"`
	ExpertElecWiring            bool   `gorm:"type:bool;default:false"`
	ExpertElecFix               bool   `gorm:"type:bool;default:false"`
	ExpertOthElecSer            bool   `gorm:"varchar(50)"`
	ExpertPlumbing              bool   `gorm:"type:bool;default:false"`
	SupplierLabor               bool   `gorm:"type:bool;default:false"`
	SupplierCivilLabor          bool   `gorm:"type:bool;default:false"`
	SupplierTileLabor           string `gorm:"type:bool;default:false"`
	SupplierPainter             bool   `gorm:"type:bool;default:false"`
	SupplierCarpenter           bool   `gorm:"type:bool;default:false"`
	SupplierPlumbing            bool   `gorm:"type:bool;default:false"`
	SupplierWaterTank           bool   `gorm:"type:bool;default:false"`
	SupplierWaterPipeAndFitting bool   `gorm:"type:bool;default:false"`
	SupplierWaterFix            bool   `gorm:"type:bool;default:false"`
	SupplierElec                bool   `gorm:"type:bool;default:false"`
	SupplierWiringMat           bool   `gorm:"type:bool;default:false"`
	SupplierSwitch              bool   `gorm:"type:bool;default:false"`
	SupplierSensorSup           bool   `gorm:"type:bool;default:false"`
	SupplierEleAppSup           bool   `gorm:"type:bool;default:false"`
	SupplierOthEle              bool   `gorm:"type:bool;default:false"`
	SupplierPaint               bool   `gorm:"type:bool;default:false"`
	SupplierIntPaint            bool   `gorm:"type:bool;default:false"`
	SupplierExtPaint            bool   `gorm:"type:bool;default:false"`
	SupplierIntPrimer           bool   `gorm:"type:bool;default:false"`
	SupplierExtPrimer           bool   `gorm:"type:bool;default:false"`
	SupplierPutty               bool   `gorm:"type:bool;default:false"`
	SupplierPaintAccessory      bool   `gorm:"type:bool;default:false"`
	SupplierOtherPaint          bool   `gorm:"type:bool;default:false"`
	ExpertFabrication           bool   `gorm:"type:bool;default:false"`
	ExpertGlassFab              bool   `gorm:"type:bool;default:false"`
	ExpertWoodFab               bool   `gorm:"type:bool;default:false"`
	ExpertMetalAl               bool   `gorm:"type:bool;default:false"`
	ExpertMetalFe               bool   `gorm:"type:bool;default:false"`
	ExpertMetalSS               bool   `gorm:"type:bool;default:false"`
	ExpertOthHardware           string `gorm:"varchar(50)"`
	SupplierConstEqp            bool   `gorm:"type:bool;default:false"`
	SupplierJcb                 bool   `gorm:"type:bool;default:false"`
	SupplierConcMixer           bool   `gorm:"type:bool;default:false"`
	SupplierHaulTruck           bool   `gorm:"type:bool;default:false"`
	SupplierCrane               bool   `gorm:"type:bool;default:false"`
	SupplierOthEqp              string `gorm:"varchar(50)"`
	SupplierHouseProduct        bool   `gorm:"type:bool;default:false"`
	SupplierFurniture           bool   `gorm:"type:bool;default:false"`
	SupplierAppliance           bool   `gorm:"type:bool;default:false"`
	SupplierFurnish             bool   `gorm:"type:bool;default:false"`
	SupplierOthFin              string `gorm:"varchar(50)"`
	ProjDesc                    string `gorm:"varchar(200)"`
	ProjLoc                     string `gorm:"varchar(200)"`
	ProjArea                    string `gorm:"varchar(100)"`
	ProjBudget                  string `gorm:"varchar(10)"`
	ProjType                    string `gorm:"varchar(20)"`
	PassingOutDate              *time.Time
	DeletedAt                   *time.Time `sql:"index"`
	CreatedAt                   time.Time
	UpdatedAt                   time.Time
}