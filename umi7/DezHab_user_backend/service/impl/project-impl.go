package impl

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"gitlab.com/umi7/DezHab_user_backend/common/logger"
	"gitlab.com/umi7/DezHab_user_backend/dao/models"
	"gitlab.com/umi7/DezHab_user_backend/dao/repos/impl"
	"gitlab.com/umi7/DezHab_user_backend/service"
)

var ProjectInfo service.Project

type ProjectImpl struct{}

func init() {
	ProjectInfo = ProjectImpl{}
}


func (p ProjectImpl) Get(ctx context.Context, organizationId string) (projectInfo []models.Project, err error) {
	projectInfo, err = impl.ProjectInfo.Search(ctx, organizationId)
	if err != nil {
		logger.Error(ctx, err)
		return
	}
	return

}

func (p ProjectImpl) Create(ctx context.Context, organizationId, projectArea, projectField, projectType,
	status, description, location string) (projectId string, err error){
	// Create UUID
	projectId = uuid.New().String()
	projectInfo := models.Project{
		ID: projectId, OrganizationId: organizationId, ProjectArea: projectArea,
		ProjectField: projectField, ProjectType: projectType, Status: status, Description: description,
		Location: location,
	}
	err = impl.ProjectInfo.Create(ctx, projectInfo)
	if err != nil {
		logger.Error(ctx, fmt.Sprintf("Error occurred while creating entry on organization info : %v", err))
	}
	return

}


func (p ProjectImpl) Update(ctx context.Context, id, projectArea, projectField, projectType,
	status, description, location string) (err error){
	projectInfo := models.Project{
		ID: id, ProjectArea: projectArea, ProjectField: projectField, ProjectType: projectType,
		Status: status, Description: description, Location: location,
	}
	err = impl.ProjectInfo.Update(ctx, projectInfo)
	if err != nil {
		logger.Error(ctx, fmt.Sprintf("Error occurred while creating entry on organization info : %v", err))
	}
	return

}


func (p ProjectImpl) Delete(ctx context.Context, id string) (err error){
	projectInfo := models.Project{
		ID: id,
	}
	err = impl.ProjectInfo.Delete(ctx, projectInfo)
	if err != nil {
		logger.Error(ctx, fmt.Sprintf("Error occurred while creating entry on organization info : %v", err))
	}
	return

}