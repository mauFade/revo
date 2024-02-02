package postmodel

import "time"

type Like struct {
	ID         string    `gorm:"type:uuid"`
	UserID     string    `gorm:"type:uuid"`
	PostID     string    `gorm:"type:uuid"`
	FollowedAt time.Time `gorm:"type:timestamp"`
}

func NewLike(id, userId, postId string, followedAt time.Time) *Like {
	return &Like{
		ID:         id,
		UserID:     userId,
		PostID:     postId,
		FollowedAt: followedAt,
	}
}
