package model

type Config struct {
	AdminPassword string `json:"admin_password"`
	Port          int16  `json:"port"`
	DraftDir      string `json:"draft_dir"`
	ArticleDir    string `json:"article_dir"`
}
