package controller

import (
	"fmt"
	"github.com/AA1HSHH/TTT/dal"
	"github.com/AA1HSHH/TTT/mw"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type UserListResponse struct {
	Response
	UserList []dal.UserInfo `json:"user_list"`
}
type FriendListResponse struct {
	Response
	UserList []dal.FriendUser `json:"user_list"`
}

type ProxyPostFollowAction struct {
	*gin.Context
	token      string
	followId   int64
	actionType int32
}

type ListAction struct {
	*gin.Context
	userId int64
	token  string
}

func RelationAction(c *gin.Context) {
	var actionType, FOLLOW, CANCEL int32
	var ac int64
	FOLLOW = 1
	CANCEL = 2
	//p:= &ProxyPostFollowAction{Context: c} //对吗？ * &*/

	token := c.Query("token")
	userId, _, err := mw.TokenStringGetUser(token)
	FId := c.Query("to_user_id")
	followId, err := strconv.ParseInt(FId, 10, 64) //转换成数字
	if err != nil {
		fmt.Println(err)
		return
	}
	AType := c.Query("action_type")
	ac, err = strconv.ParseInt(AType, 10, 64)
	actionType = int32(ac)
	if err != nil {
		fmt.Println(err)
		return
	}

	//可以提取示例，下面的由于数据库的问题还没有
	//checkNum
	if exist := dal.IsUserExistById(followId); !exist {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "Follow User is not exist"})
	} else if actionType != FOLLOW && actionType != CANCEL {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "Follow Action is not exist"})
	} else if userId == followId {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "Follow yourself is not allowed "})
	} else {
		if actionType == FOLLOW {
			err = dal.AddUserFollow(userId, followId)
			if err != nil {
				c.JSON(http.StatusOK, Response{StatusCode: 0, StatusMsg: "Fail to follow,Do not repeat follow!"})
			} else {
				c.JSON(http.StatusOK, Response{StatusCode: 0, StatusMsg: "Successful to follow"})
			}
			// UpdateVideoFavorState redis 更新点赞状态，state:true为点赞，false为取消点赞,我们没有用redis ,所以暂不考虑

		}
		if actionType == CANCEL {
			err = dal.CancelUserFollow(userId, followId)
			if err != nil {
				fmt.Println(err)
				c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: err.Error()})
			} else {
				c.JSON(http.StatusOK, Response{StatusCode: 0, StatusMsg: "Successful to cancel"})
			}

		}
	}

}

func JudgeUserFair(userId int64, token string) string {
	if exist := dal.IsUserExistById(userId); !exist {
		//c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "User is not exist"})
		return "User is not exist"
	}

	id, username, err := mw.TokenStringGetUser(token)
	if err != nil {
		//c.JSON(http.StatusOK, UserResponse{
		//	Response: Response{StatusCode: 1, StatusMsg: "Authen failed"},
		//})
		return "Authen failed"
	}
	fmt.Println("获取的信息为：id - username : ",id ,username)
	return "ok"
	//if exit := dal.IsExist(username); !exit {
	//	c.JSON(http.StatusOK, UserResponse{
	//		Response: Response{StatusCode: 1, StatusMsg: "User doesn't exist"},
	//	})
	//	return
	//}
}

// FollowList all users have same follow list
func FollowList(c *gin.Context) {

	Id := c.Query("user_id")
	token := c.Query("token")
	userId, err := strconv.ParseInt(Id, 10, 64) //转换成数字
	if err != nil {
		fmt.Println(err)
		return
	}

	information := JudgeUserFair(userId, token )
	if information != "ok" {
		c.JSON(http.StatusOK, UserListResponse{
			Response: Response{
				StatusCode: 1,
				StatusMsg:  information,
			},
			UserList: nil ,
		})
		return
	}

	//var userList []dal.UserInfo // 这里的userList 表示 关注列表
	var userInfo []dal.UserInfo
	//var userSet dal.UserInfo

	userInfo, err = dal.GetFollowListByUserId(userId)
	if err != nil {
		print("没有关注关系")
		//c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "?"})

	}

	//for i := 0; i < len(userInfo); i++ {
	//	userSet.Id = userInfo[i].Id
	//	userSet.FollowCount = userInfo[i].FollowCount
	//	userSet.FollowerCount = userInfo[i].FollowerCount
	//	userSet.IsFollow = userInfo[i].IsFollow
	//	userSet.Name = userInfo[i].Name
	//	userList = append(userList,userSet )
	//
	//}

	c.JSON(http.StatusOK, UserListResponse{
		Response: Response{
			StatusCode: 0,
			StatusMsg:"Successful to get Follow List.",
		},
		UserList: userInfo,
	})
}

// FollowerList all users have same follower list
func FollowerList(c *gin.Context) {

	Id := c.Query("user_id")
	token := c.Query("token")
	userId, err := strconv.ParseInt(Id, 10, 64) //转换成数字
	if err != nil {
		fmt.Println(err)
		return
	}

	information := JudgeUserFair(userId, token )
	if information != "ok" {
		c.JSON(http.StatusOK, UserListResponse{
			Response: Response{
				StatusCode: 1,
				StatusMsg:  information,
			},
			UserList: nil ,
		})
		return
	}

	//var userList []dal.UserInfo // 这里的userList 表示 粉丝列表
	var userInfo []dal.UserInfo
	//var userSet dal.UserInfo

	userInfo, err = dal.GetFollowerListByUserId(userId)
	if err != nil {
		fmt.Println(err)
		//c.JSON(http.StatusOK,  Response{StatusCode: 1, StatusMsg: "?"})
	}

	//for i := 0; i < len(userInfo); i++ {
	//	userSet.Id = userInfo[i].Id
	//	userSet.FollowCount = userInfo[i].FollowCount
	//	userSet.FollowerCount = userInfo[i].FollowerCount
	//	userSet.IsFollow=userInfo[i].IsFollow
	//	userSet.Name = userInfo[i].Name
	//	userList = append(userList,userSet )
	//
	//}
	//if len(userInfo)>=2 {
	//	fmt.Println("扣扣扣扣")
	//
	//	userInfo[2].IsFollow = false }

	c.JSON(http.StatusOK, UserListResponse{
		Response: Response{
			StatusCode: 0,
			StatusMsg:"Successful to get Follower List.",
		},
		UserList: userInfo,
	})
}

// FriendList all users have same friend list DemoUser 不行
func FriendList(c *gin.Context) {
	Id := c.Query("user_id")
	token := c.Query("token")
	userId, err := strconv.ParseInt(Id, 10, 64) //转换成数字
	if err != nil {
		fmt.Println(err)
		return
	}

	information := JudgeUserFair(userId, token )
	if information != "ok" {
		c.JSON(http.StatusOK, UserListResponse{
			Response: Response{
				StatusCode: 1,
				StatusMsg:  information,
			},
			UserList: nil ,
		})
		return
	}

	//var userList []dal.UserInfo // 这里的userList 表示 粉丝列表
	var userInfo []dal.FriendUser
	//var userSet dal.UserInfo

	userInfo, err = dal.GetChat(userId)
	if err != nil {
		fmt.Println("发生了错误",err)
		//c.JSON(http.StatusOK,  Response{StatusCode: 1, StatusMsg: "?"})
	}
	c.JSON(http.StatusOK, FriendListResponse{
		Response: Response{
			StatusCode: 0,
			StatusMsg:"Successful to get Friend List.",
		},
		UserList: userInfo,
	})

}
