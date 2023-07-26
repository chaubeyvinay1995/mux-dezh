package dto

import (
	"gitlab.com/umi7/DezHab_user_backend/common"
)

// GetUserPersonalInfoRequest used to validate the get user personal info request
type GetUserPersonalInfoRequest struct {
	ID string `json:"id" validate:"required"`
}

// Validate function used to validate the request.
// This is the implementation of the Validation interface.

func (g GetUserPersonalInfoRequest) Validate() error {
	return common.Validate.Struct(g)
}

// GetPrivatePersonalInfoResponse response struct
type GetPrivatePersonalInfoResponse struct {
	Response GetPrivatePersonalInfoData `json:"response"`
}

type GetPrivatePersonalInfoData struct {
	Id                        string `json:"id"`
	UserId                    string `json:"userId"`
	CompositeScore            uint   `json:"compositeScore"`
	FirstName                 string `json:"firstName"`
	MiddleName                string `json:"middleName"`
	LastName                  string `json:"lastName"`
	Email                     string `json:"email"`
	PrimaryContact            string `json:"primaryContact"`
	SecondaryContact          string `json:"secondaryContact"`
	PinCode                   string `json:"pinCode"`
	Address                   string `json:"address"`
	ResidentState             string `json:"residentState"`
	UserAliasName             string `json:"UserAliasName"`
	RegisterEdas              string `json:"registerEdas"`
	IsBlocked                 bool   `json:"IsBlocked"`
	IsVerified                bool   `json:"isVerified"`
	PendingPersonalDetail     bool   `json:"pendingPersonalDetail"`
	PendingProfessionalDetail bool   `json:"pendingProfessionalDetail"`
	GoogleID                  string `json:"googleID"`
	WebProfile                string `json:"webProfile"`
	LinkedIn                  string `json:"linkedIn"`
	DateOfBirth               string `json:"dateOfBirth"`
}

// GetPublicPersonalInfoResponse response struct
type GetPublicPersonalInfoResponse struct {
	Response GetPublicPersonalInfoData `json:"response"`
}

type GetPublicPersonalInfoData struct {
	Message string `json:"message"`
}

// CreatePersonalInfoRequest struct used to validate the Create Personal Info Request
type CreatePersonalInfoRequest struct {
	UserID                    string `json:"userId" validate:"min=1,max=50,required"`
	FirstName                 string `json:"firstName" validate:"min=1,max=50,required"`
	MiddleName                string `json:"middleName" validate:"min=1,max=80,required"`
	LastName                  string `json:"lastName" validate:"min=1,max=80,required"`
	Email                     string `json:"email" validate:"min=1,max=80,required"`
	PrimaryContact            string `json:"primaryContact" validate:"min=1,max=20,required"`
	SecondaryContact          string `json:"secondaryContact" validate:"min=1,max=20,required"`
	PinCode                   string `json:"pinCode" validate:"min=6,max=6,required"`
	Address                   string `json:"address" validate:"min=1,max=255,required"`
	ResidentState             string `json:"residentState" validate:"min=1,max=50,required"`
	UserAliasName             string `json:"userAliasName" validate:"min=1,max=50,required"`
	RegisterEdas              string `json:"registerEdas" validate:"min=1,max=6,required"`
	GoogleID                  string `json:"googleId" validate:"min=1,max=255,required"`
	WebProfile                string `json:"webProfile" validate:"min=1,max=255,required"`
	LinkedIn                  string `json:"linkedIn" validate:"min=1,max=255,required"`
	DateOfBirth               string `json:"dateOfBirth" validate:"min=10,max=10,required"`
	IsBlocked                 bool   `json:"isBlocked"`
	IsVerified                bool   `json:"isVerified"`
	PendingPersonalDetail     bool   `json:"pendingPersonalDetail"`
	PendingProfessionalDetail bool   `json:"pendingProfessionalDetail"`
}

// Validate function used to validate the request.
// This is the implementation of the Validation interface.

func (g CreatePersonalInfoRequest) Validate() error {
	return common.Validate.Struct(g)
}

// CreatePublicPersonalInfoResponse response struct
type CreatePublicPersonalInfoResponse struct {
	Response CreatePublicPersonalInfoData `json:"response"`
}

type CreatePublicPersonalInfoData struct {
	Id      string `json:"id"`
	Message string `json:"message"`
}

// DeletePersonalInfoRequest struct used to validate the delete request for Project.
type DeletePersonalInfoRequest struct {
	Id string `json:"id" validate:"required"`
}

// Validate -->Call Validate method to validate the DeleteProjectRequest Struct
func (p DeletePersonalInfoRequest) Validate() error {
	return common.Validate.Struct(p)
}

type DeletePersonalInfoResponse struct {
	Response DeletePersonalInfoData `json:"response"`
}

type DeletePersonalInfoData struct {
	Message string `json:"message"`
}

// UpdatePersonalInfoRequest struct used to validate the Update Personal Info Request
type UpdatePersonalInfoRequest struct {
	Id                        string `json:"id" validate:"min=1,max=50,required"`
	FirstName                 string `json:"firstName" validate:"min=1,max=50,required"`
	MiddleName                string `json:"middleName" validate:"min=1,max=80,required"`
	LastName                  string `json:"lastName" validate:"min=1,max=80,required"`
	PrimaryContact            string `json:"primaryContact" validate:"min=1,max=20,required"`
	SecondaryContact          string `json:"secondaryContact" validate:"min=1,max=20,required"`
	PinCode                   string `json:"pinCode" validate:"min=6,max=6,required"`
	Address                   string `json:"address" validate:"min=1,max=255,required"`
	ResidentState             string `json:"residentState" validate:"min=1,max=50,required"`
	UserAliasName             string `json:"userAliasName" validate:"min=1,max=50,required"`
	RegisterEdas              string `json:"registerEdas" validate:"min=1,max=6,required"`
	GoogleID                  string `json:"googleId" validate:"min=1,max=255,required"`
	WebProfile                string `json:"webProfile" validate:"min=1,max=255,required"`
	LinkedIn                  string `json:"linkedIn" validate:"min=1,max=255,required"`
	DateOfBirth               string `json:"dateOfBirth" validate:"min=10,max=10,required"`
	IsBlocked                 bool   `json:"isBlocked"`
	IsVerified                bool   `json:"isVerified"`
	PendingPersonalDetail     bool   `json:"pendingPersonalDetail"`
	PendingProfessionalDetail bool   `json:"pendingProfessionalDetail"`
}

// Validate function used to validate the request.
// This is the implementation of the Validation interface.

func (g UpdatePersonalInfoRequest) Validate() error {
	return common.Validate.Struct(g)
}

type UpdatePersonalInfoResponse struct {
	Response UpdatePersonalInfoData `json:"response"`
}

type UpdatePersonalInfoData struct {
	Message string `json:"message"`
}
