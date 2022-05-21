package errorinit

import "errors"

/*
该文件用于定义项目用的错误Error
*/

var PostValueError = errors.New("提交参数有误")
var SecretError = errors.New("请求验证秘钥错误")
