package initialize

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/leaper-one/bili-spider-gateway/global"
	"github.com/leaper-one/bili-spider-gateway/model"
)

/*
该文件为用于初始化全局变量的函数
*/

func InitGlobal() {
	var conf model.Config

	jsonFile, err := ioutil.ReadFile("./config.json")
	if err != nil {
		log.Printf(err.Error())
	}

	err = json.Unmarshal(jsonFile, &conf)
	if err != nil {
		log.Printf(err.Error())
	}

	global.Secret = conf.GlobalSecret
	global.Dsn = setDsn(conf)
	global.ExposePort = conf.ServerExposePort
}

func setDsn(c model.Config) (dsn string) {
	dsn = c.Mysql.User + ":" + c.Mysql.Password + "@tcp(" + c.Mysql.Host + ":" + c.Mysql.Port + ")/" + c.Mysql.Dbname + "?charset=utf8mb4&parseTime=True&loc=Local"
	return
}
