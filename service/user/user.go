package user

import (
	"shc-ai-demo/common/code"
	shc_email "shc-ai-demo/common/email"
	shc_redis "shc-ai-demo/common/redis"
	"shc-ai-demo/dao/user"
	"shc-ai-demo/model"
	"shc-ai-demo/utils"
	"shc-ai-demo/utils/myjwt"
)

func Login(username, password string) (string, code.Code) {
	var userInformation *model.User
	var ok bool

	//1、先判断用户是否存在
	if ok, userInformation = user.IsExistUser(username); !ok {
		return "", code.CodeUserNotExist
	}
	//2、验证密码是否正确
	if userInformation.Password != utils.MD5(password) {
		return "", code.CodeInvalidPassword
	}
	//3、生成token
	token, err := myjwt.GenerateToken(userInformation.ID, userInformation.Username)
	if err != nil {
		return "", code.CodeServerBusy
	}

	return token, code.CodeSuccess
}

func Register(email, password, captcha string) (string, code.Code) {

	var ok bool
	var userInformation *model.User

	//1、先判断用户是否存在
	if ok, _ := user.IsExistUser(email); ok {
		return "", code.CodeUserExist
	}

	//2、从redis中获取验证码进行比对
	if ok, _ := shc_redis.CheckCaptchaForEmail(email, captcha); !ok {
		return "", code.CodeInvalidCaptcha
	}

	//3、生成11位的帐号
	username := utils.GetRandomNumbers(11)

	//4、注册到数据库
	if userInformation, ok = user.Register(username, email, password); !ok {
		return "", code.CodeServerBusy
	}

	//5、发送账号到用户邮箱
	if err := shc_email.SendCaptcha(email, username, user.UserNameMsg); err != nil {
		return "", code.CodeServerBusy
	}

	//6、生成token
	token, err := myjwt.GenerateToken(userInformation.ID, userInformation.Username)

	if err != nil {
		return "", code.CodeServerBusy
	}

	return token, code.CodeSuccess
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
