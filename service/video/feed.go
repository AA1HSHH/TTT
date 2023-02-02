package service

import (
	"fmt"
	"time"

	"github.com/AA1HSHH/TTT/dal"
)

// 根据 输入时间戳 和 用户 id 获取视频流
func QueryFeedVideoList(userId int64, timeStamp time.Time) ([]dal.DBVideo, time.Time, error) {

	// 未登录用户按时间获取视频流
	// 登录用户按用户画像推荐（关注用户投稿视频流）

	dbVideos, err := dal.QueryVideosByTime(timeStamp)
	if err != nil {
		fmt.Println("数据库获取视频流信息")
	}
	return dbVideos, dbVideos[0].PublishTime, nil
}
