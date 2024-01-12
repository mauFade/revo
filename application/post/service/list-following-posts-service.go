package postservice

import (
	"sort"
	"time"

	postmodel "github.com/mauFade/revo/application/post/model"
	postrepository "github.com/mauFade/revo/application/post/repository"
	userrepository "github.com/mauFade/revo/application/user/repository"
)

type ListFollowingPostsInput struct {
	UserId string
}

type listFollowingPostsOutput struct {
	ID        string     `json:"id"`
	UserID    string     `json:"user_id"`
	Title     string     `json:"title"`
	Body      string     `json:"body"`
	Likes     int64      `json:"likes"`
	Shares    int64      `json:"shares"`
	Comments  int64      `json:"comments"`
	Deleted   bool       `json:"deleted"`
	DeletedAt *time.Time `json:"deleted_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	CreatedAt time.Time  `json:"created_at"`
}

type ListFollowingPostsService struct {
	pr *postrepository.PostRepository
	ur *userrepository.UserRepository
	fr *userrepository.FollowerRepository
}

func NewListFollowingPostsService(
	pr *postrepository.PostRepository,
	ur *userrepository.UserRepository,
	fr *userrepository.FollowerRepository,
) *ListFollowingPostsService {
	return &ListFollowingPostsService{
		pr: pr,
		ur: ur,
		fr: fr,
	}
}

func (s *ListFollowingPostsService) Execute(data ListFollowingPostsInput) []*listFollowingPostsOutput {
	userFollowing := s.fr.GetUserFollowing(data.UserId)
	var gollowingIDs []string
	var output []*listFollowingPostsOutput

	for _, follower := range userFollowing {
		gollowingIDs = append(gollowingIDs, follower.FollowedUserId)
	}

	posts := s.getCombinedPostsSorted(gollowingIDs, data.UserId)

	for _, post := range posts {
		output = append(output, &listFollowingPostsOutput{
			ID:        post.ID,
			UserID:    post.UserID,
			Title:     post.Title,
			Body:      post.Body,
			Likes:     post.Likes,
			Shares:    post.Shares,
			Comments:  post.Comments,
			Deleted:   post.Deleted,
			DeletedAt: post.DeletedAt,
			UpdatedAt: post.UpdatedAt,
			CreatedAt: post.CreatedAt,
		})
	}

	return output
}

func (s *ListFollowingPostsService) getCombinedPostsSorted(followerIDs []string, userID string) []*postmodel.Post {
	followingPosts := s.pr.FindByUserIDMacro(followerIDs)
	userPosts := s.pr.FindUserPosts(userID)

	combinedPosts := append(followingPosts, userPosts...)

	sort.Slice(combinedPosts, func(i, j int) bool {
		return combinedPosts[i].CreatedAt.After(combinedPosts[j].CreatedAt)
	})

	return combinedPosts
}
