package models

type Profile struct {
	ID     uint   `gorm:"primary_key" json:"id"`
	Name   string `json:"name";gorm:"type:varchar(255)";NOT NULL`
	Phone  string `json:"phone";gorm:"type:varchar(100)"`
	City   string `json:"city";gorm:"type:varchar(255)"`
	UserID uint   `json:"user_id"`
}
