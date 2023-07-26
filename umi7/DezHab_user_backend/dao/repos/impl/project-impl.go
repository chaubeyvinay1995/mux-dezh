package impl

import (
	"context"
	"errors"
	"gitlab.com/umi7/DezHab_user_backend/dao/core/impl"
	"gitlab.com/umi7/DezHab_user_backend/dao/models"
	"gitlab.com/umi7/DezHab_user_backend/dao/repos"
	"time"
)

var updateProjectQuery = "id=?"

var ProjectInfo repos.ProjectInfo

type ProjectInfoImpl struct{}


func init() {
	ProjectInfo = ProjectInfoImpl{}
}

func (projectInfoImpl ProjectInfoImpl) Search(ctx context.Context, organizationId string) (projectInfo []models.Project, err error) {
	var projectList []models.Project
	if err = impl.Read.Find(&projectList, models.Project{OrganizationId: organizationId}, &projectInfo); err != nil {
		return
	}
	if len(projectList) == 0 {
		err = errors.New(ErrNoData)
		return
	}
	projectInfo = projectList
	return
}


func (projectInfoImpl ProjectInfoImpl) Create(ctx context.Context, projectInfo models.Project) (err error) {
	projectInfo.CreatedAt = time.Now().UTC()
	projectInfo.UpdatedAt = projectInfo.CreatedAt
	err = impl.Write.Create(&projectInfo)
	return
}

func (projectInfoImpl ProjectInfoImpl) Update(ctx context.Context, projectInfo models.Project) (err error) {
	projectInfo.UpdatedAt = time.Now().UTC()
	updateValues := map[string]interface{}{
		"project_area": projectInfo.ProjectArea, "project_type": projectInfo.ProjectType,
		"project_field": projectInfo.ProjectField, "status": projectInfo.Status,
		"description": projectInfo.Description, "location": projectInfo.Location,
	}
	err = impl.Write.Update(&projectInfo, updateProjectQuery, projectInfo.ID, updateValues)
	return
}

func (projectInfoImpl ProjectInfoImpl) Delete(ctx context.Context, projectInfo models.Project) (err error) {
	err = impl.Write.Delete(&projectInfo)
	return
}
