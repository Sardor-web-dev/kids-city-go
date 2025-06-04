package models

type Group struct {
	ID   uint   `gorm:"primaryKey" json:"id"`
	Name string `json:"name"`
}