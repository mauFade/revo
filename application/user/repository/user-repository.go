package userrepository

import (
	usermodel "github.com/mauFade/revo/application/user/model"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) Create(user *usermodel.User) {
	r.db.Create(user)
}

func (r *UserRepository) FindByEmail(email string) *usermodel.User {
	var user usermodel.User

	// sql := "SELECT * FROM users WHERE email = ?"

	// // result := r.db.Raw(sql, email).Scan(&user)

	result := r.db.Where("email = ?", email).First(&user)

	if result.RowsAffected == 0 {
		return nil
	}

	return &user
}
