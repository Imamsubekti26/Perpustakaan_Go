package models

import "time"

type Presences struct {
	ID        uint `gorm:"primaryKey;AutoIncrement" json:"id"`
	UserID    uint
	CreatedAt *time.Time `gorm:"index" json:"created_at"`
}
