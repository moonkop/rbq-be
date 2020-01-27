package controller

import (
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"os"
	"rbq-be/config"
	"rbq-be/model"
	"rbq-be/utils"
)

func getArticles(context *gin.Context) {

}
func getDrafts(context *gin.Context) {

	fileInfo, err := ioutil.ReadDir(config.GetConfig().DraftDir)
	utils.Check(err)
	retArr := utils.Map(fileInfo, func(i interface{}) interface{} {
		return model.ArticleInfo{Name: i.(os.FileInfo).Name()}
	})
	if utils.IsAdmin(context) {
		utils.Response(context, utils.ResponseCodeOk, "ok", gin.H{
			"list": retArr,
		})
	} else {
		utils.Response(context, utils.ResponseCodeFail, "no auth", nil)
	}
}
func getArticleById(context *gin.Context) {
	utils.Unimplemented(context)
}
func createArticle() {

}
func createDraft() {

}
