package models

import "time"

type VerificationToken struct {
	Identifier string    `gorm:"column:identifier" json:"identifier"`
	Token      string    `gorm:"primaryKey;column:token" json:"token"`
	Expires    time.Time `gorm:"column:expires" json:"expires"`
}

func (VerificationToken) TableName() string {
	return "VerificationToken"
}
