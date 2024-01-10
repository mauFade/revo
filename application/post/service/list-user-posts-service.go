package postservice

import (
	"time"

	postrepository "github.com/mauFade/revo/application/post/repository"
)

type ListUserPostsInput struct {
	UserId string
}

type ListUserPostsOutput struct {
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

type ListUserPostsService struct {
	r *postrepository.PostRepository
}

func NewListUserPostsService(repo *postrepository.PostRepository) *ListUserPostsService {
	return &ListUserPostsService{
		r: repo,
	}
}

func (s *ListUserPostsService) Execute(data ListUserPostsInput) []*ListUserPostsOutput {
	var output []*ListUserPostsOutput

	posts := s.r.FindUserPosts(data.UserId)

	for _, post := range posts {
		output = append(output, &ListUserPostsOutput{
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
