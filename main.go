package main

import (
	"fmt"
	"log"
	shc_mysql "shc-ai-demo/common/mysql"
	shc_redis "shc-ai-demo/common/redis"
	"shc-ai-demo/config"
	"shc-ai-demo/router"
)

func StartServer(addr string, port int) error {
	r := router.InitRouter()

	return r.Run(fmt.Sprintf("%s:%d", addr, port))
}

func main() {
	conf := config.GetConfig()
	host := conf.MainConfig.Host
	port := conf.MainConfig.Port
	//初始化redis
	shc_redis.Init()
	log.Println("Redis连接成功!")
	//初始化数据库
	shc_mysql.InitMysql()
	log.Println("Mysql连接成功!")

	//启动服务
	if err := StartServer(host, port); err != nil {
		log.Fatalf("启动服务失败: %v", err)
	}
}
