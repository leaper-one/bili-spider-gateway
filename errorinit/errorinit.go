package errorinit

import "errors"

var PostValueError = errors.New("提交参数有误")
var SecretError = errors.New("请求验证秘钥错误")
