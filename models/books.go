package models

import "time"

type Books struct {
	ID          uint      `gorm:"primary_key" json:"id"`
	Title       string    `json:"title";gorm:"type:varchar(255);NOT NULL`
	Description string    `json:"description";gorm:"type:text"`
	Genre       string    `json:"genre";gorm:"type:varchar(100)";NOT NULL`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type BooksValidate struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	Genre       string `json:"genre" binding:"required"`
}
