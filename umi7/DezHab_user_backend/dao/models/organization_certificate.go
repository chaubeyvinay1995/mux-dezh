package models

import (
	"time"
)

type OrganizationCertificate struct {
	ID     				string 		  `gorm:"primary_key;varchar(50)"`
	OrganizationId      string 		  `sql:"type:varchar REFERENCES organization_infos(id)"`  	// F.K with OrganizationInfo info table
	CourseName    		string 		  `gorm:"varchar(50)"`
	Url 		  		string  	  `gorm:"varchar(50)"`
	Institute     		string  	  `gorm:"varchar(50)"`
	CreatedAt           time.Time
	UpdatedAt           time.Time
	DeletedAt           *time.Time 	   `sql:"index"`
}
