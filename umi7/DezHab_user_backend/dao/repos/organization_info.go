package repos

import (
	"context"
	"gitlab.com/umi7/DezHab_user_backend/dao/models"
)


type OrganizationInfo interface {
	Create(context.Context, models.OrganizationInfo) error
	Search(context.Context, string) (models.OrganizationInfo, error)
	Update(context.Context, models.OrganizationInfo) error
}

// Command to generate the mock implementation of the BookKeeping interface
//go:generate mockgen -source=book_keeping.go -destination=mock-impl/organization_info_impl.go -package=mock_impl
