package userservice

import (
	"errors"
	"regexp"
	"time"

	"github.com/google/uuid"
	usermodel "github.com/mauFade/revo/application/user/model"
	userrepository "github.com/mauFade/revo/application/user/repository"
)

type FollowUnfollowInput struct {
	InteractionType  string
	UserId           string
	FollowedUserId   string
	UnfollowedUserId string
}

type unfollowInput struct {
	UserId           string
	UnfollowedUserId string
}

type followInput struct {
	UserId         string
	FollowedUserId string
}

type FollowUnfollowService struct {
	ur *userrepository.UserRepository
	fr *userrepository.FollowerRepository
}

func NewFollowUnfollowService(
	ur *userrepository.UserRepository,
	fr *userrepository.FollowerRepository,
) *FollowUnfollowService {
	return &FollowUnfollowService{
		ur: ur,
		fr: fr,
	}
}

func (s *FollowUnfollowService) Execute(data FollowUnfollowInput) error {
	err := s.validateInput(data)

	if err != nil {
		return err
	}

	if data.InteractionType == "follow" {
		err := s.follow(followInput{
			UserId:         data.UserId,
			FollowedUserId: data.FollowedUserId,
		})

		if err != nil {
			return err
		}

		return nil
	}

	if data.InteractionType == "unfollow" {
		err := s.unfollow(unfollowInput{
			UserId:           data.UserId,
			UnfollowedUserId: data.UnfollowedUserId,
		})

		if err != nil {
			return err
		}

		return nil
	}

	return nil
}

func (s *FollowUnfollowService) validateInput(data FollowUnfollowInput) error {
	if !regexp.MustCompile(`^(follow|unfollow)$`).MatchString(data.InteractionType) {
		return errors.New("Interaction type must match 'follow' or 'unfollow'")
	}

	if data.InteractionType == "follow" && data.FollowedUserId == "" {
		return errors.New("In order to follow, must provid a followed user id")
	}

	if data.InteractionType == "unfollow" && data.UnfollowedUserId == "" {
		return errors.New("In order to unfollow, must provid a followed user id")
	}

	return nil
}

func (s *FollowUnfollowService) unfollow(data unfollowInput) error {
	unfollowedUser := s.ur.FindById(data.UnfollowedUserId)

	if unfollowedUser == nil {
		return errors.New("Unfollowed user not found with this id.")
	}

	follow := s.fr.FindByFollower(unfollowedUser.ID, data.UserId)

	if follow != nil {
		if unfollowedUser.Followers > 0 {
			unfollowedUser.Followers -= 1
		} else {
			unfollowedUser.Followers = 0
		}

		s.ur.Update(unfollowedUser)

		s.fr.Delete(follow.ID)
	}

	return nil
}

func (s *FollowUnfollowService) follow(data followInput) error {
	followedUser := s.ur.FindById(data.FollowedUserId)

	if followedUser == nil {
		return errors.New("Followed user not found with this id.")
	}

	followedUser.Followers += 1

	newFollow := usermodel.NewFollowerFollowed(
		uuid.NewString(),
		followedUser.ID,
		data.UserId,
		time.Now(),
	)

	s.fr.Create(newFollow)

	s.ur.Update(followedUser)

	return nil
}
