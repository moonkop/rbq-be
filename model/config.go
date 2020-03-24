package model

type Config struct {
	AdminName       string `json:"admin_name"`
	Port            int16  `json:"port"`
	ArticleDir      string `json:"Article_dir"`
	MongoContentUrl string `json:"mongo_content_url"`
}
