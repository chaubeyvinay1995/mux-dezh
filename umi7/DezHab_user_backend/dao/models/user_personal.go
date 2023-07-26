package models

import (
	"time"
)

type UserPersonalInfo struct {
	ID                        string `gorm:"primary_key;varchar(50)"`
	UserId                    string `gorm:"varchar(50), unique;not null;"`
	CompositeScore            uint
	FirstName                 string     `gorm:"varchar(80)"`
	MiddleName                string     `gorm:"varchar(80)"`
	LastName                  string     `gorm:"varchar(80)"`
	Email                     string     `gorm:"varchar(100)"`
	PrimaryContact            string     `gorm:"varchar(20)"`
	SecondaryContact          string     `gorm:"varchar(20)"`
	PinCode                   string     `gorm:"varchar(6)"`
	Address                   string     `gorm:"type:text"`
	ResidentState             string     `gorm:"varchar(50)"`
	UserAliasName             string     `gorm:"varchar(50)"`
	RegisterEdas              string     `gorm:"varchar(6)"`
	IsBlocked                 bool       `gorm:"type:bool;default:false"`
	IsVerified                bool       `gorm:"type:bool;default:false"`
	PendingPersonalDetail     bool       `gorm:"type:bool;default:false"`
	PendingProfessionalDetail bool       `gorm:"type:bool;default:false"`
	GoogleID                  string     `gorm:"varchar(255)"`
	WebProfile                string     `gorm:"type:text"`
	LinkedIn                  string     `gorm:"type:text"`
	DeletedAt                 *time.Time `sql:"index"`
	DateOfBirth               time.Time
	CreatedAt                 time.Time
	UpdatedAt                 time.Time
}
