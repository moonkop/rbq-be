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

type MessageType string

var (
	MessageTypeStage  MessageType = "stage"
	MessageTypeSubmit MessageType = "submit"
	MessageTypeRead   MessageType = "readAll"
)

type Message struct {
	Payload map[string]interface{} `json:"payload"`
	Type    MessageType            `json:"type"`
	Id      int                    `json:"id"`
}
type SubmitRequest struct {
	Id   int    `json:"id"`
	Data string `json:"data"`
}
type ReadRequest struct {
	Id int `json:"id"`
}
type StageRequest struct {
	Id      int    `json:"id"`
	Content string `json:"content"`
}

func openWs(context *gin.Context) {

	ws, err := upgrader.Upgrade(context.Writer, context.Request, nil)
	utils.Check(err)
	//	defer ws.Close()
	for {
		var msg Message
		err := ws.ReadJSON(&msg)
		utils.Check(err)
		fmt.Printf("content:%s", msg)
		switch msg.Type {
		case MessageTypeSubmit:
			{
				var submitRequest SubmitRequest
				mapstructure.Decode(msg.Payload, &submitRequest)
			}
		case MessageTypeStage:
			{
				var stageRequest StageRequest
				mapstructure.Decode(msg.Payload, &stageRequest)
				mongo := context.MustGet("mongodb").(*mgo.Database)
				id := context.Param("id")
				var article model.ArticleInfo
				mongo.C("articles").FindId(bson.ObjectIdHex(id)).One(&article)
				article.Content = stageRequest.Content
				mongo.C("articles").UpdateId(bson.ObjectIdHex(id), &article)
				ws.WriteJSON(gin.H{
					"code": 0,
				})
			}
		case MessageTypeRead:
			{
				var readRequest ReadRequest
				mapstructure.Decode(msg.Payload, &readRequest)
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
