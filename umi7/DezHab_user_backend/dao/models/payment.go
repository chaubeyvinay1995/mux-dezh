package models

import (
	"time"
)

type Payment struct {
	ID     				string 			`gorm:"primary_key;varchar(50)"`
	Type 				string  		`gorm:"varchar(20)"`					// ENUM FIELD
	Name 				string 			`gorm:"varchar(50)"`
	AccountNumber 		string 			`gorm:"varchar(50)"`
	BankName 			string 			`gorm:"varchar(50)"`
	DeletedAt       	*time.Time 	 	`sql:"index"`
	CreatedAt       	time.Time
	UpdatedAt       	time.Time
}
