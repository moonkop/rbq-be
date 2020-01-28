package controller

import (
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"os"
	"rbq-be/config"
	"rbq-be/model"
	"rbq-be/utils"
	"strings"
)

func getArticles(context *gin.Context) {

}
func getDrafts(context *gin.Context) {

	fileInfo, err := ioutil.ReadDir(config.GetConfig().DraftDir)
	utils.Check(err)
	retArr := utils.Map(fileInfo, func(i interface{}) interface{} {
		name := strings.Replace(i.(os.FileInfo).Name(), ".md", "", -1)

		return model.ArticleInfo{Name: name}
	})
	if utils.IsAdmin(context) {
		utils.Response(context, utils.ResponseCodeOk, "ok", gin.H{
			"list": retArr,
		})
	} else {
		utils.Response(context, utils.ResponseCodeFail, "no auth", nil)
	}
}
func getDraft(context *gin.Context) {
	name := context.Param("name")
	data, err := ioutil.ReadFile(config.GetConfig().DraftDir + name + ".md")
	utils.Check(err)
	utils.Response(context, utils.ResponseCodeOk, "ok", gin.H{
		"content": string(data),
	})
}
func getArticleById(context *gin.Context) {
	utils.Unimplemented(context)
}
func createArticle() {

}
func createDraft() {

}
