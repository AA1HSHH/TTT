package service

import (
	"fmt"
	"log"
	"time"

	"github.com/AA1HSHH/TTT/dal"
)

// 根据 输入时间戳 和 用户 id 获取视频流
func QueryFeedVideoList(userId int64, timeStamp time.Time) ([]dal.FeedVideo, time.Time, error) {

	var (
		dbVideosList []dal.DBVideo
		err          error
	)
	// 未登录用户按时间获取视频流
	if userId == 0 {
		dbVideosList, err = dal.QueryVideosByTime(timeStamp)
		if err != nil {
			fmt.Println("dal.QueryVideosByTime 错误")
		}
	} else { // 登录用户按用户画像推荐（关注用户投稿视频流）
		dbVideosList, err = dal.QueryVideosByIdTime(userId, timeStamp)
		if err != nil {
			fmt.Println("dal.QueryVideosByIdTime 错误")
		}
	}

	feed_videos, _ := GetFeedVideo(dbVideosList)

	var nextTime time.Time
	if !(len(feed_videos) > 0) {
		nextTime = time.Now()
	} else {
		nextTime = dbVideosList[0].PublishTime
	}
	return feed_videos, nextTime, nil
}

// 获取 FeedVideo, 合并 DBVideo 和 User 数据
func GetFeedVideo(dbvideos []dal.DBVideo) ([]dal.FeedVideo, error) {
	videoCount := len(dbvideos) // config.VideoCount
	log.Printf("用户发布视频数: %v", videoCount)
	feed_videos := make([]dal.FeedVideo, videoCount)
	feed_authors := make([]dal.FeedAuthor, videoCount)
	for index, dbvideo := range dbvideos {
		author_id := dbvideo.AuthorId
		user_info, err := dal.QueryUserbyId(author_id)
		if err == nil {
			feed_authors[index].Id = user_info.Id
			feed_authors[index].Name = user_info.Name
			feed_authors[index].FollowCount = user_info.FollowCount
			feed_authors[index].FollowerCount = user_info.FollowerCount
		}

		feed_videos[index].Id = dbvideo.Id
		feed_videos[index].PlayUrl = dbvideo.PlayUrl
		feed_videos[index].CoverUrl = dbvideo.CoverUrl
		feed_videos[index].FavoriteCount = dbvideo.FavoriteCount
		feed_videos[index].CommentCount = dbvideo.CommentCount
		feed_videos[index].Title = dbvideo.Title
		feed_videos[index].Author = feed_authors[index]

	}
	return feed_videos, nil
}
