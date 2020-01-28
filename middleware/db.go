package middleware

import (
	"github.com/gin-gonic/gin"
	"rbq-be/db"
)

func BindDb(engine *gin.Engine) {

	engine.Use(func(context *gin.Context) {
		session := db.Session.Clone()
		defer session.Clone()
		context.Set("mongodb", session.DB(db.Connection.Database))
	})
}
