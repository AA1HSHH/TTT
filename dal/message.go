package dal

import (
	"sort"
	"time"
)

type Message struct {
	ID         int64  `json:"id" gorm:"column:id"`
	ToUserId   int64  `json:"to_user_id" gorm:"column:to"`
	FromUserId int64  `json:"from_user_id" gorm:"column:from"`
	Content    string `json:"content" gorm:"column:content"`
	CreateTime int64  `json:"create_time" gorm:"column:create_time"`
}
type Msg []Message

func (msg Msg) Len() int           { return len(msg) }
func (msg Msg) Swap(i, j int)      { msg[i], msg[j] = msg[j], msg[i] }
func (msg Msg) Less(i, j int) bool { return msg[i].CreateTime < msg[j].CreateTime }

func (Message) TableName() string {
	return "t_message"
}

func UserIsExist(id int64) bool {
	_, err := QueryUserbyid(id)
	if err != nil {
		return false
	}
	return true
}
func CreateMessage(id int64, toUid int64, content string) error {
	createTime := time.Now().UnixNano() / 1e6
	print(createTime)
	message := Message{ToUserId: toUid, FromUserId: id, Content: content, CreateTime: createTime}
	rst := db.Create(&message)
	return rst.Error
}
func QueryUserbyid(id int64) (int64, error) {
	users := make([]*User, 0)
	rst := db.Where("id = ?", id).Find(&users)
	if rst.Error != nil || rst.RowsAffected == 0 {
		return int64(-1), NotFond
	}
	return users[0].Id, nil
}
func QueryMessage(id int64, toUid int64, preMsgTime int64) ([]Message, error) {
	var msgList []Message
	var fromList []Message
	var toList []Message
	if preMsgTime == 0 {
		err := db.Where(&Message{FromUserId: id, ToUserId: toUid}).Find(&toList).Error
		if err != nil {
			return msgList, NotFond
		}
		err = db.Where(&Message{FromUserId: toUid, ToUserId: id}).Find(&fromList).Error
		if err != nil {
			return msgList, NotFond
		}
		msgList = append(fromList, toList...)
		sort.Sort(Msg(msgList))
		return msgList, nil
	} else {
		err := db.Where(&Message{FromUserId: id, ToUserId: toUid}).Having("create_time > (?)", preMsgTime+1000).Find(&toList).Error //fix delay bug: 1000ms delay
		if err != nil {
			return msgList, NotFond
		}
		err = db.Where(&Message{FromUserId: toUid, ToUserId: id}).Having("create_time > (?)", preMsgTime+1000).Find(&fromList).Error
		if err != nil {
			return msgList, NotFond
		}
		msgList = append(fromList, toList...)

		sort.Sort(Msg(msgList))
		return msgList, nil
	}

}
