package dal

import (
	"errors"
	"fmt"
)

//	type UserTemplate struct {
//		Id             int64  `gorm:"column:id"`
//		Name           string `gorm:"column:name"`
//		FollowCount    int64  `gorm:"column:follow_count"`
//		FollowerCount  int64  `gorm:"column:follower_count"`
//		Avatar         string `gorm:"column:avatar"`
//		BackgroundImg  string `gorm:"column:background_image"`
//		Signature      string `gorm:"column:signature"`
//		TotalFavorited int64  `gorm:"column:total_favorited"`
//		WorkCnt        int64  `gorm:"column:work_count"`
//		FavoriteCnt    int64  `gorm:"column:favorite_count"`
//	}
type UserInfo struct {
	Id            int64  `json:"id"`
	Name          string `json:"name"`
	FollowCount   int64  `json:"follow_count"`
	FollowerCount int64  `json:"follower_count"`
	IsFollow      bool   `json:"is_follow"`
	// 下为用户新增字段
	Avatar          string `json:"avatar"`
	BackgroundImage string `json:"background_image"`
	Signature       string `json:"signature"`
	TotalFavorited  int64  `json:"total_favorited"`
	WorkCount       int64  `json:"work_count"`
	FavoriteCount   int64  `json:"favorite_count"`
}

type FriendUser struct {
	UserInfo
	Content string `gorm:"column:content" json:"message"` //首字母必须要大写，否则接受不了数据
	MsgType int64  `gorm:"column:MsgType"  json:"msgType"`
}

type Tfollow struct {
	user_id int64 `json:"user_id,omitempty"`
	fans_id int64 `json:"fans_id,omitempty"`
}

func IsUserExistById(id int64) bool {

	var tfollow []Tfollow // 仅仅使用user_id
	if err := db.Raw("SELECT `id`,`id` from `t_user` WHERE id =?", id).Scan(&tfollow).Error; err != nil {

		return false
	}
	if len(tfollow) == 0 {
		return false
	}
	return true
}

func AddUserFollow(userId, followId int64) error {

	if err := db.Exec("INSERT INTO `t_follow` (`fans_id`,`user_id`) VALUES (?,?)", userId, followId).Error; err != nil {
		return err
	}
	if err := db.Exec("UPDATE t_user SET follow_count=follow_count+1 WHERE id = ?", userId).Error; err != nil {
		return err
	}
	if err := db.Exec("UPDATE t_user SET follower_count=follower_count+1 WHERE id = ?", followId).Error; err != nil {
		return err
	}

	return nil
}

func CancelUserFollow(userId, followId int64) error {
	//var tfollow []Tfollow
	//if err := db.Raw("SELECT user_id,fans_id from t_follow where  user_id = ? AND fans_id = ?", userId, userToId).Scan(&tfollow).Error; err != nil {
	//	return err
	//}
	////fmt.Println("tfollow",tfollow[0].fans_id, db.RowsAffected)
	//if len(tfollow)==1 && tfollow[0].user_id ==0 && tfollow[0].fans_id==0{
	//	return errors.New("There is no relationship, no need to cancel")
	//}
	db_eff := db.Exec("DELETE from t_follow  where  user_id = ? AND fans_id = ?", followId,userId)
	if db_eff.Error != nil {
		return db_eff.Error
	}
	if db_eff.RowsAffected == 0 {
		return errors.New("There is no relationship, no need to cancel")
	}

	if err := db.Exec("UPDATE t_user SET follow_count=follow_count-1 WHERE id = ?", userId).Error; err != nil {
		return err
	}
	if err := db.Exec("UPDATE t_user SET follower_count=follower_count-1 WHERE id = ?", followId).Error; err != nil {
		return err
	}

	return nil
}

func GetFollowListByUserId(userId int64) ([]UserInfo, error) {
	var err error
	var results []UserInfo

	fmt.Println("userId:", userId)
	//if err = db.Raw("SELECT u.id as id ,u.name as name , u.follow_count as follow_count,u.follower_count as follower_count,1 as is_follow FROM t_follow r, t_user u WHERE r.fans_id = ? AND r.user_id = u.id", userId).Scan(&results).Error; err != nil {
	if err = db.Raw("SELECT u.*,1 as is_follow FROM t_follow r, t_user u WHERE r.fans_id = ? AND r.user_id = u.id", userId).Scan(&results).Error; err != nil {
		fmt.Println(results)
		return results, err
	} else {
		fmt.Println(results)
	}
	if len(results) == 0 || (results)[0].Id == 0 {
		return results, errors.New("List is empty")
	}
	return results, nil
}

func GetFollowerListByUserId(userId int64) ([]UserInfo, error) {
	var err error
	var results []UserInfo

	fmt.Println("userId:", userId)
	//if err = db.Raw("SELECT u.*,coalesce(r.BID,0) as is_follow From t_user as  u,(select A.AID as AID ,B.BID as BID from  (select fans_id as AID from t_follow  where user_id = ?) as A left join  (select a.user_id as BID from t_follow as a inner join t_follow as b on a.fans_id = ? and b.fans_id = ?) as B on A.AID = B.BID ) as r where  r.AID = u.id",userId,userId,userId).Scan(&results).Error; err != nil {
	if err = db.Raw("SELECT u.*,1 as is_follow FROM t_follow r, t_user u WHERE r.user_id = ? AND r.fans_id = u.id", userId).Scan(&results).Error; err != nil {
		fmt.Println("err", err)
		return results, err
	} else {
		fmt.Println(results)
	}
	if len(results) == 0 || (results)[0].Id == 0 {
		return results, errors.New("List is empty")
	}
	return results, nil
}

func GetChat(userId int64) ([]FriendUser, error) {
	var err error
	var results []FriendUser

	fmt.Println("userId:", userId)
	if err = db.Raw("SELECT t.*,MM.content as content,IF (MM.from= follower_id,0,1) as MsgType "+
		"from (SELECT f1.fans_id as follower_id FROM t_follow f1 "+
		"JOIN t_follow f2 ON f1.fans_id = f2.user_id AND f1.user_id = f2.fans_id "+
		"WHERE f1.user_id = ? ) B "+
		"left join ( SELECT * from (SELECT * FROM  t_message  having 1 ORDER BY create_time desc) t GROUP BY t.from,t.to) MM "+
		"on B.follower_id =  MM.to or B.follower_id =  MM.from "+
		"join t_user t "+
		"on t.id = B.follower_id "+
		"where MM.from = ? or MM.to =? "+
		"GROUP BY B.follower_id;",  userId,  userId,  userId).Scan(&results).Error; err != nil {
		fmt.Println("err", err)
		return results, err
	} else {
		fmt.Println(results)
	}
	if len(results) == 0 || (results)[0].Id == 0 {
		return results, errors.New("List is empty!")
	}
	fmt.Println(results)
	return results, nil
}
