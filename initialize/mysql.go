package initialize

import (
	"ginForBH/global"
	"ginForBH/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
)

/*
该文件用于初始化Mysql数据库
*/

func InitMysql() {
	dsn := global.Dsn
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{Logger: logger.Default.LogMode(logger.Info)})
	if err != nil {
		log.Fatalln(err)
	}
	db.AutoMigrate(&model.RankModel{})
	global.DB = db
}
