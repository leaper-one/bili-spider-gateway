package model

type Config struct {
	Author       string `json:"author"`
	Globalsecret string `json:"globalsecret"`
	Mysql        struct {
		Host     string `json:"host"`
		Port     string `json:"port"`
		User     string `json:"user"`
		Password string `json:"password"`
		Dbname   string `json:"dbname"`
	}
}
