package models

import (
	"time"
)

type Project struct {
	ID     				string 				`gorm:"primary_key;varchar(50)"`
	OrganizationId  	string 		  		`sql:"type:varchar REFERENCES organization_infos(id)"`		// F.K with OrganizationInfo info table
	ProjectArea     	string				`gorm:"varchar(50)"`
	ProjectField 		string    			`gorm:"varchar(50)"`
	ProjectType 		string   			`gorm:"varchar(50)"`
	Status 				string				`gorm:"varchar(20)"`							// Enum Field
	Description 		string 				`gorm:"varchar(100)"`
	Location 			string				`gorm:"varchar(50)"`
	DeletedAt       	*time.Time 	 		`sql:"index"`
	CreatedAt       	time.Time
	UpdatedAt       	time.Time
}
