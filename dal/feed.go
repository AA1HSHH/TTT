package dal

import (
	"time"
)

type DBVideo struct {
	Id          int64     `json:"id"`
	AuthorId    int64     `json:"author_id"` // 根据作者 id 去用户表找作者信息
	Title       string    `json:"title"`
	PublishTime time.Time `json:"publish_time"`
	PlayUrl     string    `json:"play_url,omitempty"`
	CoverUrl    string    `json:"cover_url,omitempty"`
	// FavoriteCount int64     `json:"favorite_count,omitempty"`
	// CommentCount  int64     `json:"comment_count,omitempty"`
	// IsFavorite    bool      `json:"is_favorite,omitempty"`
}

func (DBVideo) TableName() string {
	return "t_video"
}

func QueryVideosByTime(timeStamp time.Time) ([]DBVideo, error) {
	videoCount := 1 // 写进配置文件
	videos := make([]DBVideo, videoCount)
	result := db.Where("publish_time<?", timeStamp).Order("publish_time desc").Limit(videoCount).Find(&videos)
	if result.Error != nil {
		return videos, result.Error
	}
	return videos, nil
}
