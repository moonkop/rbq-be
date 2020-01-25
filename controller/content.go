package controller

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"os"
	"rbq-be/config"
	"rbq-be/utils"
	"time"
)

type Comments struct {
	Created time.Time `json:"created"`
	Content string    `json:"content"`
	Author  string    `json:"author"`
}
type ArticleInfo struct {
	Name     string     `json:"name"`
	Content  string     `json:"content"`
	Modified time.Time  `json:"modified"`
	Created  time.Time  `json:"created"`
	Comments []Comments `json:"comments"`
}

func getArticles(context *gin.Context) {
	session := sessions.Default(context)
	isAdmin := session.Get("isAdmin").(bool)
	fileInfo, err := ioutil.ReadDir(config.GetConfig().ArticleDir)
	utils.Check(err)
	retArr := utils.Map(fileInfo, func(i interface{}) interface{} {
		return ArticleInfo{Name: i.(os.FileInfo).Name()}
	})
	if isAdmin {
		Response(context, ResponseCodeOk, "ok", gin.H{
			"list": retArr,
		})
	} else {
		Response(context, ResponseCodeFail, "no auth", nil)
	}
}
func getDrafts(context *gin.Context) {
	Unimplemented(context)
}
func getArticleById(context *gin.Context) {
	Unimplemented(context)
}
func createArticle() {

}
func createDraft() {

}
