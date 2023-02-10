package CC

import (
	"fmt"
	"testing"
	"time"

	"github.com/AA1HSHH/TTT/dal"
	service "github.com/AA1HSHH/TTT/service/video"
	// "github.com/stretchr/testify/assert"
)

func TestDB(t *testing.T) {
	dal.Init()
	dbVideos, err := dal.QueryVideosByTime(time.Now())
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(dbVideos[0].Id)
}

func TestFeedService(t *testing.T) {
	dal.Init()
	// feedVideoList := make([]dal.Video, 0, 1)
	feedVideoList, nextTime, err := service.QueryFeedVideoList(0, time.Now())
	fmt.Println(feedVideoList[0].Id)
	fmt.Println(feedVideoList[0].AuthorId)
	fmt.Println(feedVideoList[0].Title)
	fmt.Println(feedVideoList[0].PublishTime)
	fmt.Println(feedVideoList[0].PlayUrl)
	fmt.Println(feedVideoList[0].CoverUrl)
	fmt.Println(feedVideoList[0].FavoriteCount)
	fmt.Println(feedVideoList[0].CommentCount)
	fmt.Println(feedVideoList[0].IsFavorite)
	fmt.Println(nextTime)
	if err != nil {
		fmt.Println(feedVideoList[0].Id)
		fmt.Println(nextTime)
	}
}
