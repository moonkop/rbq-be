package model

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type ArticleInfo struct {
	Id       bson.ObjectId `json:"id" bson:"_id"`
	Name     string        `json:"name,omitempty"`
	Content  string        `json:"content,omitempty"`
	Modified time.Time     `json:"modified,omitempty"`
	Created  time.Time     `json:"created,omitempty"`
	Comments []Comments    `json:"comments,omitempty"`
}
