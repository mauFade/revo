package postcontroller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	postrepository "github.com/mauFade/revo/application/post/repository"
	postservice "github.com/mauFade/revo/application/post/service"
	"github.com/mauFade/revo/infra"
)

func ListUserPostController(c *gin.Context) {
	r := postrepository.NewPostRepository(infra.DB)

	s := postservice.NewListUserPostsService(r)

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
