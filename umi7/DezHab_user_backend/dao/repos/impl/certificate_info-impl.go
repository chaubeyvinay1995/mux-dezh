package impl

import (
	"context"
	"errors"
	"gitlab.com/umi7/DezHab_user_backend/dao/core/impl"
	"gitlab.com/umi7/DezHab_user_backend/dao/models"
	"gitlab.com/umi7/DezHab_user_backend/dao/repos"
	"time"
)

var CertificateInfo repos.CertificateInfo

type CertificateInfoImpl struct{}

var updateCertificateQuery = "id=?"

func init() {
	CertificateInfo = CertificateInfoImpl{}
}

func (certificateInfoImpl CertificateInfoImpl) Search(ctx context.Context,
	organizationId string) (certificateInfo []models.OrganizationCertificate, err error) {
	var certificateList []models.OrganizationCertificate
	if err = impl.Read.Find(&certificateList, models.OrganizationCertificate{OrganizationId: organizationId},
		&certificateInfo); err != nil {
		return
	}
	if len(certificateList) == 0 {
		err = errors.New(ErrNoData)
		return
	}
	certificateInfo = certificateList
	return
}

func (certificateInfoImpl CertificateInfoImpl) Create(ctx context.Context,
	certificateInfo models.OrganizationCertificate) (err error){
	certificateInfo.CreatedAt = time.Now().UTC()
	certificateInfo.UpdatedAt = certificateInfo.CreatedAt
	err = impl.Write.Create(&certificateInfo)
	return
}

func (certificateInfoImpl CertificateInfoImpl) Update(ctx context.Context,
	certificateInfo models.OrganizationCertificate) (err error){
	certificateInfo.UpdatedAt = time.Now().UTC()
	// Map with key value pair
	updateValues := map[string]interface{}{
		"course_name": certificateInfo.CourseName, "url": certificateInfo.Url,
		"institute": certificateInfo.Institute,
	}
	err = impl.Write.Update(&certificateInfo, updateCertificateQuery, certificateInfo.ID, updateValues)
	return
}

func (certificateInfoImpl CertificateInfoImpl) UpdateFile(ctx context.Context,
	certificateInfo models.OrganizationCertificate) (err error){
	certificateInfo.UpdatedAt = time.Now().UTC()
	// Map with key value pair
	updateValues := map[string]interface{}{
		"url": certificateInfo.Url,
	}
	err = impl.Write.Update(&certificateInfo, updateCertificateQuery, certificateInfo.ID, updateValues)
	return
}

func (certificateInfoImpl CertificateInfoImpl) Delete(ctx context.Context,
	certificateInfo models.OrganizationCertificate) (err error){
	err = impl.Write.Delete(&certificateInfo)
	return

}

