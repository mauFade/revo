package postrepository

import (
	postdto "github.com/mauFade/revo/application/post/dto"
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

func (r *PostRepository) Update(post *postmodel.Post) {
	r.db.Save(post)
}

func (r *PostRepository) FindByID(id string) *postmodel.Post {
	var post postmodel.Post

	result := r.db.Where("id = ?", id).First(&post)

	if result.RowsAffected == 0 {
		return nil
	}

	return &post
}

func (r *PostRepository) FindByUserIDMacro(userIds []string) []*postdto.FindByUserIDMacroDTO {
	var posts []*postdto.FindByUserIDMacroDTO

	result := r.db.
		Table("posts").
		Select("posts.*, users.id AS user_id, users.name, users.email, users.username, users.avatar").
		Joins("JOIN users ON posts.user_id = users.id").
		Where("posts.user_id IN (?)", userIds).
		Find(&posts)

	if result.RowsAffected == 0 {
		return []*postdto.FindByUserIDMacroDTO{}
	}

	return posts
}

func (r *PostRepository) FindUserPosts(userId string) []*postdto.FindByUserIDMacroDTO {
	var posts []*postdto.FindByUserIDMacroDTO

	result := r.db.
		Table("posts").
		Select("posts.*, users.id AS user_id, users.name, users.email, users.username, users.avatar").
		Joins("JOIN users ON posts.user_id = users.id").
		Where("posts.user_id = ?", userId).
		Find(&posts)

	if result.RowsAffected == 0 {
		return []*postdto.FindByUserIDMacroDTO{}
	}

	return posts
}
