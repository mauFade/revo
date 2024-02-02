package postcontroller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	postrepository "github.com/mauFade/revo/application/post/repository"
	postservice "github.com/mauFade/revo/application/post/service"
	"github.com/mauFade/revo/infra"
)

type likePostBody struct {
	PostID string `json:"post_id"`
}

func LikePostController(c *gin.Context) {
	var body likePostBody

	id, err := infra.GetIdToken(c.GetHeader("Authorization"))

	if err != nil {
		c.JSON(
			http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			},
		)

		return
	}

	c.Bind(&body)

	pr := postrepository.NewPostRepository(infra.DB)
	lr := postrepository.NewLikeRepository(infra.DB)

	s := postservice.NewLikePostService(
		pr,
		lr,
	)

	err = s.Execute(postservice.LikePostInput{
		UserID: id,
		PostID: body.PostID,
	})

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	output := make(map[string]bool)

	output["success"] = true

	c.JSON(http.StatusCreated, output)
}
