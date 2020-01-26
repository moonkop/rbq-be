package model

import "time"

type ArticleInfo struct {
	Name     string     `json:"name"`
	Content  string     `json:"content"`
	Modified time.Time  `json:"modified"`
	Created  time.Time  `json:"created"`
	Comments []Comments `json:"comments"`
}
