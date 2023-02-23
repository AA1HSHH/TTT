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
	VideoList []dal.FeedVideo `json:"video_list,omitempty"`
	NextTime  int64           `json:"next_time,omitempty"`
}

// /douyin/feed/
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
	token := c.Query("token")           // 返回登录用户 token
	// token, isLogin := c.GetQuery("token") // 返回登录用户 token
	log.Printf("传入时间戳(字符串格式): " + inputTime)
	log.Printf("传入 token: " + token)

	// 检测 GET 表单信息
	// c.JSON(http.StatusOK, gin.H{
	// 	"latest_time": inputTime,
	// 	"token":       token,
	// })

	var feedRefTime time.Time // 根据参考时间返回视频流
	feedRefTime = time.Now()
	log.Printf("当前时间: %v", feedRefTime)
	if inputTime != "0" && inputTime != "" {
		intInputTime, _ := strconv.ParseInt(inputTime, 10, 64) // 字符串转十进制整型

		if intInputTime < 253402185600 { // 1 6762 0660 3034  抖声初始化时间戳越界
			feedRefTime = time.Unix(intInputTime, 0) // 自 1970 年 1 月 1 日 UTC 以来经过 intInputTime 秒
		}
		log.Printf("传入时间: %v", feedRefTime.UTC()) // 以 UTC 显示
	}
	log.Printf("视频流返回时间: %v", feedRefTime)

	// feedRefTime = time.Now()

	// 拉取视频流
	var (
		feedVideoList []dal.FeedVideo
		nextTime      time.Time
		err           error
	)
	_, name, _ := mw.TokenStringGetUser(token)  // 根据 token 解析用户 id
	if isLogin := dal.IsExist(name); !isLogin { // 未登录用户获取的视频流
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
		if err != nil {
			c.JSON(http.StatusOK, UserResponse{
				Response: Response{
					StatusCode: 1,
					StatusMsg:  "Authen failed"},
			})
			return
		}
		log.Printf("登录用户 id: %v, 登录用户名: %v", id, name)

		feedVideoList, nextTime, err = service.QueryFeedVideoList(id, feedRefTime) // id

		if err != nil {
			c.JSON(http.StatusOK, FeedResponse{
				Response: Response{
					StatusCode: 1,
					StatusMsg:  "已登录用户拉取视频失败"},
			})
		}
	}
	log.Printf("下次刷新时间戳: %v", nextTime) //nextTime.Unix()

	log.Printf("Feed 流长度: %v", len(feedVideoList))

	c.JSON(http.StatusOK, FeedResponse{
		Response: Response{
			StatusCode: 0,
			StatusMsg:  "feed 接口正常运行"},
		NextTime:  nextTime.Unix(), // time.Now().Unix(),
		VideoList: feedVideoList,   // DemoVideos,
	})
}
