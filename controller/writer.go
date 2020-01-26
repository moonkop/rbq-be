package controller

import (
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"os"
	"path"
	"rbq-be/config"
	"rbq-be/model"
	"rbq-be/utils"
)

func openWs(context *gin.Context) {
	utils.Unimplemented(context)

}

func newDraft(context *gin.Context) {
	var data model.ArticleInfo
	context.BindJSON(&data)
	if data.Name == "" {
		utils.Response(context, utils.ResponseCodeFail, "name should not be empty", nil)
		return
	}
	name := path.Join(config.GetConfig().DraftDir, data.Name+".md")
	ioutil.WriteFile(name, []byte(data.Content), 0644)
	utils.Response(context, utils.ResponseCodeOk, "ok", nil)
}
func editDraft(context *gin.Context) {
	var data model.ArticleInfo
	context.BindJSON(&data)
	if data.Name == "" {
		utils.Response(context, utils.ResponseCodeFail, "name should not be empty", nil)
		return
	}
	name := path.Join(config.GetConfig().DraftDir, data.Name+".md")
	ioutil.WriteFile(name, []byte(data.Content), 0644)
	utils.Response(context, utils.ResponseCodeOk, "ok", nil)
}
func deleteDraft(context *gin.Context) {
	name := context.Param("name")
	os.Remove(path.Join(config.GetConfig().DraftDir, name+".md"))
	utils.Response(context, utils.ResponseCodeOk, "ok", nil)
}
