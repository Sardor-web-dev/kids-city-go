package models

import (
	"encoding/json"
	"time"
)

type Order struct {
	ID        string          `gorm:"primaryKey" json:"id"`
	Name      string          `json:"name"`
	Surname   string          `json:"surname"`
	Adress    string          `json:"adress"`
	Number    string          `json:"number"`
	Email     string          `json:"email"`
	Payment   string          `json:"payment"`
	Total     int             `json:"total"`
	Items     json.RawMessage `gorm:"type:jsonb" json:"items"`
	CreatedAt time.Time       `gorm:"autoCreateTime" json:"createdAt"`
	Status    string          `gorm:"default:process" json:"status"`
	UserID    int            `json:"userId"`
}

func (Order) TableName() string {
	return "Order"
}
