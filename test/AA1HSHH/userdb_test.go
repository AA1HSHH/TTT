package AA1HSHH

import (
	"fmt"
	"github.com/AA1HSHH/TTT/dal"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDB(t *testing.T) {
	dal.Init()
	_, err := dal.CreateUser("101112", "123")
	if err != nil {
		fmt.Println(err)
	}

	assert.True(t, dal.IsExist("789"))
	assert.False(t, dal.IsExist("123"))

	id, err := dal.QueryUserbyName("789")

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(id)

}
