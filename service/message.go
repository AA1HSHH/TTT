package service

//
//import (
//	"encoding/json"
//	"fmt"
//	"github.com/AA1HSHH/TTT/cache"
//	"github.com/AA1HSHH/TTT/dal"
//	"github.com/AA1HSHH/TTT/pkg/e"
//	"github.com/gin-gonic/gin"
//	"github.com/gorilla/websocket"
//	"log"
//	"strconv"
//	"time"
//)
//
//const month = 60 * 60 * 24 * 30 // 按照30天算一个月
//
//// 发送消息的类型
//type SendMsg struct {
//	Type    int    `json:"type"`
//	Content string `json:"content"`
//}
//
//// 回复的消息
//type ReplyMsg struct {
//	From    string `json:"from"`
//	Code    int    `json:"code"`
//	Content string `json:"content"`
//}
//
//// 用户类
//type Client struct {
//	ID     string
//	SendID string
//	Socket *websocket.Conn
//	Send   chan []byte
//}
//
//// Message 信息转JSON (包括：发送者、接收者、内容)
//type Message struct {
//	Sender    string `json:"sender,omitempty"`
//	Recipient string `json:"recipient,omitempty"`
//	Content   string `json:"content,omitempty"`
//}
//type SendMsgService struct {
//	UserId   string `form:"userId" json:"userId" binding:""`
//	ToUserId string `form:"to_user_id" json:"to_user_id" binding:""`
//	Content  string `form:"content" json:"password"`
//}
//
//func (service SendMsgService) WsHandler(c *gin.Context) {
//	uid := c.Query("uid")     // 自己的id
//	toUid := c.Query("toUid") // 对方的id
//
//	content = c.Query("content")
//
//	if UserIsExist(uid)
//
//	//conn, err := (&websocket.Upgrader{
//	//	CheckOrigin: func(r *http.Request) bool { // CheckOrigin解决跨域问题
//	//		return true
//	//	}}).Upgrade(c.Writer, c.Request, nil) // 升级成ws协议
//	//if err != nil {
//	//	http.NotFound(c.Writer, c.Request)
//	//	return
//	//}
//	// 创建一个用户实例
//	//client := &Client{
//	//	ID:     createId(uid, toUid),
//	//	SendID: createId(toUid, uid),
//	//	Socket: conn,
//	//	Send:   make(chan []byte),
//	//}
//	//// 用户注册到用户管理上
//	//Manager.Register <- client
//	//go client.Read()
//	//go client.Write()
//}
//
////func (c *Client) Read() {
////	defer func() { // 避免忘记关闭，所以要加上close
////		Manager.Unregister <- c
////		_ = c.Socket.Close()
////	}()
////	for {
////		c.Socket.PongHandler()
////		sendMsg := new(SendMsg)
////		// _,msg,_:=c.Socket.ReadMessage()
////		err := c.Socket.ReadJSON(&sendMsg) // 读取json格式，如果不是json格式，会报错
////		if err != nil {
////			log.Println("数据格式不正确", err)
////			Manager.Unregister <- c
////			_ = c.Socket.Close()
////			break
////		}
////		if sendMsg.Type == 1 { //  发送消息
////			r1, _ := cache.RedisClient.Get(c.ID).Result()
////			r2, _ := cache.RedisClient.Get(c.SendID).Result()
////			if r1 >= "3" && r2 == "" { // 限制单聊
////				replyMsg := ReplyMsg{
////					Code:    e.WebsocketLimit,
////					Content: "达到限制",
////				}
////				msg, _ := json.Marshal(replyMsg)
////				_ = c.Socket.WriteMessage(websocket.TextMessage, msg)
////				_, _ = cache.RedisClient.Expire(c.ID, time.Hour*24*30).Result() // 防止重复骚扰，未建立连接刷新过期时间一个月
////				continue
////			} else {
////				cache.RedisClient.Incr(c.ID)
////				_, _ = cache.RedisClient.Expire(c.ID, time.Hour*24*30*3).Result() // 防止过快“分手”，建立连接三个月过期
////			}
////			log.Println(c.ID, "发送消息", sendMsg.Content)
////			Manager.Broadcast <- &Broadcast{
////				Client:  c,
////				Message: []byte(sendMsg.Content),
////			}
////		} else if sendMsg.Type == 2 { //拉取历史消息
////			timeT, err := strconv.Atoi(sendMsg.Content) // 传送来时间
////			if err != nil {
////				timeT = 999999999
////			}
////			results, _ := dal.FindMany(dal.DbName, c.SendID, c.ID, int64(timeT), 10)
////			if len(results) > 10 {
////				results = results[:10]
////			} else if len(results) == 0 {
////				replyMsg := ReplyMsg{
////					Code:    e.WebsocketEnd,
////					Content: "到底了",
////				}
////				msg, _ := json.Marshal(replyMsg)
////				_ = c.Socket.WriteMessage(websocket.TextMessage, msg)
////				continue
////			}
////			for _, result := range results {
////				replyMsg := ReplyMsg{
////					From:    result.From,
////					Content: fmt.Sprintf("%s", result.Msg),
////				}
////				msg, _ := json.Marshal(replyMsg)
////				_ = c.Socket.WriteMessage(websocket.TextMessage, msg)
////			}
////		}
////	}
////}
////
////func (c *Client) Write() {
////	defer func() {
////		_ = c.Socket.Close()
////	}()
////	for {
////		select {
////		case message, ok := <-c.Send:
////			if !ok {
////				_ = c.Socket.WriteMessage(websocket.CloseMessage, []byte{})
////				return
////			}
////			log.Println(c.ID, "接受消息:", string(message))
////			replyMsg := ReplyMsg{
////				Code:    e.WebsocketSuccessMessage,
////				Content: fmt.Sprintf("%s", string(message)),
////			}
////			msg, _ := json.Marshal(replyMsg)
////			_ = c.Socket.WriteMessage(websocket.TextMessage, msg)
////		}
////	}
////}
//
////var chatConnMap = sync.Map{}
////
////func RunMessageServer() {
////	listen, err := net.Listen("tcp", "127.0.0.1:9090")
////	if err != nil {
////		fmt.Printf("Run message sever failed: %v\n", err)
////		return
////	}
////
////	for {
////		conn, err := listen.Accept()
////		if err != nil {
////			fmt.Printf("Accept conn failed: %v\n", err)
////			continue
////		}
////
////		go process(conn)
////	}
////}
////
////func process(conn net.Conn) {
////	defer conn.Close()
////
////	var buf [256]byte
////	for {
////		n, err := conn.Read(buf[:])
////		if n == 0 {
////			if err == io.EOF {
////				break
////			}
////			fmt.Printf("Read message failed: %v\n", err)
////			continue
////		}
////
////		var event = controller.MessageSendEvent{}
////		_ = json.Unmarshal(buf[:n], &event)
////		fmt.Printf("Receive Message：%+v\n", event)
////
////		fromChatKey := fmt.Sprintf("%d_%d", event.UserId, event.ToUserId)
////		if len(event.MsgContent) == 0 {
////			chatConnMap.Store(fromChatKey, conn)
////			continue
////		}
////
////		toChatKey := fmt.Sprintf("%d_%d", event.ToUserId, event.UserId)
////		writeConn, exist := chatConnMap.Load(toChatKey)
////		if !exist {
////			fmt.Printf("User %d offline\n", event.ToUserId)
////			continue
////		}
////
////		pushEvent := controller.MessagePushEvent{
////			FromUserId: event.UserId,
////			MsgContent: event.MsgContent,
////		}
////		pushData, _ := json.Marshal(pushEvent)
////		_, err = writeConn.(net.Conn).Write(pushData)
////		if err != nil {
////			fmt.Printf("Push message failed: %v\n", err)
////		}
////	}
////}
