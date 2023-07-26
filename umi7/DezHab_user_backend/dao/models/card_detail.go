package models

import (
	"time"
)

type CardDetail struct {
	ID     			string 			`gorm:"primary_key;varchar(50)"`
	PaymentId       string 		 	`sql:"type:varchar REFERENCES payments(id)"`			// F.K with Payment info table
	Type 			string		 	`gorm:"varchar(20)"`					// ENUM TYPE i.e CREDIT CARD, DEBIT CARD, UPI ID
	Name 			string 		 	`gorm:"varchar(50)"`
	ExpiryDate  	time.Time
	CardNumber 		string		 	`gorm:"varchar(50)"`
	UpiID 			string       	`gorm:"varchar(50)"`
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       *time.Time   	`sql:"index"`
}
