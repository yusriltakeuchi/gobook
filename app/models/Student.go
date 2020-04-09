package models

type Student struct {
	ID        uint      `gorm:"primary_key" json:"id"`
}