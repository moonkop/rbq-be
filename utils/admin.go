package utils

import (
	"github.com/gin-gonic/gin"
)

func IsAdmin(context *gin.Context) bool {
	isAdmin, exist := context.Get("isAdmin")
	if !exist || isAdmin.(bool) == false {
		return false
	} else {
		return true
	}
}
