package dal

import (
	"errors"
	"fmt"
)

type UserInfo struct {
	Id            int64  `json:"id,omitempty"`
	Name          string `json:"name,omitempty"`
	FollowCount   int64  `json:"follow_count,omitempty"`
	FollowerCount int64  `json:"follower_count,omitempty"`
	IsFollow      bool   `json:"is_follow,omitempty"`
}

type FriendUser struct {
	UserInfo
	Message string `json:"message,omitempty"`
	Msgtype int64 `json:"msg_type,omitempty"`
}



type Tfollow struct {
	user_id    int64  `json:"user_id,omitempty"`
	fans_id    int64 `json:"fans_id,omitempty"`
}


func IsUserExistById(id int64) bool {

	var tfollow []Tfollow // 仅仅使用user_id
	if err:= db.Raw("SELECT `id`,`id` from `t_user` WHERE id =?",id).Scan(&tfollow).Error ; err != nil {

		return false
	}
	if len(tfollow) == 0 {
		return false
	}
	return true
}


func AddUserFollow(userId, userToId int64) error {

	if err := db.Exec("INSERT INTO `t_follow` (`user_id`,`fans_id`) VALUES (?,?)", userId, userToId).Error; err != nil {
		return err
	}
	if err := db.Exec("UPDATE t_user SET follow_count=follow_count+1 WHERE id = ?", userId).Error; err != nil {
		return err
	}
	if err := db.Exec("UPDATE t_user SET follower_count=follower_count+1 WHERE id = ?", userToId).Error; err != nil {
		return err
	}

	return nil
}

func CancelUserFollow(userId, userToId int64) error {
	var tfollow []Tfollow
	if err := db.Raw("SELECT user_id,fans_id from t_follow where  user_id = ? AND fans_id = ?", userId, userToId).Scan(&tfollow).Error; err != nil {
		return err
	}
	//fmt.Println("tfollow",tfollow[0].fans_id, db.RowsAffected)
	if len(tfollow)==1 && tfollow[0].user_id ==0 && tfollow[0].fans_id==0{
		return errors.New("Repeat to Cancel")
	}
	//if err := db.Exec("DELETE from t_follow  where  user_id = ? AND fans_id = ?", userId, userToId).Error; err != nil {
	//	return err
	//}
	//if db.RowsAffected == 0 {
	//	return errors.New("重复删除")
	//}
	if err := db.Exec("UPDATE t_user SET follow_count=follow_count-1 WHERE id = ?", userId).Error; err != nil {
		return err
	}
	if err := db.Exec("UPDATE t_user SET follower_count=follower_count-1 WHERE id = ?", userToId).Error; err != nil {
		return err
	}

	return nil
}

func  GetFollowListByUserId(userId int64) ([]UserInfo,error){
	var err error
	var results []UserInfo

	fmt.Println("userId:",userId)
	if err = db.Raw("SELECT u.id as id ,u.name as name , u.follow_count as follow_count,u.follower_count as follower_count,1 as is_follow FROM t_follow r, t_user u WHERE r.fans_id = ? AND r.user_id = u.id", userId).Scan(&results).Error; err != nil {

		return results,err
	}else{
		fmt.Println(results)
	}
	if len(results) == 0 || (results)[0].Id == 0 {
		return results,errors.New("List is empty")
	}
	return results,nil
}

func  GetFollowerListByUserId(userId int64) ([]UserInfo,error){
	var err error
	var results []UserInfo


	fmt.Println("userId:",userId)
	if err = db.Raw("SELECT DISTINCT u.id as id ,u.name as name , u.follow_count as follow_count,u.follower_count as follower_count,coalesce(r.BID,0) as is_follow From t_user as  u,(select A.AID as AID ,B.BID as BID from  (select fans_id as AID from t_follow  where user_id = ?) as A left join  (select a.user_id as BID from t_follow as a inner join t_follow as b on a.fans_id = ? and b.fans_id = ?) as B on A.AID = B.BID ) as r where  r.AID = u.id",userId,userId,userId).Scan(&results).Error; err != nil {
		fmt.Println("err",err )
		return results,err
	}else{
		fmt.Println(results)
	}
	if len(results) == 0 || (results)[0].Id == 0 {
		return results,errors.New("List is empty")
	}
	return results,nil
}

