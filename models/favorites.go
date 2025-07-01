package models

type Favorite struct {
	ID      uint `gorm:"primaryKey;column:id"`
	ClothID uint `gorm:"column:clothId"`
	UserID  int `gorm:"column:userId"`
	Cloth Cloth `gorm:"foreignKey:ClothID;references:ID"`
	User  User  `gorm:"foreignKey:UserID;references:ID"`
}

func (Favorite) TableName() string {
	return "Favorites"
}
