package main

import (
	"douban-webend/api"
	"douban-webend/config"
	"douban-webend/utils"
	"time"
)

func main() {

	config.Init("config/config.json")

	utils.RegisterLogFile() // 注册日志创建任务

	utils.RegisterUploadLogTask(time.Hour * 4) // 每四个小时上传一次日志

	<-time.NewTimer(time.Second * 3).C // 延迟 3s，让日志启动

	api.InitRouter()

}
