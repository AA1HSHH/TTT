package AA1HSHH

import (
	"fmt"
	"github.com/AA1HSHH/TTT/dal"
	"testing"
)

func TestSample(t *testing.T) {
	dal.Init()
	dal.CreatFavoirte(1, 2)
	videos, _, _ := dal.QueryFavoriteList(2)
	fmt.Println(videos)
	dal.DeleteFavoirte(1, 2)
	videos, _, _ = dal.QueryFavoriteList(2)
	fmt.Println(videos)
}
