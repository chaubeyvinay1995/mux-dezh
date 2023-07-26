package impl

import (
	"context"
	"errors"
	"gitlab.com/umi7/DezHab_user_backend/dao/core/impl"
	"gitlab.com/umi7/DezHab_user_backend/dao/models"
	"gitlab.com/umi7/DezHab_user_backend/dao/repos"
	"time"
)

var AcademicInfo repos.AcademicInfo

type AcademicInfoImpl struct{}

var updateAcademicQuery = "id=?"

func init() {
	AcademicInfo = AcademicInfoImpl{}
}

func (academicInfoImpl AcademicInfoImpl) Search(ctx context.Context,
	organizationId string) (academicInfo []models.OrganizationAcademicDetail, err error) {
	var academicList []models.OrganizationAcademicDetail
	if err = impl.Read.Find(&academicList, models.OrganizationAcademicDetail{OrganizationId: organizationId},
	&academicInfo); err != nil {
		return
	}
	if len(academicList) == 0 {
		err = errors.New(ErrNoData)
		return
	}
	academicInfo = academicList
	return
}

func (academicInfoImpl AcademicInfoImpl) Create(ctx context.Context,
	academicInfo models.OrganizationAcademicDetail) (err error) {
	academicInfo.CreatedAt = time.Now().UTC()
	academicInfo.UpdatedAt = academicInfo.CreatedAt
	err = impl.Write.Create(&academicInfo)
	return
}


func (academicInfoImpl AcademicInfoImpl) Update(ctx context.Context,
	academicInfo models.OrganizationAcademicDetail) (err error) {
	academicInfo.UpdatedAt = time.Now().UTC()
	updateValues := map[string]interface{}{
		"degree": academicInfo.Degree, "institute_name": academicInfo.InstituteName,
		"certificate_link": academicInfo.CertificateLink,
	}
	err = impl.Write.Update(&academicInfo, updateAcademicQuery, academicInfo.ID, updateValues)
	return
}

func (academicInfoImpl AcademicInfoImpl) Delete(ctx context.Context,
	academicInfo models.OrganizationAcademicDetail) (err error){
	err = impl.Write.Delete(&academicInfo)
	return
}