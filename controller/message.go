package controller

import (
	"fmt"
	"github.com/AA1HSHH/TTT/dal"
	"github.com/AA1HSHH/TTT/mw"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type ChatResponse struct {
	Response
	MessageList []dal.Message `json:"message_list"`
}

// MessageAction send a message less than 60 words to one user,message can not be null
func MessageAction(c *gin.Context) {
	token := c.Query("token")
	userId, _, err := mw.TokenStringGetUser(token)
	if err != nil {
		c.JSON(http.StatusOK, UserResponse{
			Response: Response{StatusCode: 1, StatusMsg: "Authen failed"},
		})
		return
	}
	toUserId, _ := strconv.ParseInt(c.Query("to_user_id"), 10, 64)
	content := c.Query("content")
	if len(content) == 0 {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 1,
				StatusMsg: "Message can't be null"},
		})
		return
	}
	if exit := dal.UserIsExist(userId); !exit {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 1, StatusMsg: "User doesn't exist"},
		})
		return
	}
	if exit := dal.UserIsExist(toUserId); !exit {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 1,
				StatusMsg: "User sent message to doesn't exist"},
		})
		return
	}

	if err := dal.CreateMessage(userId, toUserId, content); err == nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: 0, StatusMsg: "Message send successfully"})
	} else {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1, StatusMsg: "Message send failed"})
	}
}

// MessageChat all users have same follow list,get all messages with one user
func MessageChat(c *gin.Context) {
	token := c.Query("token")
	toUserId, _ := strconv.ParseInt(c.Query("to_user_id"), 10, 64)
	preMsgTime, _ := strconv.ParseInt(c.Query("pre_msg_time"), 10, 64)

	if exit := dal.UserIsExist(toUserId); !exit {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 1,
				StatusMsg: "User sent message to doesn't exist"},
		})
		return
	}
	myid, _, err := mw.TokenStringGetUser(token)
	if err != nil {
		c.JSON(http.StatusOK, UserResponse{
			Response: Response{StatusCode: 1, StatusMsg: "Authen failed"},
		})
		return
	}
	if msgList, err := dal.QueryMessage(myid, toUserId, preMsgTime); err == nil {
		c.JSON(http.StatusOK, ChatResponse{
			Response: Response{
				StatusCode: 0,
				StatusMsg:  "Query message list successfully"},
			MessageList: msgList,
		})
		fmt.Println(msgList)
	} else {
		c.JSON(http.StatusOK, ChatResponse{
			Response: Response{
				StatusCode: 1,
				StatusMsg:  "Query message list failed"},
			MessageList: msgList,
		})
	}
}
