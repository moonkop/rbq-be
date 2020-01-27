package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Cors(engine *gin.Engine) {
	engine.Use(func(context *gin.Context) {
		method := context.Request.Method
		context.Header("Access-Control-Allow-Origin", "http://localhost:3000")
		context.Header("Access-Control-Allow-Credentials", "true")
		context.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, PATCH, DELETE")
		if method == "OPTIONS" {
			context.AbortWithStatus(http.StatusNoContent)

		}
	})
}
