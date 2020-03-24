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
	MessageTypeStage     MessageType = "stage"
	MessageTypeSubmit    MessageType = "submit"
	MessageTypeRead      MessageType = "readAll"
	MessageTypeSubscribe MessageType = "subscribe"
)

type Message struct {
	Payload map[string]interface{} `json:"payload"`
	Type    MessageType            `json:"type"`
	Id      int                    `json:"id"`
}
type SubmitRequest struct {
	Id   string `json:"id"`
	Data string `json:"data"`
}
type ReadRequest struct {
	Id string `json:"id"`
}
type StageRequest struct {
	Id      string `json:"id"`
	Content string `json:"content"`
}
type SubscribeRequest struct {
	Ids []string `json:"ids"`
}
type UpdateResponse struct {
	Id      string `json:"id"`
	Content string `json:"content"`
}
type WsUsers struct {
	Id                int
	conn              *websocket.Conn
	subscribeArticles []string
}

var conns []WsUsers
var globalIndex int = 0

func openWs(context *gin.Context) {

	ws, err := upgrader.Upgrade(context.Writer, context.Request, nil)
	utils.Check(err)
	currentIndex := globalIndex + 1
	globalIndex = currentIndex
	fmt.Printf("ws 已连接%d", currentIndex)
	user := WsUsers{
		Id:                currentIndex,
		conn:              ws,
		subscribeArticles: make([]string, 0),
	}
	conns = append(conns, user)
	defer func() {
		err1 := recover()
		fmt.Printf("ws %d已断开 %s\n", user, err1)
		ws.Close()
	}()
	for {
		var msg Message
		err := ws.ReadJSON(&msg)
		utils.Check(err)
		err.Error()

		fmt.Printf("content:%s\n", msg)
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
				var article model.ArticleInfo
				mongo.C("articles").FindId(bson.ObjectIdHex(stageRequest.Id)).One(&article)
				article.Content = stageRequest.Content
				mongo.C("articles").UpdateId(bson.ObjectIdHex(stageRequest.Id), &article)
				ws.WriteJSON(gin.H{
					"id":   msg.Id,
					"code": 0,
				})
			}
		case MessageTypeRead:
			{
				var readRequest ReadRequest
				mapstructure.Decode(msg.Payload, &readRequest)
			}
		case MessageTypeSubscribe:

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
