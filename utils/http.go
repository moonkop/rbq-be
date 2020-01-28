package utils

import "github.com/gin-gonic/gin"

func Unimplemented(context *gin.Context) {
	context.JSON(200, gin.H{
		"code":    404,
		"message": "un implemented",
	})
}

type ResponseCode int16

const (
	ResponseCodeOk   ResponseCode = 200
	ResponseCodeFail ResponseCode = 500
)

func Response(context *gin.Context, code ResponseCode, message string, payload interface{}) {

	response := gin.H{
		"code":    code,
		"message": message,
	}
	if payload != nil {
		response["payload"] = payload
	}
	context.JSON(200, response)
}
