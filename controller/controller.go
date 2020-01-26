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
	admin.Use(middleware.AdminAuth)
	{
		admin.POST("/draft/new", newDraft)
		admin.PATCH("/article/new", editDraft)
		admin.DELETE("/draft/:name", deleteDraft)
	}
}
