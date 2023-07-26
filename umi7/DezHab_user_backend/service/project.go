package service

import (
	"context"

	"gitlab.com/umi7/DezHab_user_backend/dao/models"
)

type Project interface {
	Get(ctx context.Context, organizationId string) ([]models.Project, error)

	Create(ctx context.Context, organizationId, projectArea, projectField, projectType, status,
		description, location string) (projectId string, err error)

	Update(ctx context.Context, Id, projectArea, projectField, projectType, status,
		description, location string) error

	Delete(ctx context.Context, Id string) error
}

// Command to generate the mock implementation of the BookKeeping interface
//go:generate mockgen -source=book-keeping.go -destination=mock-impl/book_keeping_impl.go -package=mock_impl
