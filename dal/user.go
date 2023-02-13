package dal

import "errors"

type User struct {
	Id            int64  `gorm:"column:id"`
	Name          string `gorm:"column:name"`
	Passwd        string `gorm:"column:passwd"`
	FollowCount   int64  `gorm:"column:follow_count"`
	FollowerCount int64  `gorm:"column:follower_count"`
}

var (
	NotFond = errors.New("User Not Fond")
)

func (User) TableName() string {
	return "t_user"
}

func IsExist(name string) bool {
	_, err := QueryUserIDbyName(name)
	if err != nil {
		return false
	}
	return true
}
func CreateUser(name string, passwd string) (int64, error) {
	user := User{Name: name, Passwd: passwd}
	rst := db.Create(&user)
	return user.Id, rst.Error
}
func QueryUserIDbyName(name string) (int64, error) {
	users := make([]*User, 0)
	rst := db.Where("name = ?", name).Find(&users)
	if rst.Error != nil || rst.RowsAffected == 0 {
		return int64(-1), NotFond
	}
	return users[0].Id, nil
}
func QueryUserIDbyNamePasswd(name string, passwd string) (int64, error) {
	users := make([]*User, 0)
	rst := db.Where("name = ?", name).Where("passwd = ?", passwd).Find(&users)
	if rst.Error != nil || rst.RowsAffected == 0 {
		return int64(-1), NotFond
	}
	return users[0].Id, nil
}
func QueryUserbyId(id int64) (*User, error) {
	users := make([]*User, 0)
	rst := db.Where("id = ?", id).Find(&users)
	if rst.Error != nil || rst.RowsAffected == 0 {
		return users[0], NotFond
	}
	return users[0], nil
}
func IsRelationFollow(uid int64, myid int64) bool {
	if uid < 0 || myid < 0 {
		return false
	}
	if uid == myid {
		return true
	}
	relation := make([]Tfollow, 0)
	rst := db.Model(&Tfollow{}).Where("user_id = ?", uid).Where("fans_id = ?", myid).Find(&relation)
	if rst.Error != nil || len(relation) == 0 {
		return false
	}
	return true
}
func QueryUserFollowList(uid int64) map[int64]struct{} {
	relation := make([]Tfollow, 0)
	rst := db.Table("t_follow").Where("fans_id = ?", uid).Find(&relation)
	if rst.Error != nil {
		return nil
	}
	followid := make(map[int64]struct{}, len(relation))
	for _, item := range relation {
		followid[item.user_id] = struct{}{}
	}
	return followid
}
