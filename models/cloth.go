package models

import (
	"github.com/lib/pq"
)

type Cloth struct {
	ID          uint           `gorm:"column:id" json:"id"`
	Name        string         `gorm:"column:name" json:"name"`
	Description string         `gorm:"column:description" json:"description"`
	Gender      string         `gorm:"column:gender" json:"gender"`
	Image       string         `gorm:"column:image" json:"image"`
	AuthorID    uint           `gorm:"column:authorId" json:"authorId"`
	Price       float64        `gorm:"column:price" json:"price"`
	Size        pq.StringArray `gorm:"type:text[];column:size" json:"size"`
}

func (Cloth) TableName() string {
	return "Cloth"
}
