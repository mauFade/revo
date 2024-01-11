package usermodel

import (
	"time"
)

type FollowerFollowed struct {
	FollowedUserId  string    `gorm:"type:uuid"`
	FollowingUserId string    `gorm:"type:uuid"`
	CreatedAt       time.Time `gorm:"type:timestamp"`
}

type tabler interface {
	TableName() string
}

func (FollowerFollowed) tableName() string {
	return "followers"
}

func NewFollowerFollowed(followedUserId, followingUserId string, createdAt time.Time) *FollowerFollowed {
	return &FollowerFollowed{
		FollowedUserId:  followedUserId,
		FollowingUserId: followingUserId,
		CreatedAt:       createdAt,
	}
}
