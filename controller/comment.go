package controller

import (
	"github.com/AA1HSHH/TTT/dal"
	"github.com/AA1HSHH/TTT/mw"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type CommentListResponse struct {
	Response
	CommentList []APIComment `json:"comment_list,omitempty"`
}

type CommentActionResponse struct {
	Response
	Comment APIComment `json:"comment,omitempty"`
}

// CommentAction no practical effect, just check if token is valid
func CommentAction(c *gin.Context) {
	token := c.Query("token")
	videoId, e := strconv.ParseInt(c.Query("video_id"), 10, 64)
	if e != nil {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "Wrong video_id"})
		return
	}
	actionType := c.Query("action_type")
	myId, _, err := mw.TokenStringGetUser(token)
	if err != nil {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "Authed failed"})
		return
	}
	switch actionType {
	case "1":
		{
			commentText := c.Query("comment_text")
			rst, err := dal.CreateVideoComment(videoId, myId, commentText)
			if err != nil {
				c.JSON(http.StatusOK, CommentActionResponse{
					Response: Response{StatusCode: 1,
						StatusMsg: "Create error"},
				})
				return
			}
			user, _ := dal.QueryUserbyId(myId)
			videoAuthorID := dal.QueryVideoAuthorID(videoId)
			isFollow := dal.IsRelationFollow(videoAuthorID, myId)

			comment := APIComment{
				Id: rst.Id,
				User: APIUser{Id: user.Id, Name: user.Name,
					FollowCount: user.FollowCount, FollowerCount: user.FollowerCount, IsFollow: isFollow},
				Content:    rst.Content,
				CreateDate: rst.CreateTime.Format("01-02"),
			}
			c.JSON(http.StatusOK, CommentActionResponse{
				Response: Response{StatusCode: 0,
					StatusMsg: "success"},
				Comment: comment,
			})
			return
		}
	case "2":
		{
			commentId, _ := strconv.ParseInt(c.Query("comment_id"), 10, 64)
			// check user, only delete same writerid
			writerID, err := dal.QueryVideoCommentWriterID(commentId)
			if err != nil || writerID != myId {
				c.JSON(http.StatusOK, CommentActionResponse{
					Response: Response{StatusCode: 1,
						StatusMsg: "Invalid delete"},
				})
				return
			}

			err = dal.DeleteVideoComment(commentId, videoId)
			if err != nil {
				c.JSON(http.StatusOK, CommentActionResponse{
					Response: Response{StatusCode: 1,
						StatusMsg: "Delete error"},
				})
				return
			}
			c.JSON(http.StatusOK, CommentActionResponse{
				Response: Response{StatusCode: 0,
					StatusMsg: "success"},
			})
			return
		}

	}

}

// CommentList all videos have same demo comment list
func CommentList(c *gin.Context) {
	token := c.Query("token")
	videoId, _ := strconv.ParseInt(c.Query("video_id"), 10, 64)
	userId, _, err := mw.TokenStringGetUser(token)
	if err != nil {
		c.JSON(http.StatusOK, CommentListResponse{
			Response: Response{StatusCode: 1, StatusMsg: "Authen failed"},
		})
		return
	}
	videoCommentList, writermap, _ := dal.QueryVideoCommentList(videoId)
	followList := dal.QueryUserFollowList(userId)
	comments := constrctCommentList(videoCommentList, writermap, followList)
	c.JSON(http.StatusOK, CommentListResponse{
		Response:    Response{StatusCode: 0},
		CommentList: comments,
	})
}
