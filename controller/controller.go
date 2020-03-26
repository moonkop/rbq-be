package controller

import (
	"github.com/gin-gonic/gin"
)

func RunHttpServer(c *gin.Engine) {
	user := c.Group("/user")
	{
		user.GET("/info", getUserInfo)
		user.POST("/login", login)
	}
	c.GET("/articles", getArticles)
	c.GET("/article/:id", getArticleById)
	c.GET("/tags/:tag", getArticlesByTag)
	c.GET("/ws", openWs)
	//	writer.GET("/Articles", getArticles)
	c.POST("/article/new", newArticle)
	c.PATCH("/article/:id", editArticle)
	c.DELETE("/article/:id", deleteArticle)
}
