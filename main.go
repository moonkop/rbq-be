package main

import (
	"github.com/gin-gonic/gin"
	"rbq-be/config"
	"rbq-be/controller"
	"rbq-be/middleware"
	"strconv"
)

func main() {
	config := config.ReadConfig()
	engine := gin.Default()
	middleware.BindSession(engine)
	controller.RunHttpServer(engine)
	engine.Run(":" + strconv.FormatInt(int64(config.Port), 10))
}
