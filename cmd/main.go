package main

import (
	"douban-webend/api"
	"douban-webend/utils"
	"time"
)

func main() {

	utils.RegisterLogFile() // 注册日志创建任务

	utils.RegisterUploadLogTask(time.Hour * 4) // 每四个小时上传一次日志

	api.InitRouter()
}

