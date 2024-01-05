package user

import "time"

type User struct {
	Id        string
	Name      string
	Email     string
	Phone     string
	Password  string
	Username  string
	Bio       string
	Avatar    string
	City      string
	Country   string
	Deleted   bool
	DeletedAt *time.Time
	UpdatedAt time.Time
	CreatedAt time.Time
}

func NewUser(
	id string,
	name string,
	email string,
	phone string,
	password string,
	username string,
	bio string,
	avatar string,
	city string,
	country string,
	deleted bool,
	deleted_at *time.Time,
	updated_at time.Time,
	created_at time.Time,
) *User {
	return &User{
		Id:        id,
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
