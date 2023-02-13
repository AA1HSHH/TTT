package service

import (
	"fmt"
	"log"
	"time"

	"github.com/AA1HSHH/TTT/config"
	"github.com/AA1HSHH/TTT/dal"
)

// Publish
// 将传入的视频流保存在文件服务器中，并存储在 mysql 表中
func Publish(title string, userId int64, videoName string, coverName string) error {

	var publish_video dal.DBVideo
	// publish_video.Id =
	publish_video.AuthorId = userId
	publish_video.Title = title
	publish_video.PublishTime = time.Now() // .Unix()
	publish_video.PlayUrl = config.PlayUrlPrefix + videoName
	publish_video.CoverUrl = config.CoverUrlPrefix + coverName
	publish_video.FavoriteCount = 0
	publish_video.CommentCount = 0
	// publish_video.IsFavorite = false

	// publish_video := dal.DBVideo{
	// 	Id:,
	// 	AuthorId:,
	// 	Title:,
	// 	PublishTime:,
	// 	PlayUrl:,
	// 	CoverUrl:,
	// }

	// 上传视频信息保存至数据库
	err := dal.AddVideo(publish_video)
	if err != nil {
		log.Printf("视频信息保存至数据库失败%v", err)
		return err
	}
	log.Printf("视频信息保存至数据库成功")

	return nil
}

func GetPublishList(user_id int64) ([]dal.FeedVideo, error) {
	// 根据 user_id 去 t_video 表查询所有 user_id 发布视频
	dbVideos, err := dal.QueryVideosByUserId(user_id)
	if err != nil {
		fmt.Println("数据库获取视频流信息失败")
	}
	publishVideos, _ := GetFeedVideo(user_id, dbVideos)
	return publishVideos, nil
}
