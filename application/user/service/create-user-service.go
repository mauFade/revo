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

type CreateUserOutput struct {
	Id        string     `json:"id"`
	Name      string     `json:"name"`
	Email     string     `json:"email"`
	Phone     string     `json:"phone"`
	Username  string     `json:"username"`
	Bio       string     `json:"bio"`
	Avatar    *string    `json:"avatar"`
	City      string     `json:"city"`
	Country   string     `json:"country"`
	Deleted   bool       `json:"deleted"`
	Followers int64      `json:"followers"`
	DeletedAt *time.Time `json:"deleted_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	CreatedAt time.Time  `json:"created_at"`
}

type CreateUserService struct {
	repository *userrepository.UserRepository
}

func NewCreateUserService(r *userrepository.UserRepository) *CreateUserService {
	return &CreateUserService{
		repository: r,
	}
}

func (s *CreateUserService) Execute(data CreateUserInput) (*CreateUserOutput, error) {
	emailExists := s.repository.FindByEmail(data.Email)

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
		nil,
		data.City,
		data.Country,
		false,
		0,
		nil,
		time.Now(),
		time.Now(),
	)

	s.repository.Create(user)

	return &CreateUserOutput{
		Id:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Phone:     user.Phone,
		Username:  user.Username,
		Bio:       user.Bio,
		Avatar:    user.Avatar,
		City:      user.City,
		Country:   user.Country,
		Deleted:   user.Deleted,
		DeletedAt: user.DeletedAt,
		Followers: user.Followers,
		UpdatedAt: user.UpdatedAt,
		CreatedAt: user.CreatedAt,
	}, nil
}
