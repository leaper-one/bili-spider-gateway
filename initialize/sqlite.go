package initialize

import (
	"log"

	"github.com/leaper-one/bili-spider-gateway/global"
	"github.com/leaper-one/bili-spider-gateway/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

/*
该文件用于初始化Sqlite数据库
*/

func InitSqlite() {
	db, err := gorm.Open(sqlite.Open("./Database/sqlite.db"), &gorm.Config{})
	if err != nil {
		log.Panicln(err)
	}
	db.AutoMigrate(&model.RankModel{})
	global.DB = db
}
