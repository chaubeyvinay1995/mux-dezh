package dto

import (
	"gitlab.com/umi7/DezHab_user_backend/common"
	"time"
)

type GetOrganizationInfoRequest struct {
	ID string `json:"id" validate:"required"`
}

// Validate function used to validate the request.
// This is the implementation of the Validation interface.

func (g GetOrganizationInfoRequest) Validate() error {
	return common.Validate.Struct(g)
}

type CreateOrganizationInfoRequest struct {
	UserID                  string `json:"userId" validate:"min=1,max=10,required"`
	EntityType              string `json:"entityType" validate:"min=1,max=10,required"`
	FirstName               string `json:"firstName" validate:"min=1,max=50,required"`
	MiddleName              string `json:"middleName" validate:"min=1,max=10"`
	LastName                string `json:"lastName" validate:"min=1,max=50,required"`
	ContactPersonName       string `json:"contactPersonName" validate:"min=1,max=50,required"`
	PreferredContactInTime  string `json:"inTime" validate:"required"`
	PreferredContactOutTime string `json:"outTime" validate:"required"`
	Gender                  string `json:"gender" validate:"min=1,max=10,required"`
	City                    string `json:"city" validate:"min=1,max=30,required"`
	State                   string `json:"state" validate:"min=1,max=30,required"`
	Country                 string `json:"country" validate:"min=1,max=30,required"`
	Address1                string `json:"address1" validate:"min=1,max=50,required"`
	Address2                string `json:"address2" validate:"min=1,max=50,required"`
	PrimaryContactNumber    string `json:"primaryContactNumber" validate:"min=10,max=12,required"`
	SecondaryContactNumber  string `json:"secondaryContactNumber" validate:"min=10,max=12,required"`
	EmailId                 string `json:"emailId" validate:"min=3,max=20,required,email"`
	WebsiteUrl              string `json:"websiteUrl" validate:"min=5,max=50,required"`
	FacebookUrl             string `json:"facebookUrl" validate:"min=5,max=50,required"`
	About                   string `json:"about" validate:"min=1,max=250,required"`
	ProfessionalExperience  string `json:"professionalExperience" validate:"min=1,max=10,required"`
}

// Validate the function used to validate the Create OrganizationInfo request.
func (g CreateOrganizationInfoRequest) Validate() error {
	return common.Validate.Struct(g)
}

// Need to implement the custom validate to check email already exist in OrganizationInfo table

type UpdateOrganizationInfoRequest struct {
	Id                string `json:"id" validate:"required"`
	EntityType        string `json:"entityType" validate:"min=1,max=10,required"`
	UserID            string `json:"userId" validate:"min=1,max=10,required"`
	ContactPersonName string `json:"contactPersonName" validate:"min=1,max=50,required"`
	FirstName         string `json:"firstName" validate:"min=1,max=50,required"`
	MiddleName        string `json:"middleName" validate:"min=1,max=10"`
	LastName          string `json:"lastName" validate:"min=1,max=50,required"`
	//PreferredContactInTime time.Time `json:"inTime" validate:"required"`
	//PreferredContactOutTime time.Time `json:"outTime" validate:"required"`
	Gender                 string `json:"gender" validate:"min=1,max=10,required"`
	City                   string `json:"city" validate:"min=1,max=30,required"`
	State                  string `json:"state" validate:"min=1,max=30,required"`
	Country                string `json:"country" validate:"min=1,max=30,required"`
	Address1               string `json:"address1" validate:"min=1,max=50,required"`
	Address2               string `json:"address2" validate:"min=1,max=50,required"`
	PrimaryContactNumber   string `json:"primaryContactNumber" validate:"min=10,max=12,required"`
	SecondaryContactNumber string `json:"secondaryContactNumber" validate:"min=10,max=12,required"`
	WebsiteUrl             string `json:"websiteUrl" validate:"min=5,max=50,required"`
	FacebookUrl            string `json:"facebookUrl" validate:"min=5,max=50,required"`
	About                  string `json:"about" validate:"min=1,max=250,required"`
	ProfessionalExperience string `json:"professionalExperience" validate:"min=1,max=10,required"`
}

// Validate the function used to validate the Create OrganizationInfo request.
func (g UpdateOrganizationInfoRequest) Validate() error {
	return common.Validate.Struct(g)
}

// GetOrganizationPrivateInfoResponse response struct
type GetOrganizationPrivateInfoResponse struct {
	Response GetOrganizationPrivateResponse `json:"response"`
}

type GetOrganizationPrivateResponse struct {
	Id                      string    `json:"id"`
	EmailId                 string    `json:"email"`
	EntityType              string    `json:"entityType"`
	UserID                  string    `json:"userId"`
	ContactPersonName       string    `json:"contactPersonName"`
	FirstName               string    `json:"firstName"`
	MiddleName              string    `json:"middleName"`
	LastName                string    `json:"lastName"`
	Gender                  string    `json:"gender"`
	City                    string    `json:"city"`
	State                   string    `json:"state"`
	Country                 string    `json:"country"`
	Address1                string    `json:"address1"`
	Address2                string    `json:"address2"`
	PrimaryContactNumber    string    `json:"primaryContactNumber"`
	SecondaryContactNumber  string    `json:"secondaryContactNumber"`
	WebsiteUrl              string    `json:"websiteUrl"`
	FacebookUrl             string    `json:"facebookUrl"`
	About                   string    `json:"about"`
	PreferredContactInTime  time.Time `json:"inTime"`
	PreferredContactOutTime time.Time `json:"outTime"`
	ProfessionalExperience  string    `json:"professionalExperience"`
}

// GetOrganizationPublicInfoResponse response struct
type GetOrganizationPublicInfoResponse struct {
	Response GetOrganizationPublicResponse `json:"response"`
}

type GetOrganizationPublicResponse struct {
	CompositeScore         uint   `json:"compositeScore"`
	About                  string `json:"about"`
	ProfessionalExperience string `json:"professionalExperience"`
}

// OrganizationInfoResponse response struct

type CreateOrganizationInfoResponse struct {
	Response CreateOrganizationResponse `json:"response"`
}

type CreateOrganizationResponse struct {
	UserId  string `json:"userId"`
	Message string `json:"message"`
}

type UpdateOrganizationInfoResponse struct {
	Response UpdateOrganizationResponse `json:"response"`
}

type UpdateOrganizationResponse struct {
	Id      string `json:"Id"`
	Message string `json:"message"`
}
