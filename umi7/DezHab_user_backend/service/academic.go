package service

import (
	"context"
	"gitlab.com/umi7/DezHab_user_backend/dao/models"
)

type Academic interface {
	Get(ctx context.Context, OrganizationId string) ([]models.OrganizationAcademicDetail, error)
	Create(ctx context.Context, OrganizationId, Degree, InstituteName,
		CertificateLink string) (academicId string, err error)
	Update(ctx context.Context, Id, Degree, InstituteName, CertificateLink string) error
	Delete(ctx context.Context, Id string) error
}
