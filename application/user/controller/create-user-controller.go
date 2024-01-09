package usercontroller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	userrepository "github.com/mauFade/revo/application/user/repository"
	userservice "github.com/mauFade/revo/application/user/service"
	"github.com/mauFade/revo/infra"
)

type createUserBody struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Password string `json:"password"`
	Username string `json:"username"`
	Bio      string `json:"bio"`
	City     string `json:"city"`
	Country  string `json:"country"`
}

func CreateUserController(c *gin.Context) {
	var body createUserBody

	c.Bind(&body)

	userRepository := userrepository.NewUserRepository(infra.DB)

	createUserService := userservice.NewCreateUserService(userRepository)

	user, err := createUserService.Execute(userservice.CreateUserInput{
		Name:     body.Name,
		Email:    body.Email,
		Phone:    body.Phone,
		Password: body.Password,
		Username: body.Username,
		Bio:      body.Bio,
		City:     body.City,
		Country:  body.Country,
	})

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"data": user,
	})
}
