package postrepository

import (
	postmodel "github.com/mauFade/revo/application/post/model"
	"gorm.io/gorm"
)

type LikeRepository struct {
	db *gorm.DB
}

func NewLikeRepository(db *gorm.DB) *LikeRepository {
	return &LikeRepository{
		db: db,
	}
}

func (r *LikeRepository) Insert(entity *postmodel.Like) {
	r.db.Create(entity)
}

func (r *LikeRepository) Update(entity *postmodel.Like) {
	r.db.Save(entity)
}

func (r *LikeRepository) FindByUserIDAndPostID(postId, userId string) *postmodel.Like {
	var like *postmodel.Like

	result := r.db.
		Table("likes").
		Select("likes.*").
		Where("user_id = ? AND post_id = ?", userId, postId).
		Find(&like)

	if result.RowsAffected == 0 {
		return nil
	}

	return like
}

func (r *LikeRepository) FindByPostIDMacro(postIDs []string) []*postmodel.Like {
	var likes []*postmodel.Like

	result := r.db.
		Table("likes").
		Select("likes.*").
		Where("post_id IN (?)", postIDs).
		Find(&likes)

	if result.RowsAffected == 0 {
		return []*postmodel.Like{}
	}

	return likes
}
