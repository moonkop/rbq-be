package middleware

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"rbq-be/utils"
)

func AdminAuth(context *gin.Context) {

	session := sessions.Default(context)
	isAdmin := session.Get("isAdmin").(bool)
	if isAdmin {
		context.Next()
	} else {
		utils.Response(context, utils.ResponseCodeFail, "no auth", nil)
		context.Abort()
	}
}
