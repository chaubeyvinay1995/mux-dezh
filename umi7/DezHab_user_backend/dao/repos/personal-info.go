package repos

import (
	"context"
	"gitlab.com/umi7/DezHab_user_backend/dao/models"
)

type PersonalInfo interface {
	Search(context.Context, string) (models.UserPersonalInfo, error)
	Create(context.Context, models.UserPersonalInfo) error
	Delete(ctx context.Context, personalInfo models.UserPersonalInfo) error
	Update(ctx context.Context, personalInfo models.UserPersonalInfo) error
}
