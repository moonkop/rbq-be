package controller

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func RunHttpServer(c *gin.Engine) {
	user := c.Group("/user")
	{
		user.Any("/adminLogin", adminLogin)
		user.Any("/login", login)
	}
	writing := c.Group("/writing")
	{
		writing.GET("/ws", openWs)
	}
	content := c.Group("/content")
	{
		content.GET("/articles", getArticles)
		content.GET("/drafts", getDrafts)
		content.GET("/article/:id", getArticleById)
	}
	admin := c.Group("/admin")
	admin.Use(AdminAuth)
	{
		admin.POST("/draft/new")
		admin.POST("/article/new")
	}
}

func newDraft(context *gin.Context) {

}
func AdminAuth(context *gin.Context) {

	session := sessions.Default(context)
	isAdmin := session.Get("isAdmin").(bool)
	if isAdmin {
		context.Next()
	} else {
		Response(context, ResponseCodeFail, "no auth", nil)
		context.Abort()
	}
}

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

func Response(context *gin.Context, code ResponseCode, message string, payload gin.H) {

	response := gin.H{
		"code":    code,
		"message": message,
	}
	if payload != nil {
		response["payload"] = payload
	}
	context.JSON(200, response)
}
