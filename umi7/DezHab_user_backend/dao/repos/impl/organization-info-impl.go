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
	updateOrganizationQuery = "id=?"
)

var OrganizationInfo repos.OrganizationInfo

type OrganizationInfoImpl struct{}

func init() {
	OrganizationInfo = OrganizationInfoImpl{}
}

func (organizationInfoImpl OrganizationInfoImpl) Search(ctx context.Context, id string) (organizationInfo models.OrganizationInfo, err error) {
	var organizationInfoList []models.OrganizationInfo
	if err = impl.Read.Last(&organizationInfoList, models.OrganizationInfo{ID: id}); err != nil {
		return
	}
	if len(organizationInfoList) == 0 {
		err = errors.New(ErrNoData)
		return
	}
	organizationInfo = organizationInfoList[0]
	return
}

func (organizationInfoImpl OrganizationInfoImpl) Create(ctx context.Context, organizationInfo models.OrganizationInfo) (err error) {
	organizationInfo.CreatedAt = time.Now().UTC()
	organizationInfo.UpdatedAt = organizationInfo.CreatedAt
	err = impl.Write.Create(&organizationInfo)
	return
}

func (organizationInfoImpl OrganizationInfoImpl) Update(ctx context.Context, organizationInfo models.OrganizationInfo) (err error) {
	organizationInfo.UpdatedAt = time.Now().UTC()
	updateValues := map[string]interface{}{
		"user_id": organizationInfo.UserId, "first_name": organizationInfo.FirstName,
		"middle_name": organizationInfo.MiddleName, "last_name": organizationInfo.LastName,
		"about": organizationInfo.About, "contact_person_name": organizationInfo.ContactPersonName,
		"gender": organizationInfo.Gender, "city": organizationInfo.City, "state": organizationInfo.State,
		"country": organizationInfo.Country, "address1": organizationInfo.Address1,
		"address2": organizationInfo.Address2, "primary_contact_number": organizationInfo.PrimaryContactNumber,
		"secondary_contact_number": organizationInfo.SecondaryContactNumber,
		"website_url":              organizationInfo.WebsiteUrl, "facebook_url": organizationInfo.FacebookUrl,
		"professional_experience": organizationInfo.ProfessionalExperience,
	}
	err = impl.Write.Update(&organizationInfo, updateOrganizationQuery, organizationInfo.ID, updateValues)
	return
}
