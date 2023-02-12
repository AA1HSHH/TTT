package dal

import (
	"time"

	"github.com/AA1HSHH/TTT/config"
)

type DBVideo struct {
	Id            int64     `json:"id"`
	AuthorId      int64     `json:"author_id"` // 根据作者 id 去用户表找作者信息
	Title         string    `json:"title"`
	PublishTime   time.Time `json:"publish_time"`
	PlayUrl       string    `json:"play_url,omitempty"`
	CoverUrl      string    `json:"cover_url,omitempty"`
	FavoriteCount int64     `json:"favorite_count,omitempty"`
	CommentCount  int64     `json:"comment_count,omitempty"`
	// IsFavorite    bool      `json:"is_favorite,omitempty"`
}

type FeedAuthor struct {
	Id            int64  `gorm:"column:id"`
	Name          string `gorm:"column:name"`
	FollowCount   int64  `gorm:"column:follow_count"`
	FollowerCount int64  `gorm:"column:follower_count"`
	// IsFollow    bool      `json:"is_follow,omitempty"`
}

type FeedVideo struct {
	Id            int64  `json:"id"`
	PlayUrl       string `json:"play_url,omitempty"`
	CoverUrl      string `json:"cover_url,omitempty"`
	FavoriteCount int64  `json:"favorite_count,omitempty"`
	CommentCount  int64  `json:"comment_count,omitempty"`
	// IsFavorite    bool      `json:"is_favorite,omitempty"`
	Title  string     `json:"title"`
	Author FeedAuthor `json:"feed_author"`
}

func (DBVideo) TableName() string {
	return "t_video"
}

// 未登录用户根据时间戳获取 feed 列表
func QueryVideosByTime(timeStamp time.Time) ([]DBVideo, error) {
	videoCount := config.VideoCount // 写进配置文件
	videos := make([]DBVideo, videoCount)
	result := db.Where("publish_time<?", timeStamp).Order("publish_time desc").Limit(videoCount).Find(&videos)
	if result.Error != nil {
		return videos, result.Error
	}
	return videos, nil
}

// 登录用户根据用户 ID 和 时间戳 获取 feed 列表
func QueryVideosByIdTime(userId int64, timeStamp time.Time) ([]DBVideo, error) {
	videoCount := config.VideoCount // 写进配置文件
	videos := make([]DBVideo, videoCount)

	// 根据用户关注推荐视频流
	result := db.Where("publish_time<?", timeStamp).Order("publish_time desc").Limit(videoCount).Find(&videos)
	if result.Error != nil {
		return videos, result.Error
	}
	return videos, nil
}
