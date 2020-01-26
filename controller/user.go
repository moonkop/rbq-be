package controller

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"rbq-be/config"
	"rbq-be/utils"
)

func login(context *gin.Context) {
	utils.Unimplemented(context)
}

func adminLogin(context *gin.Context) {
	config := config.GetConfig()
	var a struct {
		Password string `json:"password"`
	}
	context.BindJSON(&a)
	if config.AdminPassword == a.Password {
		session := sessions.Default(context)
		session.Set("isAdmin", true)
		session.Save()
		utils.Response(context, utils.ResponseCodeOk, "success", nil)
	} else {
		session := sessions.Default(context)
		session.Set("isAdmin", false)
		session.Save()
		utils.Response(context, utils.ResponseCodeFail, "fail", nil)
	}
}
