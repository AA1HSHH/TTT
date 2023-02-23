package controller

import "github.com/AA1HSHH/TTT/dal"

type Response struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg,omitempty"`
}

type Video struct {
	Id            int64  `json:"id,omitempty"`
	Author        User   `json:"author"`
	PlayUrl       string `json:"play_url" json:"play_url,omitempty"`
	CoverUrl      string `json:"cover_url,omitempty"`
	FavoriteCount int64  `json:"favorite_count,omitempty"`
	CommentCount  int64  `json:"comment_count,omitempty"`
	IsFavorite    bool   `json:"is_favorite,omitempty"`
}

type Comment struct {
	Id         int64  `json:"id,omitempty"`
	User       User   `json:"user"`
	Content    string `json:"content,omitempty"`
	CreateDate string `json:"create_date,omitempty"`
}

type User struct {
	Id            int64  `json:"id,omitempty"`
	Name          string `json:"name,omitempty"`
	FollowCount   int64  `json:"follow_count,omitempty"`
	FollowerCount int64  `json:"follower_count,omitempty"`
	IsFollow      bool   `json:"is_follow,omitempty"`
}

type MessageSendEvent struct {
	UserId     int64  `json:"user_id,omitempty"`
	ToUserId   int64  `json:"to_user_id,omitempty"`
	MsgContent string `json:"msg_content,omitempty"`
}

type MessagePushEvent struct {
	FromUserId int64  `json:"user_id,omitempty"`
	MsgContent string `json:"msg_content,omitempty"`
}

type APIVideo struct {
	Id            int64   `json:"id,omitempty"`
	Author        APIUser `json:"author"`
	PlayUrl       string  `json:"play_url" json:"play_url,omitempty"`
	CoverUrl      string  `json:"cover_url,omitempty"`
	FavoriteCount int64   `json:"favorite_count,omitempty"`
	CommentCount  int64   `json:"comment_count,omitempty"`
	IsFavorite    bool    `json:"is_favorite,omitempty"`
}
type APIUser struct {
	Id             int64  `json:"id,omitempty"`
	Name           string `json:"name,omitempty"`
	FollowCount    int64  `json:"follow_count,omitempty"`
	FollowerCount  int64  `json:"follower_count,omitempty"`
	IsFollow       bool   `json:"is_follow,omitempty"`
	Avatar         string `json:"avatar"`
	BackgroundImg  string `json:"background_image"`
	Signature      string `json:"signature"`
	TotalFavorited int64  `json:"total_favorited"`
	WorkCnt        int64  `json:"work_count"`
	FavoriteCnt    int64  `json:"favorite_count"`
}
type APIComment struct {
	Id         int64   `json:"id,omitempty"`
	User       APIUser `json:"user"`
	Content    string  `json:"content,omitempty"`
	CreateDate string  `json:"create_date,omitempty"`
}

func constructFavoriteVideoList(videos []dal.DBVideo, mp map[int64]dal.User, followList map[int64]struct{}) []APIVideo {
	apiVideos := make([]APIVideo, len(videos))
	for i := 0; i < len(videos); i++ {
		video := videos[i]
		user := mp[videos[i].AuthorId]
		_, ok := followList[videos[i].AuthorId]
		author := APIUser{Id: user.Id, Name: user.Name, FollowCount: user.FollowCount,
			FollowerCount: user.FollowerCount, IsFollow: ok, Avatar: user.Avatar, BackgroundImg: user.BackgroundImg,
			Signature: user.Signature, TotalFavorited: user.TotalFavorited, WorkCnt: user.WorkCnt, FavoriteCnt: user.FavoriteCnt}
		apiVideos[i] = APIVideo{Id: video.Id, Author: author, PlayUrl: video.PlayUrl, CoverUrl: video.CoverUrl,
			FavoriteCount: video.FavoriteCount, IsFavorite: true}
	}
	return apiVideos
}
func constrctCommentList(comments []dal.VideoComment, mp map[int64]dal.User, followList map[int64]struct{}) []APIComment {
	apiComments := make([]APIComment, len(comments))
	for i := 0; i < len(comments); i++ {
		rst := comments[i]
		user := mp[rst.WriterId]
		_, ok := followList[rst.WriterId]
		comment := APIComment{
			Id: rst.Id,
			User: APIUser{Id: user.Id, Name: user.Name, FollowCount: user.FollowCount,
				FollowerCount: user.FollowerCount, IsFollow: ok, Avatar: user.Avatar, BackgroundImg: user.BackgroundImg,
				Signature: user.Signature, TotalFavorited: user.TotalFavorited, WorkCnt: user.WorkCnt, FavoriteCnt: user.FavoriteCnt},
			Content:    rst.Content,
			CreateDate: rst.CreateTime.Format("01-02"),
		}
		apiComments[i] = comment
	}
	return apiComments
}
