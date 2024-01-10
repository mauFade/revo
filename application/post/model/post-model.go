package postmodel

import (
	"time"
)

type Post struct {
	ID        string     `gorm:"type:uuid"`
	UserID    string     `gorm:"type:uuid"`
	Title     string     `gorm:"type:varchar"`
	Body      string     `gorm:"type:text"`
	Likes     int64      `gorm:"type:int8"`
	Shares    int64      `gorm:"type:int8"`
	Comments  int64      `gorm:"type:int8"`
	Deleted   bool       `gorm:"type:bool"`
	DeletedAt *time.Time `gorm:"type:timestamp"`
	UpdatedAt time.Time  `gorm:"type:timestamp"`
	CreatedAt time.Time  `gorm:"type:timestamp"`
}

func NewPost(
	id,
	userId,
	title,
	body string,
	likes,
	shares,
	comments int64,
	deleted bool,
	deleted_at *time.Time,
	updated_at,
	created_at time.Time,
) *Post {
	return &Post{
		ID:        id,
		UserID:    userId,
		Title:     title,
		Body:      body,
		Likes:     likes,
		Shares:    shares,
		Comments:  comments,
		Deleted:   deleted,
		DeletedAt: deleted_at,
		UpdatedAt: updated_at,
		CreatedAt: created_at,
	}
}
