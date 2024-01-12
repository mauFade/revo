package postcontroller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	postrepository "github.com/mauFade/revo/application/post/repository"
	postservice "github.com/mauFade/revo/application/post/service"
	userrepository "github.com/mauFade/revo/application/user/repository"
	"github.com/mauFade/revo/infra"
)

func ListFollowingPostsController(c *gin.Context) {
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
	pr := postrepository.NewPostRepository(infra.DB)

	s := postservice.NewListFollowingPostsService(
		pr,
		ur,
		fr,
	)

	data := s.Execute(postservice.ListFollowingPostsInput{
		UserId: id,
	})

	c.JSON(http.StatusOK, data)
}
