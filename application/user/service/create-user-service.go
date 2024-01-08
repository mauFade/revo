package userservice

import (
	"errors"
	"time"

	"github.com/google/uuid"
	usermodel "github.com/mauFade/revo/application/user/model"
	userrepository "github.com/mauFade/revo/application/user/repository"
	"golang.org/x/crypto/bcrypt"
)

type CreateUserInput struct {
	Name     string
	Email    string
	Phone    string
	Password string
	Username string
	Bio      string
	City     string
	Country  string
}

type CreateUserService struct {
	Repository *userrepository.UserRepository
}

func (s *CreateUserService) Execute(data *CreateUserInput) (*usermodel.User, error) {
	emailExists := s.Repository.FindByEmail(data.Email)

	if emailExists != nil {
		return nil, errors.New("This email is already in use.")
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(data.Password), 6)

	if err != nil {
		return nil, errors.New("Error on generating hash: " + err.Error())
	}

	user := usermodel.NewUser(
		uuid.NewString(),
		data.Name,
		data.Email,
		data.Phone,
		string(hash),
		data.Username,
		data.Bio,
		"",
		data.City,
		data.Country,
		false,
		nil,
		time.Now(),
		time.Now(),
	)

	s.Repository.Create(user)

}
