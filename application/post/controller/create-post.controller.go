package postcontroller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	postrepository "github.com/mauFade/revo/application/post/repository"
	postservice "github.com/mauFade/revo/application/post/service"
	"github.com/mauFade/revo/infra"
)

type createPostBody struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}

func CreatePostController(c *gin.Context) {
	var body createPostBody

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

	r := postrepository.NewPostRepository(infra.DB)

	s := postservice.NewCreatePostService(r)

	data, err := s.Execute(postservice.CreatePostInput{
		UserID: id,
		Title:  body.Title,
		Body:   body.Body,
	})

	if err != nil {
		c.JSON(
			http.StatusBadRequest, gin.H{
				"error": err.Error(),
			},
		)

		return
	}

	c.JSON(http.StatusCreated, data)
}
