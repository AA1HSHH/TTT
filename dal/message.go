package dal

import "time"

type Message struct {
	MsgId   int64     `gorm:"column:id"`
	Id      int64     `gorm:"column:from"`
	ToUid   int64     `gorm:"column:to"`
	Content string    `gorm:"column:content"`
	Time    time.Time `gorm:"column:create_time"`
}

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
	createTime := time.Now()
	message := Message{Id: id, ToUid: toUid, Content: content, Time: createTime}
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
func QueryMessage(id int64, toUid int64) ([]Message, error) {
	var msgList []Message
	err := db.Where(&Message{Id: id, ToUid: toUid}).Order("create_time").Find(&msgList).Error
	if err != nil {
		return msgList, NotFond
	}
	return msgList, nil
}
