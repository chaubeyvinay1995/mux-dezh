package impl

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"gitlab.com/umi7/DezHab_user_backend/common/logger"
	"gitlab.com/umi7/DezHab_user_backend/dao/models"
	"gitlab.com/umi7/DezHab_user_backend/dao/repos/impl"
	"gitlab.com/umi7/DezHab_user_backend/service"
	"time"
)

var PersonalInfo service.PersonalInfo

type PersonalInfoImpl struct{}

func init() {
	PersonalInfo = PersonalInfoImpl{}
}

func (p PersonalInfoImpl) Get(ctx context.Context, id string) (personalInfo models.UserPersonalInfo, err error) {
	personalInfo, err = impl.PersonalInfo.Search(ctx, id)
	if err != nil {
		logger.Error(ctx, err)
		return
	}
	return

}

func (p PersonalInfoImpl) Create(ctx context.Context, userId, firstName, middleName, lastName, email,
	primaryContact, secondaryContact, pinCode, address, residentState, userAliasName, registerEdas, googleID,
	webProfile, linkedIn string, dateOfBirth time.Time, isBlocked, isVerified, pendingPersonalDetail,
	pendingProfessionalDetail bool) (id string, err error) {
	id = uuid.New().String()

	personalInfo := models.UserPersonalInfo{
		ID: id, UserId: userId, FirstName: firstName, MiddleName: middleName, LastName: lastName,
		Email: email, PrimaryContact: primaryContact, SecondaryContact: secondaryContact, PinCode: pinCode,
		Address: address, ResidentState: residentState, UserAliasName: userAliasName, RegisterEdas: registerEdas,
		GoogleID: googleID, WebProfile: webProfile, LinkedIn: linkedIn, DateOfBirth: dateOfBirth,
		IsBlocked: isBlocked, IsVerified: isVerified, PendingPersonalDetail: pendingPersonalDetail,
		PendingProfessionalDetail: pendingProfessionalDetail,
	}
	err = impl.PersonalInfo.Create(ctx, personalInfo)
	if err != nil {
		logger.Error(ctx, fmt.Sprintf("Error occurred while creating entry on personal info : %v", err))
	}
	return

}

func (p PersonalInfoImpl) Delete(ctx context.Context, id string) (err error) {
	personalInfo := models.UserPersonalInfo{
		ID: id,
	}
	err = impl.PersonalInfo.Delete(ctx, personalInfo)
	if err != nil {
		logger.Error(ctx, fmt.Sprintf("Error occurred while creating entry on organization info : %v", err))
	}
	return

}

// Update PersonalInfo
func (p PersonalInfoImpl) Update(ctx context.Context, id, firstName, middleName, lastName, primaryContact,
	secondaryContact, pinCode, address, residentState, userAliasName, registerEdas, googleID, webProfile,
	linkedIn string, dateOfBirth time.Time, isBlocked, isVerified, pendingPersonalDetail,
	pendingProfessionalDetail bool) (err error) {
	personalInfo := models.UserPersonalInfo{
		FirstName: firstName, MiddleName: middleName, LastName: lastName,
		PrimaryContact: primaryContact, SecondaryContact: secondaryContact, PinCode: pinCode, Address: address,
		ResidentState: residentState, UserAliasName: userAliasName, RegisterEdas: registerEdas,
		GoogleID: googleID, WebProfile: webProfile, LinkedIn: linkedIn, DateOfBirth: dateOfBirth,
		IsBlocked: isBlocked, IsVerified: isVerified, PendingPersonalDetail: pendingPersonalDetail,
		PendingProfessionalDetail: pendingProfessionalDetail, ID: id,
	}
	err = impl.PersonalInfo.Update(ctx, personalInfo)
	if err != nil {
		logger.Error(ctx, err)
	}
	return
}
