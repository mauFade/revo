package usercontroller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	userrepository "github.com/mauFade/revo/application/user/repository"
	userservice "github.com/mauFade/revo/application/user/service"
	"github.com/mauFade/revo/infra"
)

type authenticateBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func AuthenticateUserController(c *gin.Context) {
	var body authenticateBody

	c.Bind(&body)

	r := userrepository.NewUserRepository(infra.DB)

	s := userservice.NewAuthenticateUserService(r)

	data, err := s.Execute(userservice.AuthenticateInput{
		Email:    body.Email,
		Password: body.Password,
	})

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, data)
}
