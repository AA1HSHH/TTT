package controller

import (
	"fmt"
	"log"
	"net/http"
	"os/exec"
	"path/filepath"
	"strconv"
	"time"

	"github.com/AA1HSHH/TTT/dal"
	"github.com/AA1HSHH/TTT/mw"
	service "github.com/AA1HSHH/TTT/service/video"
	"github.com/gin-gonic/gin"
)

type VideoListResponse struct {
	Response
	VideoList []dal.FeedVideo `json:"video_list"`
}

/**
 * @brief: /douyin/publish/action
 * @请求参数:
 *     - file data
 *     - string token
 *     - string title
 *
 * @返回响应:
 *     - status_code
 *     - status_msg
 */
func Publish(c *gin.Context) {

	// 获取请求参数：鉴权 token
	// 判断登录用户
	token := c.PostForm("token")
	id, name, err := mw.TokenStringGetUser(token) // 根据 token 解析用户 id
	if err != nil {
		c.JSON(http.StatusOK, UserResponse{
			Response: Response{
				StatusCode: 1,
				StatusMsg:  "User doesn't exist"},
		})
		return
	}
	log.Printf("登录用户 id: %v, 登录用户名: %v", id, name)

	// 获取请求参数: 视频名
	title := c.PostForm("title")
	log.Printf("上传视频名: " + title)

	// 判断视频名合法性

	// 获取请求参数：上传视频文件 data
	data, err := c.FormFile("data")
	if err != nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}

	// 视频上传至服务器(本地文件夹)
	// 后续考虑 ftp 服务器
	// fileName := filepath.Base(data.Filename)  // bear.mp4
	// fileExt := path.Ext(fileName)             // .mp4  用于判断上传格式合法性
	// filename_wo_ext := strings.TrimSuffix(fileName, fileExt)  // bear  用于判断视频名合法性
	videoName := fmt.Sprintf("%d_%s_%s", id, title, time.Now().Format("2006-01-02-150405")) // 后续需修改
	saveFile := filepath.Join("./public/", videoName)
	if err := c.SaveUploadedFile(data, saveFile); err != nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}
	log.Printf("保存视频文件为: %v", saveFile)

	// ffmpeg 对上传视频进行截图
	// "ffmpeg -ss 00:00:01 -i /home/ftpuser/video/" + videoName + ".mp4 -vframes 1 /home/ftpuser/images/" + imageName + ".jpg"

	coverName := fmt.Sprintf("%d_%s_%s.jpg", id, title, time.Now().Format("2006-01-02-150405"))
	coverFile := filepath.Join("./public/", coverName)
	cmd := exec.Command(
		"ffmpeg",
		"-ss", "00:00:01",
		"-i", saveFile,
		"-vframes", "1",
		coverFile,
	)
	cmd.Run()

	// 上传视频信息入数据库
	err = service.Publish(title, id, videoName, coverName)
	if err != nil {
		log.Printf("视频信息入库失败: %v", err)
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}
	log.Printf("视频信息入库成功")

	c.JSON(http.StatusOK, Response{
		StatusCode: 0,
		StatusMsg:  videoName + " uploaded successfully",
	})
}

/**
 * @brief:
 *    /douyin/publish/list
 *    登录用户的视频发布列表，直接列出用户所有投稿过的视频
 * @请求参数:
 *    - string token
 *    - string user_id
 *
 * @返回响应:
 *    - status_code
 *    - status_msg
 *    - video_list
 */
func PublishList(c *gin.Context) {
	token := c.Query("token")
	user_id := c.Query("user_id")
	log.Printf("传入 token  : " + token)
	log.Printf("传入 user id: " + user_id)
	id, name, err := mw.TokenStringGetUser(token)
	if err != nil {
		c.JSON(http.StatusOK, VideoListResponse{
			Response: Response{
				StatusCode: 1,
				StatusMsg:  "User doesn't exist"},
		})
		return
	}
	log.Printf("登录用户 id: %v, 登录用户名: %v", id, name)

	// if (user_id == id)

	user_id_i64, _ := strconv.ParseInt(user_id, 10, 64)
	publish_list, err := service.GetPublishList(user_id_i64)

	if err != nil {
		log.Printf("service.GetPublishList(%v)错误：%v\n", user_id, err)
		c.JSON(http.StatusOK, VideoListResponse{
			Response: Response{
				StatusCode: 1,
				StatusMsg:  "获取视频列表失败"},
		})
		return
	}
	log.Printf("service.GetPublishList(%v)成功", user_id)

	c.JSON(http.StatusOK, VideoListResponse{
		Response: Response{
			StatusCode: 0,
		},
		VideoList: publish_list,
	})
}
