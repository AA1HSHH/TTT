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
	vc := VideoComment{VideoId: vid, WriterId: uid, Content: commenttext, CreateTime: time.Now()}
	if err := db.Model(&VideoComment{}).Create(&vc).Error; err != nil {
		return nil, err
	}
	return &vc, nil
}
func DeleteVideoComment(commentId int64) error {
	vc := VideoComment{}
	if err := db.Model(&VideoComment{}).Where("id = ?", commentId).Delete(&vc).Error; err != nil {
		return err
	}
	return nil
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
