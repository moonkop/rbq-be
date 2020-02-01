package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
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
	MessageTypeStage   = "stage"
	MessageTypeSubmit  = "submit"
	MessageTypeReadAll = "readAll"
)

type Message struct {
	Content string      `json:"content"`
	Type    MessageType `json:"type"`
	Id      int         `json:"id"`
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
	}
}

func newDraft(context *gin.Context) {
	mongo := context.MustGet("mongodb").(*mgo.Database)
	var data model.ArticleInfo
	context.BindJSON(&data)
	data.Id = bson.NewObjectId()
	mongo.C("articles").Insert(&data)
	utils.Response(context, utils.ResponseCodeOk, "ok", nil)
}
func editDraft(context *gin.Context) {
	mongo := context.MustGet("mongodb").(*mgo.Database)
	id := context.Param("id")
	var data model.ArticleInfo
	context.BindJSON(&data)
	if id == "" {
		utils.Response(context, utils.ResponseCodeFail, "Id should not be empty", nil)
		return
	}
	mongo.C("articles").UpdateId(bson.ObjectIdHex(id), &data)
	utils.Response(context, utils.ResponseCodeOk, "ok", nil)
}
func deleteDraft(context *gin.Context) {
	mongo := context.MustGet("mongodb").(*mgo.Database)
	id := context.Param("id")
	mongo.C("articles").RemoveId(bson.ObjectIdHex(id))
	utils.Response(context, utils.ResponseCodeOk, "ok", nil)
}
