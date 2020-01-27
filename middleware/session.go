package middleware

import (
	//	"github.com/gin-contrib/sessions"
	//"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

var num int64 = 10

func BindSession(engine *gin.Engine) *gin.Engine {
	store := cookie.NewStore([]byte("userName"))
	engine.Use(sessions.Sessions("goSession", store))
	engine.Use(func(context *gin.Context) {
		session := sessions.Default(context)
		isAdmin := session.Get("isAdmin")
		context.Set("isAdmin", isAdmin != nil && isAdmin.(bool) == true)
		session.Save()
	})
	return engine
}
