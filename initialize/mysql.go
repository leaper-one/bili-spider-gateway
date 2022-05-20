package initialize

import (
	"ginForBH/global"
	"ginForBH/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

func InitMysql() {
	dsn := global.Dsn
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}
	db.AutoMigrate(&model.RankModel{})
	global.DB = db
}
