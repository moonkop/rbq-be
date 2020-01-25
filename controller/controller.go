package controller

import (
	"../utils"
	"github.com/gin-gonic/gin"
	"io/ioutil"
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

	c.GET("/getFile", GetFile)
}
func GetFile(c *gin.Context) {
	data, err := ioutil.ReadFile("./1.txt")
	utils.Check(err)
	str := string(data)
	c.JSON(200, gin.H{
		"code": 200,
		"err":  "ok",
		"payload": gin.H{
			"content": str,
		},
	})
}
