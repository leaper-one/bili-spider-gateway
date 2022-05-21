package global

import "gorm.io/gorm"

/*
该文件用于定义项目的全局变量
*/

//数据库相关
var DB *gorm.DB
var Dsn string

//验证密钥
var Secret string

//api暴露端口
var ExposePort = "8000"

//用于排名的rank基础值
var Rankmax = int(10e8)
var Rankmin = 0
var Rankadd = int(10e6)
