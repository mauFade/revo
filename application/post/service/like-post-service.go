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
	like := s.FindLike(data.PostID, data.UserID)

	post := s.pr.FindByID(data.PostID)

	if post == nil {
		return errors.New("Post not found with this id")
	}

	if like == nil {
		post.Likes += 1

		s.pr.Update(post)

		like := postmodel.NewLike(
			uuid.NewString(),
			data.UserID,
			post.ID,
			time.Now(),
		)

		s.lr.Insert(like)
	} else {
		if post.Likes > 0 {
			post.Likes -= 1
		} else {
			post.Likes = 0
		}

		s.pr.Update(post)

		s.lr.Delete(like.ID)
	}

	return nil
}

func (s *LikePostService) FindLike(postId, userId string) *postmodel.Like {
	like := s.lr.FindByUserIDAndPostID(postId, userId)

	return like
}
