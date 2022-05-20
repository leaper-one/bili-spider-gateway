package initialize

import (
	"ginForBH/global"
	"ginForBH/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

func InitMysql() {
	dsn := "root:cunky812819@tcp(cloud.cunoe.com:3306)/rankv1?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}
	db.AutoMigrate(&model.RankModel{})
	global.DB = db
}
