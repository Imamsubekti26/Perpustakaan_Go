package models

import "gorm.io/gorm"

type Borrows struct {
	gorm.Model
	BookID uint `gorm:"not null" json:"book_id"`
	UserID uint `gorm:"not null" json:"user_id"`
	Status int  `gorm:"not null" json:"status"`
}
