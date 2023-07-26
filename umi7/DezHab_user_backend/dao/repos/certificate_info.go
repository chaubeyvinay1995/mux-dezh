package repos

import (
	"context"
	"gitlab.com/umi7/DezHab_user_backend/dao/models"
)


type CertificateInfo interface {
	Search(context.Context, string) ([]models.OrganizationCertificate, error)
	Create(context.Context, models.OrganizationCertificate) error
	Update(ctx context.Context, certificate models.OrganizationCertificate) error
	UpdateFile(ctx context.Context, certificate models.OrganizationCertificate) error
	Delete(ctx context.Context, certificate models.OrganizationCertificate) error
}