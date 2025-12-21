package redis

import (
	"fmt"
	"shc-ai-demo/config"
)

// github.com/go-redis/redis/v8
// key:特定邮箱-> 验证码
func GenerateCaptcha(email string) string {
	return fmt.Sprintf(config.DefaultRedisKeyConfig.CaptchaPrefix, email)
}
