package postrepository

import (
	postmodel "github.com/mauFade/revo/application/post/model"
	"gorm.io/gorm"
)

type PostRepository struct {
	db *gorm.DB
}

func NewPostRepository(d *gorm.DB) *PostRepository {
	return &PostRepository{
		db: d,
	}
}

func (r *PostRepository) Create(post *postmodel.Post) {
	r.db.Create(post)
}

func (r *PostRepository) FindByID(id string) *postmodel.Post {
	var post postmodel.Post

	result := r.db.Where("id = ?", id).First(&post)

	if result.RowsAffected == 0 {
		return nil
	}

	return &post
}

func (r *PostRepository) FindUserPosts(userId string) []*postmodel.Post {
	var posts []*postmodel.Post

	result := r.db.Where("user_id = ?", userId).Find(&posts)

	if result.RowsAffected == 0 {
		return []*postmodel.Post{}
	}

	return posts
}
