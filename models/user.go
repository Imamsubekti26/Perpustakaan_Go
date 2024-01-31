package models

import "gorm.io/gorm"

type Users struct {
	gorm.Model
	Surename  string      `gorm:"not null" json:"surename"`
	Username  string      `gorm:"not null" json:"username"`
	Password  string      `gorm:"not null" json:"password"`
	IsAdmin   bool        `gorm:"not null" json:"is_admin"`
	Presences []Presences `gorm:"foreignKey:UserID"`
	Borrows   []Borrows   `gorm:"foreignKey:UserID"`
}

type ResponseUser struct {
	ID       uint
	Surename string
	Username string
	IsAdmin  bool
}
