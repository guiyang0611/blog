package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"unique" json:"username"`
	Password string `json:"password"`
	Phone    string `gorm:"unique" json:"phone"`
	Address  string `json:"address"`
	Status   int    `json:"status"`
}

func (*User) TableName() string {
	return "users"
}
