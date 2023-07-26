package repos

import (
	"context"
	"gitlab.com/umi7/DezHab_user_backend/dao/models"
)


type AcademicInfo interface {
	Search(context.Context, string) ([]models.OrganizationAcademicDetail, error)
	Create(context.Context, models.OrganizationAcademicDetail) error
	Update(ctx context.Context, detail models.OrganizationAcademicDetail) error
	Delete(context.Context, models.OrganizationAcademicDetail) error
}