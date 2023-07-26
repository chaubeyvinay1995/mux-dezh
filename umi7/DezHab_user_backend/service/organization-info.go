package service

import (
	"context"
	"gitlab.com/umi7/DezHab_user_backend/dao/models"
	"time"
)

type OrganizationInfo interface {
	Get(ctx context.Context, userId string) (models.OrganizationInfo, error)

	Create(ctx context.Context, userId , entityType, firstName, middleName , lastName, contactPersonName,
		gender, city, state, country, address1, address2, primaryContactNumber, secondaryContactNumber,
		emailId, websiteUrl, facebookUrl, about, professionalExperience string, inTime, outTime time.Time) (organizationId string, err error)

	Update(ctx context.Context, id, userID, entityType, firstName, middleName, lastName, contactPersonName, gender, city, state,
		country, address1, address2, primaryContactNumber, secondaryContactNumber,
		websiteUrl, facebookUrl, about, professionalExperience string) error
}

// Command to generate the mock implementation of the BookKeeping interface
//go:generate mockgen -source=book-keeping.go -destination=mock-impl/organization_info_impl.go -package=mock_impl
