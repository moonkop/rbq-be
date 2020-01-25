package main

import (
	"./config"
	"./controller"
	"./middleware"
	"github.com/gin-gonic/gin"
)

func main() {
	config.ReadConfig()
	engine := gin.Default()
	controller.RunHttpServer(engine)
	middleware.BindSession(engine)
	engine.Run(":8080")
}
