package postservice

import (
	"errors"
	"time"

	"github.com/google/uuid"
	postmodel "github.com/mauFade/revo/application/post/model"
	postrepository "github.com/mauFade/revo/application/post/repository"
)

type LikePostInput struct {
	UserID string
	PostID string
}

type LikePostService struct {
	pr *postrepository.PostRepository
	lr *postrepository.LikeRepository
}

func NewLikePostService(
	p *postrepository.PostRepository,
	l *postrepository.LikeRepository,
) *LikePostService {
	return &LikePostService{
		pr: p,
		lr: l,
	}
}

func (s *LikePostService) Execute(data LikePostInput) error {
	post := s.pr.FindByID(data.PostID)

	if post == nil {
		return errors.New("Post not foun with this id")
	}

	post.Likes += 1

	s.pr.Update(post)

	like := postmodel.NewLike(
		uuid.NewString(),
		data.UserID,
		post.ID,
		time.Now(),
	)

	s.lr.Insert(like)

	return nil
}
