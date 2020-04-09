package models

import "time"

type User struct {
	ID        uint      `gorm:"primary_key" json:"id"`
	Username  string    `json:"username";gorm:"type:varchar(255);NOT NULL`
	Password  string    `json:"password";gorm:"type:text"`
	Email     string    `json:"email";gorm:"type:varchar(255)";NOT NULL`
	UniqueKey string    `json:"unique_key";gorm:"type:varchar(100)"`
	Profile   Profile   `gorm:"foreignkey:user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type LoginValidate struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type RegisterValidate struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Email    string `json:"email" binding:"required"`
}
