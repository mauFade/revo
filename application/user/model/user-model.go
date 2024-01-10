package usermodel

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID  `gorm:"type:uuid"`
	Name      string     `gorm:"type:varchar"`
	Email     string     `gorm:"type:varchar"`
	Phone     string     `gorm:"type:varchar"`
	Password  string     `gorm:"type:varchar"`
	Username  string     `gorm:"type:varchar"`
	Bio       string     `gorm:"type:varchar"`
	Avatar    *string    `gorm:"type:varchar"`
	City      string     `gorm:"type:varchar"`
	Country   string     `gorm:"type:varchar"`
	Deleted   bool       `gorm:"type:bool"`
	DeletedAt *time.Time `gorm:"type:timestamp"`
	UpdatedAt time.Time  `gorm:"type:timestamp"`
	CreatedAt time.Time  `gorm:"type:timestamp"`
}

func NewUser(
	id uuid.UUID,
	name string,
	email string,
	phone string,
	password string,
	username string,
	bio string,
	avatar *string,
	city string,
	country string,
	deleted bool,
	deleted_at *time.Time,
	updated_at time.Time,
	created_at time.Time,
) *User {
	return &User{
		ID:        id,
		Name:      name,
		Email:     email,
		Phone:     phone,
		Password:  password,
		Username:  username,
		Bio:       bio,
		Avatar:    avatar,
		City:      city,
		Country:   country,
		Deleted:   deleted,
		DeletedAt: deleted_at,
		UpdatedAt: updated_at,
		CreatedAt: created_at,
	}
}
