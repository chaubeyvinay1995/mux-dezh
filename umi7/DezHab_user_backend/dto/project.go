package dto

import (
	"gitlab.com/umi7/DezHab_user_backend/common"
	"gitlab.com/umi7/DezHab_user_backend/dao/models"
)

// GetProjectRequest struct used to validate the get request for Project.
type GetProjectRequest struct {
	OrganizationId string `json:"organizationId" validate:"required"`
}

// Validate -->Call Validate method to validate the GetProject Struct
func(p GetProjectRequest) Validate() error{
	return common.Validate.Struct(p)
}

// CreateProjectRequest struct used to validate the post request for Project.
type CreateProjectRequest struct {
	OrganizationId		string 		`json:"organizationId" validate:"required"`
	ProjectArea 		string 		`json:"projectArea" validate:"required"`
	ProjectField		string 		`json:"projectField" validate:"required"`
	ProjectType			string		`json:"projectType" validate:"required"`
	Status 				string		`json:"status" validate:"required"`
	Description         string 		`json:"description" validate:"required"`
	Location 			string		`json:"location"`
}

// Validate -->Call Validate method to validate the CreateProjectRequest Struct
func(p CreateProjectRequest) Validate() error{
	return common.Validate.Struct(p)
}

// UpdateProjectRequest struct used to validate the patch request for Project.
type UpdateProjectRequest struct {
	Id					string 		`json:"id" validate:"required"`
	ProjectArea 		string 		`json:"projectArea" validate:"min=1,max=50,required"`
	ProjectField		string 		`json:"projectField" validate:"min=1,max=50,required"`
	ProjectType			string		`json:"projectType" validate:"min=1,max=20,required"`
	Status 				string		`json:"status" validate:"min=1,max=20,required"`
	Description         string 		`json:"description" validate:"max=100,required"`
	Location 			string		`json:"location"`
}

// Validate --> Call Validate method to validate the UpdateProject Struct.
func(p UpdateProjectRequest) Validate() error{
	return common.Validate.Struct(p)
}


// DeleteProjectRequest struct used to validate the delete request for Project.
type DeleteProjectRequest struct {
	Id 		string 	`json:"id" validate:"required"`
}

// Validate -->Call Validate method to validate the DeleteProjectRequest Struct
func(p DeleteProjectRequest) Validate() error{
	return common.Validate.Struct(p)
}


// GetProjectPrivateInfoResponse response struct
type GetProjectPrivateInfoResponse struct {
	Response []models.Project `json:"response"`
}

// GetProjectPublicInfoResponse response struct
type GetProjectPublicInfoResponse struct {
	Response []ProjectPublicResponse `json:"response"`
}

type ProjectPublicResponse struct {
	OrganizationId		string 		`json:"OrganizationId"`
	Id 					string		`json:"ID"`
	ProjectType			string		`json:"ProjectType"`
}


type CreateProjectResponse struct {
	Response CreateProjectData `json:"response"`
}

type CreateProjectData struct {
	OrganizationId 			string     	`json:"id"`
	Message 				string		`json:"message"`
}


type UpdateProjectResponse struct {
	Response UpdateProjectData `json:"response"`
}

type UpdateProjectData struct {
	Id 			string     	`json:"id"`
	Message 	string		`json:"message"`
}


type DeleteProjectResponse struct {
	Response DeleteProjectData `json:"response"`
}

type DeleteProjectData struct {
	Id 			string     	`json:"id"`
	Message 	string		`json:"message"`
}
