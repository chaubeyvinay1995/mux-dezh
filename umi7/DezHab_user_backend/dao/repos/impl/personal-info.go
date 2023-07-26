package impl

import (
	"context"
	"gitlab.com/umi7/DezHab_user_backend/dao/core/impl"
	"gitlab.com/umi7/DezHab_user_backend/dao/models"
	"gitlab.com/umi7/DezHab_user_backend/dao/repos"
	"time"
)

const (
	updatePersonalInfoQuery = "id=?"
)

var PersonalInfo repos.PersonalInfo

type PersonalInfoImpl struct{}

func init() {
	PersonalInfo = PersonalInfoImpl{}
}

func (personalInfoImpl PersonalInfoImpl) Search(ctx context.Context, id string) (projectInfo models.UserPersonalInfo, err error) {
	var projectInfoDetail models.UserPersonalInfo
	if err = impl.Read.First(&projectInfoDetail, models.UserPersonalInfo{ID: id}); err != nil {
		return
	}
	projectInfo = projectInfoDetail
	return
}

func (personalInfoImpl PersonalInfoImpl) Create(ctx context.Context, personalInfo models.UserPersonalInfo) (err error) {
	personalInfo.CreatedAt = time.Now().UTC()
	personalInfo.UpdatedAt = personalInfo.CreatedAt
	err = impl.Write.Create(&personalInfo)
	return
}

func (personalInfoImpl PersonalInfoImpl) Delete(ctx context.Context, personalInfo models.UserPersonalInfo) (err error) {
	err = impl.Write.Delete(&personalInfo)
	return
}

func (personalInfoImpl PersonalInfoImpl) Update(ctx context.Context, personalInfo models.UserPersonalInfo) (err error) {
	personalInfo.UpdatedAt = time.Now().UTC()
	updateValues := map[string]interface{}{
		"first_name": personalInfo.FirstName, "middle_name": personalInfo.MiddleName,
		"last_name": personalInfo.LastName, "primary_contact": personalInfo.PrimaryContact,
		"secondary_contact": personalInfo.SecondaryContact, "pin_code": personalInfo.PinCode,
		"address": personalInfo.Address, "resident_state": personalInfo.ResidentState,
		"user_alias_name": personalInfo.UserAliasName, "register_edas": personalInfo.RegisterEdas,
		"google_id": personalInfo.GoogleID, "web_profile": personalInfo.WebProfile,
		"linked_in": personalInfo.LinkedIn, "is_blocked": personalInfo.IsBlocked,
		"is_verified": personalInfo.IsVerified, "pending_personal_detail": personalInfo.PendingPersonalDetail,
		"pending_professional_detail": personalInfo.PendingProfessionalDetail, "date_of_birth": personalInfo.DateOfBirth,
	}
	err = impl.Write.Update(&personalInfo, updatePersonalInfoQuery, personalInfo.ID, updateValues)
	return
}
