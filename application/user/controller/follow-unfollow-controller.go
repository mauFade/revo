package usercontroller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	userrepository "github.com/mauFade/revo/application/user/repository"
	userservice "github.com/mauFade/revo/application/user/service"
	"github.com/mauFade/revo/infra"
)

type followUnfollowBody struct {
	InteractionType  string `json:"interaction_type"`
	FollowedUserId   string `json:"followed_user_id"`
	UnfollowedUserId string `json:"unfollowed_user_id"`
}

func FollowUnfollowController(c *gin.Context) {
	var body followUnfollowBody

	c.Bind(&body)

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

	s := userservice.NewFollowUnfollowService(ur, fr)

	err = s.Execute(userservice.FollowUnfollowInput{
		UserId:           id,
		InteractionType:  body.InteractionType,
		FollowedUserId:   body.FollowedUserId,
		UnfollowedUserId: body.UnfollowedUserId,
	})

	if err != nil {
		c.JSON(
			http.StatusBadRequest, gin.H{
				"error": err.Error(),
			},
		)

		return
	}

	output := make(map[string]bool)

	output["success"] = true

	c.JSON(http.StatusOK, output)
}
