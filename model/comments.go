package model

import "time"

type Comments struct {
	Created time.Time `json:"created"`
	Content string    `json:"content"`
	Author  string    `json:"author"`
}
