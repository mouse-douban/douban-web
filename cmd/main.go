package main

import (
	"douban-webend/api"
	"douban-webend/config"
	"douban-webend/dao"
	"douban-webend/utils"
	"time"
)

func main() {

	config.Init("config/config.json")

	//config.InitWithCOS() 部署时换成这个

	utils.RegisterLogFile() // 注册日志创建任务

	utils.RegisterUploadLogTask(time.Hour * 4) // 每四个小时上传一次日志

	<-time.NewTimer(time.Second * 2).C // 延迟 2s，让日志启动

	utils.ConnectRedis() // 连接 redisK

	dao.InitDao()

	api.InitRouter(false)

}
