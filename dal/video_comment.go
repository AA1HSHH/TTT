package dal

import (
	"gorm.io/gorm"
	"time"
)

type VideoComment struct {
	Id         int64          `gorm:"column:id"`
	VideoId    int64          `gorm:"column:video_id"`
	WriterId   int64          `gorm:"column:writer_id"`
	Content    string         `gorm:"column:content"`
	CreateTime time.Time      `gorm:"column:create_date"`
	IsDelete   gorm.DeletedAt `gorm:"column:is_delete"`
}

func (VideoComment) TableName() string {
	return "t_video_comment"
}

func CreateVideoComment(vid int64, uid int64, commenttext string) (*VideoComment, error) {
	//
	//if err := db.Model(&VideoComment{}).Create(&vc).Error; err != nil {
	//	return nil, err
	//}
	//return &vc, nil
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if err := tx.Error; err != nil {
		return nil, err
	}
	vc := VideoComment{VideoId: vid, WriterId: uid, Content: commenttext, CreateTime: time.Now()}
	if err := tx.Model(&VideoComment{}).Create(&vc).Error; err != nil {
		tx.Rollback()
		return nil, err
	}
	if err := tx.Model(&DBVideo{}).
		Where("id = ?", vid).
		UpdateColumn("comment_count", gorm.Expr("comment_count + ?", 1)).Error; err != nil {
		tx.Rollback()
		return nil, err
	}
	return &vc, tx.Commit().Error

}
func DeleteVideoComment(commentId, vid int64) error {
	//
	//if err := db.Model(&VideoComment{}).Where("id = ?", commentId).Delete(&vc).Error; err != nil {
	//	return err
	//}
	//return nil
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if err := tx.Error; err != nil {
		return err
	}
	vc := VideoComment{}
	if err := db.Model(&VideoComment{}).
		Where("id = ?", commentId).
		Delete(&vc).Error; err != nil {
		tx.Rollback()
		return err
	}
	if err := tx.Model(&DBVideo{}).
		Where("id = ?", vid).
		UpdateColumn("comment_count", gorm.Expr("comment_count - ?", 1)).Error; err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}
func QueryVideoCommentWriterID(commentId int64) (int64, error) {
	vc := VideoComment{}
	if err := db.Model(&VideoComment{}).Where("id = ?", commentId).Find(&vc).Error; err != nil {
		return -1, NotFond
	}
	return vc.WriterId, nil
}
func QueryVideoCommentList(vid int64) ([]VideoComment, map[int64]User, error) {
	vc := make([]VideoComment, 0)
	rst := db.Model(&VideoComment{}).
		Where("video_id = ?", vid).
		Order("create_date desc").Find(&vc)
	if rst.Error != nil {
		return nil, nil, rst.Error
	}

	type WriterId int64
	var uids []WriterId
	rst.Select("writer_id").Scan(&uids)
	users := make([]User, 0)
	if rst = db.Model(&User{}).
		Where("id in (?)", uids).Find(&users); rst.Error != nil {
		return nil, nil, rst.Error
	}
	writerMap := make(map[int64]User, len(users))
	for _, usr := range users {
		writerMap[usr.Id] = usr
	}
	return vc, writerMap, nil
}

func QueryVideoAuthorID(vid int64) int64 {
	video := make([]DBVideo, 0)
	if err := db.Model(&DBVideo{}).Where("id = ?", vid).Find(&video).Error; err != nil || len(video) == 0 {
		return -1
	}
	return video[0].AuthorId
}
