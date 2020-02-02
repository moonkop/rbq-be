package controller

import (
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"rbq-be/model"
	"rbq-be/utils"
	"sort"
)

func getDrafts(context *gin.Context) {
	mongo := context.MustGet("mongodb").(*mgo.Database)
	var articles []model.ArticleInfo
	mongo.C("articles").Find(nil).All(&articles)
	if articles == nil {
		articles = make([]model.ArticleInfo, 0)
	}
	sort.Slice(articles, func(i, j int) bool {
		return articles[i].Created.After(articles[j].Created)
	})
	utils.Response(context, utils.ResponseCodeOk, "ok", gin.H{
		"list": articles,
	})
}
func getDraftsByTag(context *gin.Context) {
	mongo := context.MustGet("mongodb").(*mgo.Database)
	var articles []model.ArticleInfo
	tag := context.Param("tag")
	mongo.C("articles").Find(gin.H{
		"tags": gin.H{
			"$elemMatch": gin.H{
				"$eq": tag,
			},
		},
	}).All(&articles)
	if articles == nil {
		articles = make([]model.ArticleInfo, 0)
	}
	sort.Slice(articles, func(i, j int) bool {
		return articles[i].Created.After(articles[j].Created)
	})
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
