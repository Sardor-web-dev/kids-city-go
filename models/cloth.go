package models

import "gorm.io/gorm"

type Cloth struct {
    gorm.Model
    Name        string
    Description string
    Image       string
    Gender      string
    Price       float64
    Size        string
    AuthorID    uint
}
