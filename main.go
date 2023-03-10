package main

import (
	"github.com/Joeyzsy/douyin-app-demo/dal/db"
	"github.com/Joeyzsy/douyin-app-demo/service"
	"github.com/gin-gonic/gin"
)

func main() {
	go service.RunMessageServer()

	r := gin.Default()

	initRouter(r)
	db.Init()

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
