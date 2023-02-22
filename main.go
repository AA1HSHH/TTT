package main

import (
	"github.com/AA1HSHH/TTT/config"
	"github.com/AA1HSHH/TTT/dal"
	"github.com/AA1HSHH/TTT/mw"
	"github.com/gin-gonic/gin"
)

func main() {

	if err := dal.Init(); err != nil {
		panic("db init error")
	}
	if err := mw.JwtInit(config.JwtPrivateKey, config.JwtPublicKey); err != nil {
		panic("JWT init error")
	}

	r := gin.Default()

	initRouter(r)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
