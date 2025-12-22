package user

import (
	"shc-ai-demo/common/code"
	shc_email "shc-ai-demo/common/email"
	shc_redis "shc-ai-demo/common/redis"
	"shc-ai-demo/dao/user"
	"shc-ai-demo/utils"
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

func SendCaptcha(email_ string) code.Code {
	//1、生成验证码
	send_code := utils.GetRandomNumbers(6)

	//2、存入redis
	if err := shc_redis.SetCaptchaForEmail(email_, send_code); err != nil {
		return code.CodeServerBusy
	}

	//3、发送验证码到用户邮箱
	if err := shc_email.SendCaptcha(email_, send_code, shc_email.CodeMsg); err != nil {
		return code.CodeServerBusy
	}

	return code.CodeSuccess
}
