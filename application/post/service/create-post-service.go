package postservice

import (
	"errors"
	"time"

	"github.com/google/uuid"
	postmodel "github.com/mauFade/revo/application/post/model"
	postrepository "github.com/mauFade/revo/application/post/repository"
)

type CreatePostInput struct {
	UserID string
	Title  string
	Body   string
}

type CreatePostOutput struct {
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

type CreatePostService struct {
	r *postrepository.PostRepository
}

func NewCreatePostService(repo *postrepository.PostRepository) *CreatePostService {
	return &CreatePostService{
		r: repo,
	}
}

func (s *CreatePostService) Execute(data CreatePostInput) (*CreatePostOutput, error) {
	err := s.validateInput(data)

	if err != nil {
		return nil, err
	}

	post := postmodel.NewPost(
		uuid.NewString(),
		data.UserID,
		data.Title,
		data.Body,
		0,
		0,
		0,
		false,
		nil,
		time.Now(),
		time.Now(),
	)

	s.r.Create(post)

	return &CreatePostOutput{
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
	}, nil
}

func (s *CreatePostService) validateInput(data CreatePostInput) error {
	if data.Body == "" {
		return errors.New("Body is required.")
	}

	if data.Title == "" {
		return errors.New("Title is required.")
	}

	if data.UserID == "" {
		return errors.New("User ID is required.")
	}

	return nil
}
