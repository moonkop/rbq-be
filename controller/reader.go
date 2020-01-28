package controller

import (
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"rbq-be/model"
	"rbq-be/utils"
)

func getArticles(context *gin.Context) {

}
func getDrafts(context *gin.Context) {
	mongo := context.MustGet("mongodb").(*mgo.Database)
	var articles []model.ArticleInfo
	mongo.C("articles").Find(nil).All(&articles)
	utils.Response(context, utils.ResponseCodeOk, "ok", gin.H{
		"list": articles,
	})
}
func getDraftById(context *gin.Context) {
	mongo := context.MustGet("mongodb").(*mgo.Database)
	id := context.Param("id")
	var article model.ArticleInfo
	mongo.C("articles").FindId(bson.ObjectIdHex(id)).One(&article)
	utils.Response(context, utils.ResponseCodeOk, "ok", article)
}
func getArticleById(context *gin.Context) {
	utils.Unimplemented(context)
}
func createArticle() {

}
func createDraft() {

}
