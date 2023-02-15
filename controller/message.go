package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

//var tempChat = map[string][]Message{}
//
//var messageIdSequence = int64(1)
//
//type ChatResponse struct {
//	Response
//	MessageList []Message `json:"message_list"`
//}

// MessageAction no practical effect, just check if token is valid
//func MessageAction(c *gin.Context) {
//	var sendMsgService service.SendMsgService
//	token := c.Query("token")
//	//userId, _ := strconv.ParseInt(c.Query("user_id"), 10, 64)
//	//toUserId := c.Query("to_user_id")
//	//content := c.Query("content")
//
//	_, _, err := mw.TokenStringGetUser(token)
//	if err != nil {
//		c.JSON(http.StatusOK, UserResponse{
//			Response: Response{StatusCode: 1, StatusMsg: "Authen failed"},
//		})
//		return
//	}
//
//	if err := c.ShouldBind(&sendMsgService); err == nil {
//		sendMsgService.WsHandler(c)
//		c.JSON(http.StatusOK, Response{
//			StatusCode: 0, StatusMsg: "Message send successfully"})
//	} else {
//		c.JSON(http.StatusOK, Response{
//			StatusCode: 1, StatusMsg: "Message send failed"})
//	}
//
//	//userIdB, _ := strconv.Atoi(toUserId)
//	//chatKey := genChatKey(userId, int64(userIdB))
//	//
//	//atomic.AddInt64(&messageIdSequence, 1)
//	//curMessage := Message{
//	//	Id:         messageIdSequence,
//	//	Content:    content,
//	//	CreateTime: time.Now().Format(time.Kitchen),
//	//}
//	//
//	//if messages, exist := tempChat[chatKey]; exist {
//	//	tempChat[chatKey] = append(messages, curMessage)
//	//} else {
//	//	tempChat[chatKey] = []Message{curMessage}
//	//}
//
//	//c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "Not implement"})
//}

// MessageChat all users have same follow list
func MessageChat(c *gin.Context) {
	//token := c.Query("token")
	//toUserId := c.Query("to_user_id")
	//
	//if user, exist := usersLoginInfo[token]; exist {
	//	userIdB, _ := strconv.Atoi(toUserId)
	//	chatKey := genChatKey(user.Id, int64(userIdB))
	//
	//	c.JSON(http.StatusOK, ChatResponse{Response: Response{StatusCode: 0}, MessageList: tempChat[chatKey]})
	//} else {
	//	c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
	//}
	c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "Not implement"})
}

func genChatKey(userIdA int64, userIdB int64) string {
	if userIdA > userIdB {
		return fmt.Sprintf("%d_%d", userIdB, userIdA)
	}
	return fmt.Sprintf("%d_%d", userIdA, userIdB)
}
