package controller

import (
	"github.com/AA1HSHH/TTT/dal"
	"github.com/AA1HSHH/TTT/mw"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type FavoriteListResponse struct {
	Response
	VideoList []APIVideo `json:"video_list"`
}

// FavoriteAction
func FavoriteAction(c *gin.Context) {
	token := c.Query("token")
	videoId, _ := strconv.ParseInt(c.Query("video_id"), 10, 64)
	actionType := c.Query("action_type")
	userId, _, err := mw.TokenStringGetUser(token)
	if err != nil {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "Authed failed"})
		return
	}
	switch actionType {
	case "1":
		{
			if err := dal.CreatFavoirte(videoId, userId); err != nil {
				c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "Favorite failed"})
				return
			}

		}
	case "2":
		{
			if err := dal.DeleteFavoirte(videoId, userId); err != nil {
				c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "Cancel Favorite failed"})
				return
			}
		}
	}
	c.JSON(http.StatusOK, Response{StatusCode: 0, StatusMsg: "success"})
}

// FavoriteList all users have same favorite video list
func FavoriteList(c *gin.Context) {
	token := c.Query("token")
	userId, _ := strconv.ParseInt(c.Query("user_id"), 10, 64)
	_, _, err := mw.TokenStringGetUser(token)
	if err != nil {
		c.JSON(http.StatusOK, FavoriteListResponse{
			Response: Response{StatusCode: 1, StatusMsg: "Authed failed"},
		})
		return
	}
	videos, authors, err := dal.QueryFavoriteList(userId)
	followList := dal.QueryUserFollowList(userId)
	videosList := constructFavoriteVideoList(videos, authors, followList)

	if err != nil {
		c.JSON(http.StatusOK, FavoriteListResponse{
			Response: Response{
				StatusCode: 1,
				StatusMsg:  "Query failed",
			},
		})
		return
	}

	c.JSON(http.StatusOK, FavoriteListResponse{
		Response: Response{
			StatusCode: 0,
			StatusMsg:  "success",
		},
		VideoList: videosList,
	})
}
