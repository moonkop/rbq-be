package middleware

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

var num int64 = 10

func BindSession(engine *gin.Engine) *gin.Engine {
	store := cookie.NewStore([]byte("userName"))
	engine.Use(sessions.Sessions("goSession", store))
	return engine
}
