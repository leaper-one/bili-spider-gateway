package model

/*
该文件为配置文件config.json的结构体
*/

type Config struct {
	Author       string `json:"author"`
	GlobalSecret string `json:"global_secret"`
	Mysql        struct {
		Host     string `json:"host"`
		Port     string `json:"port"`
		User     string `json:"user"`
		Password string `json:"password"`
		Dbname   string `json:"dbname"`
	}
	ServerExposePort string `json:"server_expose_port"`
}
