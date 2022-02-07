package main

import (
	"douban-webend/api"
	"douban-webend/config"
	"douban-webend/dao"
	"douban-webend/utils"
	"os"
	"time"
)

// 需要配置环境变量
// BUCKET_URL  			-->  配置文件/密钥 储存桶链接
// TENCENT_SECRET_ID  	-->  腾讯云 secret id
// TENCENT_SECRET_KEY  	-->  腾讯云 secret key

var EnableLog = true
var EnableCOS = false

func main() {

	if EnableCOS {
		config.InitWithCOS("config2.json")
	} else {
		config.Init("config/config.json")
	}

	if config.Config.UseTLS { // 远程同步 key 文件
		utils.DownloadFile(os.Getenv("BUCKET_URL"), "tls_keys/douban-api.key", "config/api.key")
		utils.DownloadFile(os.Getenv("BUCKET_URL"), "tls_keys/douban-api.pem", "config/api.pem")
	}

	if EnableLog {
		utils.EnableLog()
		utils.RegisterUploadLogTask(time.Hour * 6) // 每六个小时上传一次日志

		<-time.NewTimer(time.Second).C // 延迟 1s，让日志启动
	}

	utils.ConnectRedis() // 连接 redis

	dao.InitDao()

	api.InitRouter(config.Config.UseTLS)

}
