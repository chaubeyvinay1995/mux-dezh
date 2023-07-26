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

var OrganizationInfo service.OrganizationInfo

type OrganizationInfoImpl struct{}

func init() {
	OrganizationInfo = OrganizationInfoImpl{}
}


func (o OrganizationInfoImpl) Get(ctx context.Context, userId string) (organizationInfo models.OrganizationInfo, err error) {
	organizationInfo, err = impl.OrganizationInfo.Search(ctx, userId)
	if err != nil {
		logger.Error(ctx, err)
		return
	}
	return

}

// Create OrganizationInfo
func (o OrganizationInfoImpl) Create(ctx context.Context, userId , entityType, firstName, middleName , lastName,
	contactPersonName, gender, city, state, country, address1, address2, primaryContactNumber,
	secondaryContactNumber, emailId, websiteUrl, facebookUrl, about,
	professionalExperience string, inTime, outTime time.Time) (organizationId string, err error) {
	// Create UUID
	organizationId = uuid.New().String()
	organizationInfo := models.OrganizationInfo{
		ID:	organizationId, UserId: userId, EntityType: entityType, FirstName: firstName, LastName: lastName,
		MiddleName: middleName, ContactPersonName: contactPersonName, Gender: gender, City: city, State: state,
		Country: country, Address1: address1, Address2: address2, PrimaryContactNumber: primaryContactNumber,
		SecondaryContactNumber: secondaryContactNumber, EmailId: emailId, WebsiteUrl: websiteUrl,
		FacebookUrl: facebookUrl, About: about, PreferredContactInTime: inTime, PreferredContactOutTime: outTime,
		ProfessionalExperience: professionalExperience,
	}
	err = impl.OrganizationInfo.Create(ctx, organizationInfo)
	if err != nil {
		logger.Error(ctx, fmt.Sprintf("Error occurred while creating entry on organization info : %v", err))
	}
	return
}

// Update OrganizationInfo
func (o OrganizationInfoImpl) Update(ctx context.Context, id, userID, entityType, firstName, middleName,
	lastName, contactPersonName, gender, city, state, country, address1, address2, primaryContactNumber,
	secondaryContactNumber, websiteUrl, facebookUrl, about, professionalExperience string) (err error){
	organizationInfo := models.OrganizationInfo{
		ID: id, UserId: userID, EntityType: entityType, FirstName: firstName, MiddleName: middleName,
		LastName: lastName, ContactPersonName: contactPersonName, Gender: gender, City: city, State: state,
		Country: country, Address1: address1, Address2: address2, PrimaryContactNumber: primaryContactNumber,
		SecondaryContactNumber: secondaryContactNumber, WebsiteUrl: websiteUrl, FacebookUrl: facebookUrl,
		About: about, ProfessionalExperience: professionalExperience,
	}
	err = impl.OrganizationInfo.Update(ctx, organizationInfo)
	if err != nil {
		logger.Error(ctx, err)
	}
	return
}