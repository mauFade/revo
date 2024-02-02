package postdto

import "time"

type FindByUserIDMacroDTO struct {
	ID        string
	UserID    string
	Title     string
	Body      string
	Likes     int64
	Shares    int64
	Comments  int64
	Deleted   bool
	DeletedAt *time.Time
	UpdatedAt time.Time
	CreatedAt time.Time
	Name      string
	Email     string
	Username  string
	Avatar    *string
}
