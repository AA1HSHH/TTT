package dal

import (
	"gorm.io/gorm"
)

type VideoFavorite struct {
	Id       int64          `gorm:"column:id"`
	VideoId  int64          `gorm:"column:video_id"`
	LikerId  int64          `gorm:"column:liker_id"`
	IsDelete gorm.DeletedAt `gorm:"column:is_delete"`
}

func (VideoFavorite) TableName() string {
	return "t_video_favorite"
}
func ISFavorite(videoId, userId int64) bool {
	vf := make([]VideoFavorite, 0)
	rst := db.Model(&VideoFavorite{}).
		Where("video_id = ?", videoId).
		Where("liker_id = ?", userId).
		Where("is_delete IS NULL").Find(&vf)
	if rst.Error != nil || len(vf) == 0 {
		return false
	}
	return true
}

func CreatFavoirte(vid, likeId int64) error {
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if err := tx.Error; err != nil {
		return err
	}
	vf := VideoFavorite{VideoId: vid, LikerId: likeId}
	if err := tx.Model(&VideoFavorite{}).Create(&vf).Error; err != nil {
		tx.Rollback()
		return err
	}
	if err := tx.Model(&DBVideo{}).
		Where("id = ?", vid).
		UpdateColumn("favorite_count", gorm.Expr("favorite_count + ?", 1)).Error; err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

func DeleteFavoirte(vid, likeId int64) error {
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if err := tx.Error; err != nil {
		return err
	}
	vf := VideoFavorite{}
	if err := tx.Model(&VideoFavorite{}).
		Where("video_id = ?", vid).
		Where("liker_id = ?", likeId).Delete(&vf).Error; err != nil {
		tx.Rollback()
		return err
	}
	if err := tx.Model(&DBVideo{}).
		Where("id = ?", vid).
		UpdateColumn("favorite_count", gorm.Expr("favorite_count - ?", 1)).Error; err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

func QueryFavoriteList(uid int64) ([]DBVideo, map[int64]User, error) {
	videos := make([]DBVideo, 0)
	// select video_id from t_video_favorite
	query := db.Select("video_id").
		Where("liker_id = ?", uid).
		Where("is_delete IS NULL").Table("t_video_favorite")
	// select * from t_video
	rst := db.Model(&DBVideo{}).Where("id in (?)", query).Find(&videos)
	if rst.Error != nil {
		return nil, nil, rst.Error
	}
	// select author_id from t_video
	type AuthorId int64
	var uids []AuthorId
	rst.Select("author_id").Scan(&uids)
	// select * from t_user
	users := make([]User, 0)
	if rst = db.Model(&User{}).Where("id in (?)", uids).Find(&users); rst.Error != nil {
		return nil, nil, rst.Error
	}

	authorMap := make(map[int64]User, len(users))
	for _, usr := range users {
		authorMap[usr.Id] = usr
	}
	return videos, authorMap, nil
}
