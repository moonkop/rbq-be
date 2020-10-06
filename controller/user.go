package controller

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"rbq-be/config"
	"rbq-be/utils"
)

func login(context *gin.Context) {
	var a struct {
		Name string `json:"name"`
	}
	context.BindJSON(&a)
	if a.Name == "" {
		utils.Response(context, utils.ResponseCodeFail, "invalid name", nil)
		return
	}
	session := sessions.Default(context)
	session.Set("name", a.Name)
	isAdmin := config.GetConfig().AdminName == a.Name
	session.Set("isAdmin", isAdmin)
	session.Save()
	utils.Response(context, utils.ResponseCodeOk, "login as admin", gin.H{
		"name":    a.Name,
		"isAdmin": isAdmin,
	})
}
func logout(context *gin.Context) {
	sess := sessions.Default(context)
	sess.Clear()
	utils.Response(context, utils.ResponseCodeOk, "logout successful", gin.H{})
}
func getUserInfo(context *gin.Context) {
	session := sessions.Default(context)
	name := session.Get("name")
	isAdmin := session.Get("isAdmin")
	utils.Response(context, utils.ResponseCodeOk, "success", gin.H{
		"name":    name,
		"isAdmin": isAdmin,
	})

}
