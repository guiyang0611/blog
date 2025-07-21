package models

import (
	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	Title    string
	Content  string
	UserID   uint
	Status   CommentStatus
	Comments []Comment `gorm:"foreignKey:PostID"`
}

func (*Post) TableName() string {
	return "posts"
}
