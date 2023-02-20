package controller

import (
	"fmt"
	"github.com/AA1HSHH/TTT/dal"
	"github.com/AA1HSHH/TTT/mw"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

//var tempChat = map[string][]Message{}

//var messageIdSequence = int64(1)

type ChatResponse struct {
	Response
	MessageList []dal.Message `json:"message_list"`
}

// MessageAction no practical effect, just check if token is valid
func MessageAction(c *gin.Context) {
	token := c.Query("token")
	userId, _ := strconv.ParseInt(c.Query("user_id"), 10, 64)
	toUserId, _ := strconv.ParseInt(c.Query("to_user_id"), 10, 64)
	content := c.Query("content")

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

	_, _, err := mw.TokenStringGetUser(token)
	if err != nil {
		c.JSON(http.StatusOK, UserResponse{
			Response: Response{StatusCode: 1, StatusMsg: "Authen failed"},
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

	//userIdB, _ := strconv.Atoi(toUserId)
	//chatKey := genChatKey(userId, int64(userIdB))
	//
	//atomic.AddInt64(&messageIdSequence, 1)
	//curMessage := Message{
	//	Id:         messageIdSequence,
	//	Content:    content,
	//	CreateTime: time.Now().Format(time.Kitchen),
	//}
	//
	//if messages, exist := tempChat[chatKey]; exist {
	//	tempChat[chatKey] = append(messages, curMessage)
	//} else {
	//	tempChat[chatKey] = []Message{curMessage}
	//}

	//c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "Not implement"})
}

// MessageChat all users have same follow list
func MessageChat(c *gin.Context) {
	token := c.Query("token")
	toUserId, _ := strconv.ParseInt(c.Query("to_user_id"), 10, 64)
	//preMsgTime := c.Query("pre_msg_time")

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
	//t := time.NewTicker(time.Second * 2)
	//
	//go func() {
	//	for {
	//		select {
	//		case <-t.C:
	//			fmt.Println(1)
	//		}
	//	}
	//}()
	//
	//time.Sleep(10 * time.Second)
	//t.Stop()
	//time.Sleep(10 * time.Second)
	//var msgList []Message
	if msgList, err := dal.QueryMessage(myid, toUserId); err == nil {
		c.JSON(http.StatusOK, ChatResponse{
			Response: Response{
				StatusCode: 0,
				StatusMsg:  "Query message list successfully"},
			MessageList: msgList,
		})
	} else {
		c.JSON(http.StatusOK, ChatResponse{
			Response: Response{
				StatusCode: 1,
				StatusMsg:  "Query message list failed"},
			MessageList: msgList,
		})
	}
	//
	//if user, exist := usersLoginInfo[token]; exist {
	//	userIdB, _ := strconv.Atoi(toUserId)
	//	chatKey := genChatKey(user.Id, int64(userIdB))
	//
	//	c.JSON(http.StatusOK, ChatResponse{Response: Response{StatusCode: 0}, MessageList: tempChat[chatKey]})
	//} else {
	//	c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
	//}
	//c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "Not implement"})
}

func genChatKey(userIdA int64, userIdB int64) string {
	if userIdA > userIdB {
		return fmt.Sprintf("%d_%d", userIdB, userIdA)
	}
	return fmt.Sprintf("%d_%d", userIdA, userIdB)
}
