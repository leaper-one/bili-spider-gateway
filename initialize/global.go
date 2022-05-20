package initialize

import (
	"encoding/json"
	"ginForBH/global"
	"ginForBH/model"
	"io/ioutil"
	"log"
)

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

	global.Secret = conf.Globalsecret
	global.Dsn = setDsn(conf)
}

func setDsn(c model.Config) (dsn string) {
	dsn = c.Mysql.User + ":" + c.Mysql.Password + "@tcp(" + c.Mysql.Host + ":" + c.Mysql.Port + ")/" + c.Mysql.Dbname + "?charset=utf8mb4&parseTime=True&loc=Local"
	return
}
