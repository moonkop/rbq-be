package main

import (
	"github.com/gin-gonic/gin"
	"rbq-be/config"
	"rbq-be/controller"
	"rbq-be/db"
	"rbq-be/middleware"
	"strconv"
)

func main() {

	_config := config.ReadConfig()
	engine := gin.Default()
	db.Connect()
	middleware.BindDb(engine)
	middleware.BindSession(engine)
	middleware.Cors(engine)
	controller.RunHttpServer(engine)

	engine.Run(":" + strconv.FormatInt(int64(_config.Port), 10))
}
