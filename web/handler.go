package web

import (
	"github.com/gin-gonic/gin"
	usercontroller "github.com/mauFade/revo/application/user/controller"
)

func NewHttpHandler() *gin.Engine {
	router := gin.Default()
	v1 := router.Group("/v1")

	users := v1.Group("/users")

	v1.POST("/login", usercontroller.AuthenticateUserController)

	users.POST("/", usercontroller.CreateUserController)

	return router
}
