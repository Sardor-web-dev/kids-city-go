package models

import "time"

type Session struct {
	ID           string    `gorm:"primaryKey;column:id" json:"id"`
	SessionToken string    `gorm:"column:sessionToken" json:"sessionToken"`
	UserID       uint      `gorm:"column:userId" json:"userId"`
	Expires      time.Time `gorm:"column:expires" json:"expires"`
}

func (Session) TableName() string {
	return "Session"
}
