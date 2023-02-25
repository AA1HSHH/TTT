package AA1HSHH

import (
	"fmt"
	"github.com/AA1HSHH/TTT/dal"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSample(t *testing.T) {
	dal.Init()

	err := dal.DeleteFavoirte(1, 2)
	fmt.Println(err)
	err = dal.CreatFavoirte(1, 2)
	assert.Nil(t, err)

	videos, _, _ := dal.QueryFavoriteList(2)
	fmt.Println(videos)

	err = dal.CreatFavoirte(1, 2)
	assert.NotNil(t, err)

	dal.DeleteFavoirte(1, 2)

	videos, _, _ = dal.QueryFavoriteList(2)
	fmt.Println(videos)
	err = dal.DeleteFavoirte(1, 2)
	assert.NotNil(t, err)
}
