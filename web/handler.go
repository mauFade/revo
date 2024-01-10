package web

import (
	"github.com/gin-gonic/gin"
	postcontroller "github.com/mauFade/revo/application/post/controller"
	usercontroller "github.com/mauFade/revo/application/user/controller"
)

func NewHttpHandler() *gin.Engine {
	router := gin.Default()
	v1 := router.Group("/v1")

	users := v1.Group("/users")
	posts := v1.Group("/posts")

	v1.POST("/login", usercontroller.AuthenticateUserController)

	{
		users.POST("/", usercontroller.CreateUserController)
	}

	{
		posts.POST("/", postcontroller.CreatePostController)
	}

	return router
}
