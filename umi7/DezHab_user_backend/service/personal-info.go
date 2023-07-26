package service

import (
	"context"
	"gitlab.com/umi7/DezHab_user_backend/dao/models"
	"time"
)

type PersonalInfo interface {
	Get(ctx context.Context, id string) (models.UserPersonalInfo, error)

	Create(ctx context.Context, userId, firstName, middleName, lastName, email, primaryContact, secondaryContact,
		pinCode, address, residentState, userAliasName, registerEdas, googleID, webProfile,
		linkedIn string, dateOfBirth time.Time, isBlocked, isVerified, pendingPersonalDetail,
		pendingProfessionalDetail bool) (id string, err error)

	Update(ctx context.Context, id, firstName, middleName, lastName, primaryContact, secondaryContact,
		pinCode, address, residentState, userAliasName, registerEdas, googleID, webProfile,
		linkedIn string, dateOfBirth time.Time, isBlocked, isVerified, pendingPersonalDetail,
		pendingProfessionalDetail bool) (err error)

	Delete(ctx context.Context, id string) error
}
