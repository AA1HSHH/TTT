package AA1HSHH

import (
	"fmt"
	"github.com/AA1HSHH/TTT/dal"
	"testing"
	"time"
)

func TestNegativeNumber(t *testing.T) {
	dal.Init()
	videos, _, _ := dal.QueryVideoCommentList(1)
	fmt.Println(videos)
	dal.CreateVideoComment(1, 1, "good1")
	time.Sleep(2 * time.Second)
	dal.CreateVideoComment(1, 2, "good2")
	time.Sleep(2 * time.Second)
	dal.CreateVideoComment(1, 3, "good3")
	videos, _, _ = dal.QueryVideoCommentList(1)
	fmt.Println(videos)
}
