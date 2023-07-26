package repos

import (
	"context"
	"gitlab.com/umi7/DezHab_user_backend/dao/models"
)


type ProjectInfo interface {
	Search(context.Context, string) ([]models.Project, error)
	Create(context.Context, models.Project) error
	Update(context.Context, models.Project) error
	Delete(context.Context, models.Project) error
}

// Command to generate the mock implementation of the BookKeeping interface
//go:generate mockgen -source=book_keeping.go -destination=mock-impl/project-info-impl.go -package=mock_impl
