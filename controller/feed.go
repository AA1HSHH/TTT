package controller

import (
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/AA1HSHH/TTT/dal"
	"github.com/AA1HSHH/TTT/mw"
	service "github.com/AA1HSHH/TTT/service/video"
	"github.com/gin-gonic/gin"
)

type FeedResponse struct {
	Response
	VideoList []dal.DBVideo `json:"video_list,omitempty"`
	NextTime  int64         `json:"next_time,omitempty"`
}

// /feed/
// 请求参数:
//   - string latest_time
//   - string token
//
// 返回响应:
//   - status_code
//   - status_msg
//   - next_time
//   - video_list
func Feed(c *gin.Context) {

	// 获取请求参数
	inputTime := c.Query("latest_time") // 返回视频的最新投稿时间戳
	// token := c.Query("token")           // 返回登录用户 token
	token, isLogin := c.GetQuery("token") // 返回登录用户 token
	log.Printf("传入时间: " + inputTime)
	log.Printf("传入 token: " + token)
	log.Printf("登录状态: %v ", isLogin)

	// 检测 GET 表单信息
	// c.JSON(http.StatusOK, gin.H{
	// 	"latest_time": inputTime,
	// 	"token":       token,
	// })

	// 根据传入时间计算参考返回时间戳
	var feedRefTime time.Time // 根据参考时间返回视频流
	if inputTime != "0" {
		intInputTime, err := strconv.ParseInt(inputTime, 10, 64) // 字符串转十进制整型
		if err != nil {
			feedRefTime = time.Unix(0, intInputTime*1e6) // 前端传来的时间戳以 ms 为单位
		}
	} else { // 不输入, 则表示当前时间
		feedRefTime = time.Now()
	}
	log.Printf("视频流返回时间: %v", feedRefTime)

	// 拉取视频流
	var (
		feedVideoList []dal.DBVideo
		nextTime      time.Time
		err           error
	)
	if !isLogin { // 未登录用户获取的视频流
		feedVideoList, nextTime, err = service.QueryFeedVideoList(0, feedRefTime)
		if err != nil {
			c.JSON(http.StatusOK, FeedResponse{
				Response: Response{
					StatusCode: 1,
					StatusMsg:  "未登录用户拉取视频失败"},
			})
		}
	} else { // 已登录用户获取的视频流
		id, name, err := mw.TokenStringGetUser(token) // 根据 token 解析用户 id
		log.Printf("登录用户 id: %v, 登录用户名: %v", id, name)

		if err != nil {
			c.JSON(http.StatusOK, UserResponse{
				Response: Response{
					StatusCode: 1,
					StatusMsg:  "Authen failed"},
			})
			return
		}

		feedVideoList, nextTime, err = service.QueryFeedVideoList(0, feedRefTime) // id
		if err != nil {
			c.JSON(http.StatusOK, FeedResponse{
				Response: Response{
					StatusCode: 1,
					StatusMsg:  "已登录用户拉取视频失败"},
			})
		}
	}

	log.Printf("第一个视频 id: %v", feedVideoList[0].Id)
	log.Printf("下次刷新时间戳: %v", nextTime.Unix())

	// 接口正常返回参数
	// {
	// 	"status_code": 0,
	// 	"status_msg": "string",
	// 	"next_time": 0,
	// 	"video_list": [
	// 		{
	// 			"id": 0,
	// 			"author": {             // TODO
	// 				"id": 0,
	// 				"name": "string",
	// 				"follow_count": 0,
	// 				"follower_count": 0,
	// 				"is_follow": true
	// 			},
	// 			"play_url": "string",
	// 			"cover_url": "string",
	// 			"favorite_count": 0,   // TODO
	// 			"comment_count": 0,    // TODO
	// 			"is_favorite": true,   // TODO
	// 			"title": "string"
	// 		}
	// 	]
	// }
	c.JSON(http.StatusOK, FeedResponse{
		Response: Response{
			StatusCode: 0,
			StatusMsg:  "feed 接口正常运行"},
		NextTime:  nextTime.Unix(), // time.Now().Unix(),
		VideoList: feedVideoList,   // DemoVideos,
	})
}
