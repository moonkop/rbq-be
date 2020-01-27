package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"rbq-be/config"
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
