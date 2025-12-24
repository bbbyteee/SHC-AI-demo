package jwt

import (
	"log"
	"net/http"
	"shc-ai-demo/common/code"
	"shc-ai-demo/controller"
	"shc-ai-demo/utils/myjwt"
	"strings"

	"github.com/gin-gonic/gin"
)

// 读取jwt
func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		res := new(controller.Response) //准备统一返回结构

		var token string
		authHeader := c.GetHeader("Authorization")                        //从Header读取
		if authHeader != "" && strings.HasPrefix(authHeader, "Bearer ") { //Brarer是认证方案名
			token = strings.TrimPrefix(authHeader, "Bearer ")
		} else {
			// 兼容 URL 参数传 token
			token = c.Query("token")
		}

		if token == "" {
			c.JSON(http.StatusOK, res.CodeOf(code.CodeInvalidToken))
			c.Abort() //阻止后续中间件执行
			return
		}

		log.Println("token is ", token)
		userName, ok := myjwt.ParseToken(token)
		if !ok {
			c.JSON(http.StatusOK, res.CodeOf(code.CodeInvalidToken))
			c.Abort()
			return
		}

		c.Set("userName", userName)
		c.Next() //放行，不然直接卡在这
	}
}
