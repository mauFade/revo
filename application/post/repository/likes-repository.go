package postrepository

import (
	postmodel "github.com/mauFade/revo/application/post/model"
	"gorm.io/gorm"
)

type LikeRepository struct {
	DB *gorm.DB
}

func NewLikeRepository(db *gorm.DB) *LikeRepository {
	return &LikeRepository{
		DB: db,
	}
}

func (r *LikeRepository) Insert(entity *postmodel.Like) {
	r.DB.Create(entity)
}

func (r *LikeRepository) Update(entity *postmodel.Like) {
	r.DB.Save(entity)
}
