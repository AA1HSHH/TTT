package dal

type UserFollow struct {
	UserId int64 `json:"user_id"`
	FansId int64 `json:"fans_id"`
}

func (UserFollow) TableName() string {
	return "t_follow"
}

func IsFollow(fansId int64, userId int64) bool {
	fans := make([]UserFollow, 0)
	rst := db.Where(&UserFollow{UserId: userId, FansId: fansId}).Find(&fans)
	if rst.Error != nil || len(fans) == 0 {
		return false
	}
	return true
}
