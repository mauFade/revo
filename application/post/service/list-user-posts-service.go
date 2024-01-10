package postservice

import (
	"time"

	postrepository "github.com/mauFade/revo/application/post/repository"
	userrepository "github.com/mauFade/revo/application/user/repository"
)

type ListUserPostsInput struct {
	UserId string
}

type UserStruct struct {
	ID       string  `json:"id"`
	Name     string  `json:"name"`
	Username string  `json:"username"`
	Avatar   *string `json:"avatar"`
}

type ListUserPostsOutput struct {
	ID        string     `json:"id"`
	UserID    string     `json:"user_id"`
	Title     string     `json:"title"`
	Body      string     `json:"body"`
	Likes     int64      `json:"likes"`
	Shares    int64      `json:"shares"`
	Comments  int64      `json:"comments"`
	User      UserStruct `json:"user"`
	Deleted   bool       `json:"deleted"`
	DeletedAt *time.Time `json:"deleted_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	CreatedAt time.Time  `json:"created_at"`
}

type ListUserPostsService struct {
	pr *postrepository.PostRepository
	ur *userrepository.UserRepository
}

func NewListUserPostsService(repo *postrepository.PostRepository, ur *userrepository.UserRepository) *ListUserPostsService {
	return &ListUserPostsService{
		pr: repo,
		ur: ur,
	}
}

func (s *ListUserPostsService) Execute(data ListUserPostsInput) []*ListUserPostsOutput {
	var output []*ListUserPostsOutput
	var userIds []string

	posts := s.pr.FindUserPosts(data.UserId)

	for _, post := range posts {
		userIds = append(userIds, post.UserID)
	}

	users := s.ur.FindByIdMacro(userIds)

	for _, post := range posts {
		for _, user := range users {
			if post.UserID == user.ID {
				output = append(output, &ListUserPostsOutput{
					ID:       post.ID,
					UserID:   post.UserID,
					Title:    post.Title,
					Body:     post.Body,
					Likes:    post.Likes,
					Shares:   post.Shares,
					Comments: post.Comments,
					Deleted:  post.Deleted,
					User: UserStruct{
						ID:       user.ID,
						Name:     user.Name,
						Username: "@" + user.Username,
						Avatar:   user.Avatar,
					},
					DeletedAt: post.DeletedAt,
					UpdatedAt: post.UpdatedAt,
					CreatedAt: post.CreatedAt,
				})
			}
		}

	}

	return output
}
