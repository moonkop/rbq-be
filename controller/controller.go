package controller

import (
	"github.com/gin-gonic/gin"
	"rbq-be/middleware"
)

func RunHttpServer(c *gin.Engine) {
	user := c.Group("/user")
	{
		user.POST("/adminLogin", adminLogin)
		user.POST("/login", login)
	}
	content := c.Group("/reader")
	{
		content.GET("/articles", getDrafts)
		content.GET("/article/:id", getArticleById)
	}
	writer := c.Group("/writer")
	writer.Use(middleware.AdminAuth)
	{
		writer.GET("/ws", openWs)
		//	writer.GET("/drafts", getDrafts)
		writer.GET("/drafts", getDrafts)
		writer.POST("/draft/new", newDraft)
		writer.PATCH("/draft/:id", editDraft)
		writer.DELETE("/draft/:id", deleteDraft)
		//writer.GET("/draft/:name", getDraft)
		writer.GET("/draft/:id", getDraftById)
	}
}
