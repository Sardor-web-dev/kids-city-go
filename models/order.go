package models

type Order struct {
	ID      uint   `gorm:"primaryKey" json:"id"`
	UserID  uint   `json:"user_id"`
	Status  string `json:"status"`
}