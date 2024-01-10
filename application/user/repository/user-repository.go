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

	result := r.db.Where("email = ?", email).First(&user)

	if result.RowsAffected == 0 {
		return nil
	}

	return &user
}

func (r *UserRepository) FindByPhone(phone string) *usermodel.User {
	var user usermodel.User

	result := r.db.Where("phone = ?", phone).First(&user)

	if result.RowsAffected == 0 {
		return nil
	}

	return &user
}

func (r *UserRepository) FindByUsername(username string) *usermodel.User {
	var user usermodel.User

	result := r.db.Where("username = ?", username).First(&user)

	if result.RowsAffected == 0 {
		return nil
	}

	return &user
}

func (r *UserRepository) FindByIdMacro(ids []string) []*usermodel.User {
	var users []*usermodel.User

	result := r.db.Where("id IN ?", ids).Find(&users)

	if result.RowsAffected == 0 {
		return []*usermodel.User{}
	}

	return users
}
