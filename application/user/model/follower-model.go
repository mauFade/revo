package usermodel

import (
	"time"
)

type FollowerFollowed struct {
	ID              string    `gorm:"type:uuid"`
	FollowedUserId  string    `gorm:"type:uuid"`
	FollowingUserId string    `gorm:"type:uuid"`
	CreatedAt       time.Time `gorm:"type:timestamp"`
}

type Tabler interface {
	TableName() string
}

func (FollowerFollowed) TableName() string {
	return "followers"
}

func NewFollowerFollowed(id, followedUserId, followingUserId string, createdAt time.Time) *FollowerFollowed {
	return &FollowerFollowed{
		ID:              id,
		FollowedUserId:  followedUserId,
		FollowingUserId: followingUserId,
		CreatedAt:       createdAt,
	}
}
