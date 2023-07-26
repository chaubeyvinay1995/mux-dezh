package models

import (
	"time"
)

type OrganizationInfo struct {
	ID     					  string 			`gorm:"primary_key;varchar(50)"`
	UserId                    string 			`gorm:"varchar(50)"`
	CompositeScore            uint
	FirstName                 string 			`gorm:"varchar(50)"`
	MiddleName                string 			`gorm:"varchar(50)"`
	LastName                  string 			`gorm:"varchar(50)"`
	ProfessionalExperience    string 			`gorm:"varchar(10)"`
	About                     string 			`gorm:"type:text"`
	EntityType                string 			`gorm:"varchar(20)"`		// it is of Enum type
	ContactPersonName         string 			`gorm:"varchar(50)"`
	Gender                    string 			`gorm:"varchar(10)"`		// Male, Female and Other
	State                     string 			`gorm:"varchar(50)"`
	Country                   string 			`gorm:"varchar(50)"`
	City                      string 			`gorm:"varchar(50)"`
	Address1                  string 			`gorm:"varchar(100)"`
	Address2                  string 			`gorm:"varchar(100)"`
	PrimaryContactNumber      string 			`gorm:"varchar(20)"`
	SecondaryContactNumber    string 			`gorm:"varchar(20)"`
	EmailId                   string 			`gorm:"varchar(50)"`
	WebsiteUrl                string 			`gorm:"varchar(50)"`
	FacebookUrl               string 			`gorm:"varchar(50)"`
	PreferredContactInTime    time.Time
	PreferredContactOutTime   time.Time
	DeletedAt                 *time.Time 		`sql:"index"`
	CreatedAt                 time.Time
	UpdatedAt                 time.Time

}
