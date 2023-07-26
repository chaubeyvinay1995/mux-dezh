package repos

import (
	"context"
	"gitlab.com/umi7/DezHab_user_backend/dao/models"
)

type ProfessionalInfo interface {
	Search(context.Context, string) ([]models.UserProfessionalInfo, error)
	Delete(ctx context.Context, professionalInfo models.UserProfessionalInfo) error
	Create(context.Context, models.UserProfessionalInfo) error
	Update(ctx context.Context, professionalInfo models.UserProfessionalInfo) error
}
