package models

type User struct {
	ID        int        `gorm:"primaryKey;autoIncrement" json:"id"`
	Name      *string    `gorm:"column:name" json:"name"`
	Email     string     `gorm:"unique;column:email" json:"email"`
	Password  *string    `gorm:"column:password" json:"password,omitempty"`
	Image     *string    `gorm:"column:image" json:"image,omitempty"`
	Role      string     `gorm:"column:role;default:USER" json:"role"`
	IsBlocked bool       `gorm:"column:isBlocked;default:false" json:"isBlocked"`
	Clothes   []Cloth    `gorm:"foreignKey:AuthorID"`
	Orders    []Order    `gorm:"foreignKey:UserID"`
	Favorites []Favorite `gorm:"foreignKey:UserID"`
	Accounts  []Account  `gorm:"foreignKey:UserID"`
	Sessions  []Session  `gorm:"foreignKey:UserID"`
}

func (User) TableName() string {
	return "User"
}
