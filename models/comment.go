package models

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	Content string
	PostID  uint
	UserID  uint
}

func (*Comment) TableName() string {
	return "comments"
}

type CommentStatus int

const (
	CommentDisabled CommentStatus = iota
	CommentEnabled
)

func (s CommentStatus) Desc() string {
	return [...]string{"无评论", "有评论"}[s]
}
