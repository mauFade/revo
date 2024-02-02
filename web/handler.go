package web

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	postcontroller "github.com/mauFade/revo/application/post/controller"
	usercontroller "github.com/mauFade/revo/application/user/controller"
)

func NewHttpHandler() *gin.Engine {
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	router.POST("/v1/login", usercontroller.AuthenticateUserController)

	router.POST("/v1/users", usercontroller.CreateUserController)

	router.Use(EnsureAuthenticatedMiddleware())

	router.GET("/v1/users/followers", usercontroller.ListUserFollowersController)
	router.POST("/v1/users/followers", usercontroller.FollowUnfollowController)

	router.POST("/v1/posts", postcontroller.CreatePostController)
	router.GET("/v1/posts/profile", postcontroller.ListUserPostController)
	router.GET("/v1/posts", postcontroller.ListFollowingPostsController)
	router.POST("/v1/posts/like", postcontroller.LikePostController)

	return router
}
