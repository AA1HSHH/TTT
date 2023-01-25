package AA1HSHH

import (
	"fmt"
	"github.com/AA1HSHH/TTT/mw"
	"testing"
)

func TestToken(t *testing.T) {

	mw.JwtInit("../../config/jwtkey/sample_key", "../../config/jwtkey/sample_key.pub")
	ss, err := mw.CreateToken(12, "AA1HSHH")
	if err != nil {
		fmt.Println("create error:", err)
	}
	id, user, err := mw.TokenStringGetUser(ss)
	if err != nil {
		fmt.Println("authen error:", err)
	}
	fmt.Println(id, user)

}
