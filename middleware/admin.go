package middleware

import (
	"github.com/gin-gonic/gin"
	"rbq-be/utils"
)

func AdminAuth(context *gin.Context) {
	if utils.IsAdmin(context) {
		context.Next()
	} else {
		utils.Response(context, utils.ResponseCodeFail, "no auth", nil)
		context.Abort()
	}
}
