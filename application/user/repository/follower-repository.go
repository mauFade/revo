package userrepository

import (
	usermodel "github.com/mauFade/revo/application/user/model"
	"gorm.io/gorm"
)

type FollowerRepository struct {
	db *gorm.DB
}

func NewFollowerRepository(d *gorm.DB) *FollowerRepository {
	return &FollowerRepository{
		db: d,
	}
}

func (r *FollowerRepository) Create(model *usermodel.FollowerFollowed) {
	r.db.Create(model)
}

func (r *FollowerRepository) Delete(followId string) {
	r.db.Delete(&usermodel.FollowerFollowed{}, "id = ?", followId)
}

func (r *FollowerRepository) FindByFollower(followedId, followerId string) *usermodel.FollowerFollowed {
	var follow *usermodel.FollowerFollowed

	result := r.db.Where(map[string]interface{}{"followed_user_id": followedId, "following_user_id": followerId}).First(&follow)

	if result.RowsAffected == 0 {
		return nil
	}

	return follow
}

func (r *FollowerRepository) GetUserFollowers(userId string) []*usermodel.FollowerFollowed {
	var followers []*usermodel.FollowerFollowed

	result := r.db.Where("followed_user_id = ?", userId).Find(&followers)

	if result.RowsAffected == 0 {
		return []*usermodel.FollowerFollowed{}
	}

	return followers
}

func (r *FollowerRepository) GetUserFollowing(userId string) []*usermodel.FollowerFollowed {
	var followers []*usermodel.FollowerFollowed

	result := r.db.Where("following_user_id = ?", userId).Find(&followers)

	if result.RowsAffected == 0 {
		return []*usermodel.FollowerFollowed{}
	}

	return followers
}
