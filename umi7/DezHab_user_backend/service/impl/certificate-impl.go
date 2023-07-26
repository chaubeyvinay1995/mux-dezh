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

var CertificateInfo service.Certificate

type CertificateImpl struct{}

func init() {
	CertificateInfo = CertificateImpl{}
}

func (c CertificateImpl) Get(ctx context.Context,
	organizationId string) (certificateInfo []models.OrganizationCertificate, err error) {
	certificateInfo, err = impl.CertificateInfo.Search(ctx, organizationId)
	if err != nil {
		logger.Error(ctx, err)
		return
	}
	return
}

func (c CertificateImpl) Create(ctx context.Context, organizationId, courseName, institute,
	url string) (certificateId string, err error){
	// Create UUID
	certificateId = uuid.New().String()
	// Create Certificate Model Object of type models.OrganizationCertificate
	certificateInfoData := models.OrganizationCertificate{
		ID: certificateId, OrganizationId: organizationId, CourseName: courseName, Institute: institute, Url: url,
	}
	err = impl.CertificateInfo.Create(ctx, certificateInfoData)
	if err != nil {
		logger.Error(ctx, fmt.Sprintf("Error occurred while creating entry on certificate info : %v", err))
	}
	return
}

func (c CertificateImpl) Update(ctx context.Context, id, courseName, institute, url string) (err error){
	// create struct of the model.OrganizationCertificate type
	certificateInfo := models.OrganizationCertificate{
		ID: id, CourseName: courseName, Institute: institute, Url: url,
	}
	err = impl.CertificateInfo.Update(ctx, certificateInfo)
	if err != nil {
		logger.Error(ctx, fmt.Sprintf("Error occurred while creating entry on Certificate info : %v", err))
	}
	return
}

func (c CertificateImpl) UpdateFile(ctx context.Context, id, url string) (err error){
	// create struct of the model.OrganizationCertificate type
	certificateInfo := models.OrganizationCertificate{
		ID: id, Url: url,
	}
	err = impl.CertificateInfo.Update(ctx, certificateInfo)
	if err != nil {
		logger.Error(ctx, fmt.Sprintf("Error occurred while creating entry on Certificate info : %v", err))
	}
	return
}

func (c CertificateImpl) Delete(ctx context.Context, id string) (err error){
	certificateInfo := models.OrganizationCertificate{
		ID: id,
	}
	err = impl.CertificateInfo.Delete(ctx, certificateInfo)
	if err != nil {
		logger.Error(ctx, fmt.Sprintf("Error occurred while deleting entry from certificate info : %v", err))
	}
	return

}
