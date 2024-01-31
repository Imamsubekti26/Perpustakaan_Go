package models

type Categories struct {
	ID    uint     `gorm:"primaryKey;autoIncrement"`
	Name  string   `gorm:"not null" json:"name"`
	Books []*Books `gorm:"many2many:book_category"`
}
