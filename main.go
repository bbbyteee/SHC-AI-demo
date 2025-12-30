package main

import (
	"fmt"
	"log"
	"shc-ai-demo/common/aihelper"
	shc_mysql "shc-ai-demo/common/mysql"
	shc_rabbitmq "shc-ai-demo/common/rabbitmq"
	shc_redis "shc-ai-demo/common/redis"
	"shc-ai-demo/config"
	"shc-ai-demo/dao/message"
	"shc-ai-demo/router"
)

func StartServer(addr string, port int) error {
	r := router.InitRouter()

	return r.Run(fmt.Sprintf("%s:%d", addr, port))
}

// 从数据库加载消息并初始化 AIHelperManager
func readDataFromDB() error {
	manager := aihelper.GetGlobalManager()
	// 从数据库读取所有消息
	msgs, err := message.GetAllMessages()
	if err != nil {
		return err
	}
	// 遍历数据库消息
	for i := range msgs {
		m := &msgs[i]
		//默认openai模型
		modelType := "1"
		config := make(map[string]interface{})

		// 创建对应的 AIHelper
		helper, err := manager.GetOrCreateAIHelper(m.UserName, m.SessionID, modelType, config)
		if err != nil {
			log.Printf("[readDataFromDB] failed to create helper for user=%s session=%s: %v", m.UserName, m.SessionID, err)
			continue
		}
		log.Println("readDataFromDB init:  ", helper.SessionID)
		// 添加消息到内存中(不开启存储功能)
		helper.AddMessage(m.Content, m.UserName, m.IsUser, false)
	}

	log.Println("AIHelperManager init success ")
	return nil
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
	log.Printf("[InitMysql] DB addr = %p\n", shc_mysql.DB)
	//初始化AIHelperManager
	readDataFromDB()
	//初始化RabbitMQ
	shc_rabbitmq.InitRabbitMQ()
	log.Println("RabbitMQ连接成功")

	//启动服务
	if err := StartServer(host, port); err != nil {
		log.Fatalf("启动服务失败: %v", err)
	}
}
