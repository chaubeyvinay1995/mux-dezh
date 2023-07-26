package models

import (
	"time"
)


type OrganizationAcademicDetail struct {
	ID     				string 			`gorm:"primary_key;varchar(50)"`
	OrganizationId      string 	    	`sql:"type:varchar REFERENCES organization_infos(id)"`		// F.K with OrganizationInfo info table
	Degree             	string      	`gorm:"varchar(50)"`
	InstituteName      	string      	`gorm:"varchar(50)"`
	CertificateLink    	string      	`gorm:"varchar(50)"`
	CreatedAt           time.Time
	UpdatedAt           time.Time
	DeletedAt          *time.Time   	`sql:"index"`
}
