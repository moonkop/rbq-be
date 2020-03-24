package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/mitchellh/mapstructure"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"net/http"
	"rbq-be/model"
	"rbq-be/utils"
	"time"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:   1024,
	WriteBufferSize:  1024,
	HandshakeTimeout: 5 * time.Second,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type WsRequestType string

var (
	WsRequestTypeStage     WsRequestType = "stage"
	WsRequestTypeSubmit    WsRequestType = "submit"
	WsRequestTypeRead      WsRequestType = "readAll"
	WsRequestTypeSubscribe WsRequestType = "subscribe"
)

type WsResponseType string

var (
	WsResponseTypeUpdate WsResponseType = "update"
)

type WsResponse struct {
	Id      int            `json:"id"`
	Type    WsResponseType `json:"type"`
	Payload gin.H
}
type WsRequest struct {
	Payload gin.H         `json:"payload"`
	Type    WsRequestType `json:"type"`
	Id      int           `json:"id"`
}
type WsSubmitRequest struct {
	Id   string `json:"id"`
	Data string `json:"data"`
}
type WsReadRequest struct {
	Id string `json:"id"`
}
type WsStageRequest struct {
	Id      string `json:"id"`
	Content string `json:"content"`
}
type WsSubscribeRequest struct {
	Ids []string `json:"ids"`
}
type WsUpdateResponse struct {
	Id      string            `json:"id"`
	article model.ArticleInfo `json:"article"`
}

type WsUser struct {
	Id                int
	conn              *websocket.Conn
	subscribeArticles []string
}

var users []*WsUser
var globalIndex int = 0

func openWs(context *gin.Context) {

	ws, err := upgrader.Upgrade(context.Writer, context.Request, nil)
	utils.Check(err)
	currentIndex := globalIndex + 1
	globalIndex = currentIndex
	fmt.Printf("ws 已连接%d", currentIndex)
	user := WsUser{
		Id:                currentIndex,
		conn:              ws,
		subscribeArticles: make([]string, 0),
	}
	users = append(users, &user)
	defer func() {
		err1 := recover()
		fmt.Printf("ws %d已断开 %s\n", user, err1)
		ws.Close()
	}()
	for {
		var msg WsRequest
		err := ws.ReadJSON(&msg)
		utils.Check(err)
		fmt.Printf("content:%s\n", msg)
		switch msg.Type {
		case WsRequestTypeSubmit:
			{
				var submitRequest WsSubmitRequest
				mapstructure.Decode(msg.Payload, &submitRequest)
			}
		case WsRequestTypeStage:
			{
				var stageRequest WsStageRequest
				mapstructure.Decode(msg.Payload, &stageRequest)
				mongo := context.MustGet("mongodb").(*mgo.Database)
				var article model.ArticleInfo
				mongo.C("articles").FindId(bson.ObjectIdHex(stageRequest.Id)).One(&article)
				article.Content = stageRequest.Content
				mongo.C("articles").UpdateId(bson.ObjectIdHex(stageRequest.Id), &article)
				ws.WriteJSON(gin.H{
					"id":   msg.Id,
					"code": 0,
				})
				emitArticleUpdate(article)
			}
		case WsRequestTypeRead:
			{
				var readRequest WsReadRequest
				mapstructure.Decode(msg.Payload, &readRequest)
			}
		case WsRequestTypeSubscribe:
			{
				var subscribeRequest WsSubscribeRequest
				mapstructure.Decode(msg.Payload, &subscribeRequest)
				user.subscribeArticles = subscribeRequest.Ids
			}
		}
	}
}
func emitArticleUpdate(info model.ArticleInfo) {
	for i := range users {
		for i2 := range users[i].subscribeArticles {
			if info.Id == bson.ObjectIdHex(users[i].subscribeArticles[i2]) {
				users[i].conn.WriteJSON(WsResponse{
					Payload: utils.StructToMap(info),
					Type:    WsResponseTypeUpdate,
				})
			}
		}
	}
}

func newArticle(context *gin.Context) {
	mongo := context.MustGet("mongodb").(*mgo.Database)
	var data model.ArticleInfo
	context.BindJSON(&data)
	data.Id = bson.NewObjectId()
	data.Created = time.Now()
	mongo.C("articles").Insert(&data)
	utils.Response(context, utils.ResponseCodeOk, "ok", nil)
}
func editArticle(context *gin.Context) {
	mongo := context.MustGet("mongodb").(*mgo.Database)
	id := context.Param("id")
	var data model.ArticleInfo
	context.BindJSON(&data)
	if id == "" {
		utils.Response(context, utils.ResponseCodeFail, "Id should not be empty", nil)
		return
	}
	data.Modified = time.Now()
	mongo.C("articles").UpdateId(bson.ObjectIdHex(id), &data)
	utils.Response(context, utils.ResponseCodeOk, "ok", nil)
}
func deleteArticle(context *gin.Context) {
	mongo := context.MustGet("mongodb").(*mgo.Database)
	id := context.Param("id")
	mongo.C("articles").RemoveId(bson.ObjectIdHex(id))
	utils.Response(context, utils.ResponseCodeOk, "ok", nil)
}
