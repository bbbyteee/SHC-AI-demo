package user

import (
	"shc-ai-demo/common/code"
	"shc-ai-demo/dao/user"
)

func Register(email, password, captcha string) (string, code.Code) {

	// var ok bool
	// var userInfromation *model.User

	//1、先判断用户是否存在
	if ok, _ := user.IsExistUser(email); ok {
		return "用户已存在", code.CodeUserExist
	}

	//2、从redis中获取验证码进行比对

	return "", code.CodeSuccess
}
