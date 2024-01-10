package userservice

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	userrepository "github.com/mauFade/revo/application/user/repository"
	"golang.org/x/crypto/bcrypt"
)

type AuthenticateUserService struct {
	r *userrepository.UserRepository
}

type AuthenticateInput struct {
	Email    string
	Password string
}

type AuthenticateOutput struct {
	Id       uuid.UUID `json:"id"`
	Name     string    `json:"name"`
	Email    string    `json:"email"`
	Username string    `json:"username"`
	Token    string    `json:"token"`
}

func NewAuthenticateUserService(repository *userrepository.UserRepository) *AuthenticateUserService {
	return &AuthenticateUserService{
		r: repository,
	}
}

func (s *AuthenticateUserService) Execute(data AuthenticateInput) (*AuthenticateOutput, error) {
	err := s.validateInput(data)

	if err != nil {
		return nil, err
	}

	user := s.r.FindByEmail(data.Email)

	if user == nil {
		return nil, errors.New("User not found with this email.")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(data.Password))

	if err != nil {
		return nil, errors.New("Invalid password")
	}

	token, err := s.generateToken(user.ID)

	if err != nil {
		return nil, err
	}

	return &AuthenticateOutput{
		Id:       user.ID,
		Name:     user.Name,
		Email:    user.Email,
		Username: user.Username,
		Token:    token,
	}, nil
}

func (s *AuthenticateUserService) validateInput(data AuthenticateInput) error {
	if data.Email == "" {
		return errors.New("Email is required.")
	}

	if data.Password == "" {
		return errors.New("Password is required.")
	}

	return nil
}

func (s *AuthenticateUserService) generateToken(userId uuid.UUID) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["userID"] = userId
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	tokenString, err := token.SignedString([]byte("AUTH_KEY"))

	if err != nil {
		return "", err
	}

	return tokenString, nil
}
