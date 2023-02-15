package main

import (
	"github.com/AA1HSHH/TTT/config"
	"github.com/AA1HSHH/TTT/dal"
	"github.com/AA1HSHH/TTT/mw"
	"github.com/AA1HSHH/TTT/service"
	"github.com/gin-gonic/gin"
)

func main() {
	//go service.RunMessageServer()
	//
	if err := dal.Init(); err != nil {
		panic("db init error")
	}
	if err := mw.JwtInit(config.JwtPrivateKey, config.JwtPublicKey); err != nil {
		panic("JWT init error")
	}

	go service.Manager.Start()
	r := gin.Default()
	r.Use(gin.Recovery(), gin.Logger())
	initRouter(r)
	_ = r.Run(dal.HttpPort)
	//

	//r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

}
