package impl

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"gitlab.com/umi7/DezHab_user_backend/common/logger"
	"gitlab.com/umi7/DezHab_user_backend/dao/models"
	"gitlab.com/umi7/DezHab_user_backend/dao/repos/impl"
	"gitlab.com/umi7/DezHab_user_backend/service"
)

// AcademicInfo is the variable of the interface type
var AcademicInfo service.Academic

type AcademicImpl struct{}

func init() {
	AcademicInfo = AcademicImpl{}
}

func (p AcademicImpl) Get(ctx context.Context,
	organizationId string) (academicInfo []models.OrganizationAcademicDetail, err error) {
	academicInfo, err = impl.AcademicInfo.Search(ctx, organizationId)
	if err != nil {
		logger.Error(ctx, err)
		return
	}
	return
}

func (p AcademicImpl) Create(ctx context.Context, organizationId, degree, instituteName,
	certificateLink string) (academicId string, err error) {
	// Create UUID
	academicId = uuid.New().String()
	academicInfoData := models.OrganizationAcademicDetail{
		ID: academicId, OrganizationId: organizationId, Degree: degree,
		InstituteName: instituteName, CertificateLink: certificateLink,
	}
	err = impl.AcademicInfo.Create(ctx, academicInfoData)
	if err != nil {
		logger.Error(ctx, fmt.Sprintf("Error occurred while creating entry on academic info : %v", err))
	}
	return
}

func (p AcademicImpl) Update(ctx context.Context, id, degree, instituteName, certificateLink string) (err error) {
	academicInfo := models.OrganizationAcademicDetail{
		ID: id, Degree: degree, InstituteName: instituteName, CertificateLink: certificateLink,
	}
	err = impl.AcademicInfo.Update(ctx, academicInfo)
	if err != nil {
		logger.Error(ctx, fmt.Sprintf("Error occurred while creating entry on organization info : %v", err))
	}
	return
}

func (p AcademicImpl) Delete(ctx context.Context, id string) (err error) {
	academicInfo := models.OrganizationAcademicDetail{
		ID: id,
	}
	err = impl.AcademicInfo.Delete(ctx, academicInfo)
	if err != nil {
		logger.Error(ctx, fmt.Sprintf("Error occurred while creating entry on organization info : %v", err))
	}
	return

}
