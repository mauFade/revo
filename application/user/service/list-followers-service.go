package userservice

import (
	"time"

	userrepository "github.com/mauFade/revo/application/user/repository"
)

type ListUserFollowersService struct {
	ur *userrepository.UserRepository
	fr *userrepository.FollowerRepository
}

type ListFollowersInput struct {
	UserId string
}

type listFollowersOutput struct {
	FollowedUserId  string    `json:"followed_user_id"`
	FollowingUserId string    `json:"follower_user_id"`
	CreatedAt       time.Time `json:"created_at"`
}

func NewListUserFollowersService(
	ur *userrepository.UserRepository,
	fr *userrepository.FollowerRepository,
) *ListUserFollowersService {
	return &ListUserFollowersService{
		ur: ur,
		fr: fr,
	}
}

func (r *ListUserFollowersService) Execute(data ListFollowersInput) []*listFollowersOutput {
	followers := r.fr.GetUserFollowers(data.UserId)

	var output []*listFollowersOutput

	if len(followers) > 0 {
		for _, follower := range followers {
			output = append(output, &listFollowersOutput{
				FollowedUserId:  follower.FollowedUserId,
				FollowingUserId: follower.FollowingUserId,
				CreatedAt:       follower.CreatedAt,
			})
		}

		return output
	}

	return []*listFollowersOutput{}

}
