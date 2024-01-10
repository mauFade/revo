package postcontroller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	postrepository "github.com/mauFade/revo/application/post/repository"
	postservice "github.com/mauFade/revo/application/post/service"
	userrepository "github.com/mauFade/revo/application/user/repository"
	"github.com/mauFade/revo/infra"
)

func ListUserPostController(c *gin.Context) {
	r := postrepository.NewPostRepository(infra.DB)
	ur := userrepository.NewUserRepository(infra.DB)

	s := postservice.NewListUserPostsService(r, ur)

	id, err := infra.GetIdToken(c.GetHeader("Authorization"))

	if err != nil {
		c.JSON(
			http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			},
		)

		return
	}

	data := s.Execute(postservice.ListUserPostsInput{
		UserId: id,
	})

	c.JSON(http.StatusOK, data)
}
