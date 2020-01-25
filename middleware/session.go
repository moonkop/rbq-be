package middleware

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"strconv"
)

var num int64 = 10

func BindSession(engine *gin.Engine) *gin.Engine {
	store := cookie.NewStore([]byte("userName"))
	engine.Use(sessions.Sessions("goSession", store))
	engine.GET("/getSession", func(context *gin.Context) {
		session := sessions.Default(context)
		testKey := session.Get("testKey")
		num++
		context.JSON(200, gin.H{
			"code":    200,
			"content": testKey,
		})
	})
	engine.GET("/setSession", func(context *gin.Context) {
		session := sessions.Default(context)
		num++
		str := "smjibadongxi" + strconv.FormatInt(num, 10)
		session.Set("testKey", str)
		session.Save()
		context.JSON(200, gin.H{
			"code":    200,
			"content": str,
		})
	})
	return engine

}
