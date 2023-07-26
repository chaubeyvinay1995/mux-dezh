package service

import (
	"context"
	"gitlab.com/umi7/DezHab_user_backend/dao/models"
)

type Certificate interface {
	Get(ctx context.Context, OrganizationId string) ([]models.OrganizationCertificate, error)
	Create(ctx context.Context, OrganizationId, CourseName, Url, Institute string)(certificateId string, err error)
	Update(ctx context.Context, Id, CourseName, Url, Institute string)(err error)
	Delete(ctx context.Context, Id string) error
}
