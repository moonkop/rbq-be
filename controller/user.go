package controller

import (
	"../config"
	"github.com/gin-gonic/gin"
)

func login(context *gin.Context) {

}

func adminLogin(context *gin.Context) {
	config := config.GetConfig()
	context.BindJSON()
}
