package models

import "gorm.io/gorm"

type Books struct {
	gorm.Model
	Title      string        `gorm:"not null" json:"title"`
	Author     string        `gorm:"not null" json:"author"`
	Publisher  string        `gorm:"not null" json:"publisher"`
	Year       int           `gorm:"not null" json:"year"`
	InBorrow   bool          `gorm:"not null" json:"in_borrow"`
	Categories []*Categories `gorm:"many2many:book_category"`
	Borrows    []Borrows     `gorm:"foreignKey:BookID"`
}
