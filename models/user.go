package models

import (
    "gorm.io/gorm"
)

type User struct {
    gorm.Model
    Name     string `gorm:"type:varchar(100)"`
    Email    string `gorm:"uniqueIndex;type:varchar(100)"`
    Password string `gorm:"type:varchar(255)"`
    Role     string `gorm:"type:varchar(50)"`
}
