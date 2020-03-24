package controller

import (
	"github.com/gin-gonic/gin"
	"rbq-be/middleware"
)

func RunHttpServer(c *gin.Engine) {
	user := c.Group("/user")
	{
		user.GET("/info", getUserInfo)
		user.POST("/login", login)
	}
	content := c.Group("/reader")
	{
		content.GET("/articles", getArticles)
		content.GET("/article/:id", getArticleById)
		content.GET("/tags/:tag", getArticlesByTag)

	}
	writer := c.Group("/writer")
	writer.Use(middleware.AdminAuth)
	{
		writer.GET("/ws", openWs)
		//	writer.GET("/Articles", getArticles)
		writer.POST("/article/new", newArticle)
		writer.PATCH("/article/:id", editArticle)
		writer.DELETE("/article/:id", deleteArticle)
		//writer.GET("/Article/:name", getArticle)
	}
}
