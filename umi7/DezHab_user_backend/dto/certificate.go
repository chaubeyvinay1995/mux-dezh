package dto

import (
	"gitlab.com/umi7/DezHab_user_backend/common"
	"gitlab.com/umi7/DezHab_user_backend/dao/models"
	"mime/multipart"
)

// GetCertificateRequest struct used to validate the Get Academic Request
type GetCertificateRequest struct {
	UserId string `json:"userId" validate:"required"`
}

// Validate --> Call Validate method used to validate the GetAcademicRequest
func (ca GetCertificateRequest) Validate() error {
	return common.Validate.Struct(ca)
}

// CreateCertificateRequest struct used to validate the Create Certificate Request.
type CreateCertificateRequest struct {
	OrganizationId string `json:"organization_id" validate:"required"`
	CourseName     string `json:"course_name" validate:"required"`
	Url            string `json:"url" validate:"required"`
	Institute      string `json:"institute" validate:"required"`
}

// Validate struct used to validate the CreateCertificateRequest
func (ca CreateCertificateRequest) Validate() error {
	return common.Validate.Struct(ca)
}

// UpdateCertificateRequest struct used to validate the Update Certificate Request.
type UpdateCertificateRequest struct {
	Id         string `json:"id" validate:"required"`
	CourseName string `json:"course_name" validate:"required"`
	Url        string `json:"url" validate:"required"`
	Institute  string `json:"institute" validate:"required"`
}

// Validate struct used to validate the UpdateCertificateRequest
func (ca UpdateCertificateRequest) Validate() error {
	return common.Validate.Struct(ca)
}

// DeleteCertificateRequest struct used to validate Delete Academic Request
type DeleteCertificateRequest struct {
	Id string `json:"id" validate:"required"`
}

// Validate --> Call Validate method used to validate the DeleteCertificateRequest
func (ca DeleteCertificateRequest) Validate() error {
	return common.Validate.Struct(ca)
}

// UpdateCertificateFileRequest struct used to validate the update certificate file Request
type UpdateCertificateFileRequest struct {
	Id         string         `json:"id" validate:"required"`
	FileString multipart.Form `json:"fileString" validate:"required"`
}

// Validate --> Call Validate method used to validate the DeleteCertificateRequest
func (ca UpdateCertificateFileRequest) Validate() error {
	return common.Validate.Struct(ca)
}

// GetCertificatePublicDetailResponse used to send the get response.
type GetCertificatePublicDetailResponse struct {
	Response []models.OrganizationCertificate `json:"response"`
}

// GetCertificatePrivateDetailResponse used to send the get response.
type GetCertificatePrivateDetailResponse struct {
	Response []CertificatePrivateResponse `json:"response"`
}

type CertificatePrivateResponse struct {
	CertificateId string `json:"CertificateId"`
	CourseName    string `json:"CourseName"`
}

// CreateCertificateResponse struct used to send create Certificate response.
type CreateCertificateResponse struct {
	Response CreateCertificateData `json:"response"`
}

type CreateCertificateData struct {
	CertificateId string `json:"id"`
	Message       string `json:"message"`
}

// UpdateCertificateResponse struct used to send update certificate response.
type UpdateCertificateResponse struct {
	Response UpdateCertificateData `json:"response"`
}

type UpdateCertificateData struct {
	CertificateId string `json:"id"`
	Message       string `json:"message"`
}

// DeleteCertificateResponse struct used to send delete certificate response.
type DeleteCertificateResponse struct {
	Response DeleteCertificateData `json:"response"`
}

type DeleteCertificateData struct {
	Message string `json:"message"`
}
