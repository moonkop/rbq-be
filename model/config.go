package model

type Config struct {
	AdminName       string `json:"admin_name"`
	Port            int16  `json:"port"`
	DraftDir        string `json:"draft_dir"`
	ArticleDir      string `json:"article_dir"`
	MongoContentUrl string `json:"mongo_content_url"`
}
