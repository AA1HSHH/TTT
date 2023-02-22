package AA1HSHH

import (
	"fmt"
	"github.com/AA1HSHH/TTT/dal"
	"github.com/AA1HSHH/TTT/mw"
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

	id, err := dal.QueryUserIDbyName("789")

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(id)

}

func TestQuery(t *testing.T) {
	dal.Init()
	dal.QueryUserbyId(18)
}

func TestGetencryPwd(t *testing.T) {
	pwds := []string{"1234561", "1234562", "1234563", "1234564", "1234565"}
	for _, pwd := range pwds {
		fmt.Printf("%v:%v\n", pwd, mw.HashPassword(pwd))
	}
}
