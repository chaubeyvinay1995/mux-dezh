package dto

import (
	"gitlab.com/umi7/DezHab_user_backend/common"
	"gitlab.com/umi7/DezHab_user_backend/dao/models"
)

// GetAcademicRequest struct used to validate the Get Academic Request
type GetAcademicRequest struct {
	OrganizationId		string 		`json:"organizationId" validate:"required"`
}

// Validate --> Call Validate method used to validate the GetAcademicRequest
func (ca GetAcademicRequest) Validate() error{
	return common.Validate.Struct(ca)
}

// CreateAcademicRequest struct used to validate the Create Academic Request
type CreateAcademicRequest struct {
	OrganizationId		string 		`json:"organizationId" validate:"required"`
	Degree 				string 		`json:"degree" validate:"required"`
	InstituteName       string 		`json:"institute_name" validate:"required"`
	CertificateLink     string		`json:"certificate_link" validate:"required"`
}

// Validate --> Call Validate method used to validate the CreateAcademicRequest
func (ca CreateAcademicRequest) Validate() error{
	return common.Validate.Struct(ca)
}

// UpdateAcademicDetailRequest struct used to validate the Update Academic Request
type UpdateAcademicDetailRequest struct {
	Id					string 		`json:"id" validate:"required"`
	Degree 				string 		`json:"degree" validate:"required"`
	InstituteName       string 		`json:"institute_name" validate:"required"`
	CertificateLink     string		`json:"certificate_link" validate:"required"`
}

// Validate --> Call Validate method used to validate the UpdateAcademicDetailRequest
func (ca UpdateAcademicDetailRequest) Validate() error{
	return common.Validate.Struct(ca)
}

// DeleteAcademicRequest struct used to validate Delete Academic Request
type DeleteAcademicRequest struct {
	Id		string 			`json:"id" validate:"required"`
}

// Validate --> Call Validate method used to validate the DeleteAcademicRequest
func (ca DeleteAcademicRequest) Validate() error{
	return common.Validate.Struct(ca)
}

// GetAcademicDetailPublicResponse struct used to send get academic detail response.
type GetAcademicDetailPublicResponse struct {
	Response []models.OrganizationAcademicDetail `json:"response"`
}

// GetAcademicDetailPrivateResponse struct used to send get academic detail response.
type GetAcademicDetailPrivateResponse struct {
	Response []AcademicPrivateResponse `json:"response"`
}

type AcademicPrivateResponse struct {
	OrganizationId 			string 			`json:"OrganizationId"`
	Id 						string			`json:"Id"`
	Degree					string			`json:"Degree"`
}

// CreateAcademicResponse struct used to send create academic response.
type CreateAcademicResponse struct {
	Response 		CreateAcademicData 		`json:"response"`
}

type CreateAcademicData struct {
	AcademicId 				string     		`json:"id"`
	Message 				string			`json:"message"`
}

// UpdateAcademicResponse struct used to send update academic response.
type UpdateAcademicResponse struct {
	Response 		UpdateAcademicData 		`json:"response"`
}

type UpdateAcademicData struct {
	AcademicId 				string     		`json:"id"`
	Message 				string			`json:"message"`
}

// DeleteAcademicResponse struct used to send delete academic response.
type DeleteAcademicResponse struct {
	Response 		DeleteAcademicData 		`json:"response"`
}

type DeleteAcademicData struct {
	Message 				string			`json:"message"`
}

