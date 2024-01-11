package usercontroller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	userrepository "github.com/mauFade/revo/application/user/repository"
	userservice "github.com/mauFade/revo/application/user/service"
	"github.com/mauFade/revo/infra"
)

func ListUserFollowersController(c *gin.Context) {
	id, err := infra.GetIdToken(c.GetHeader("Authorization"))

	if err != nil {
		c.JSON(
			http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			},
		)

		return
	}

	ur := userrepository.NewUserRepository(infra.DB)
	fr := userrepository.NewFollowerRepository(infra.DB)

	s := userservice.NewListUserFollowersService(ur, fr)

	data := s.Execute(userservice.ListFollowersInput{
		UserId: id,
	})

	c.JSON(http.StatusOK, data)
}
