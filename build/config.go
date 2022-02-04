package main

import (
	"douban-webend/config"
	"douban-webend/utils"
	"os"
)

// 下载密钥
func main() {
	config.InitWithCOS()

	if config.Config.UseTLS { // 远程同步 key 文件
		utils.DownloadFile(os.Getenv("BUCKET_URL"), "tls_keys/douban-api.key", "config/api.key")
		utils.DownloadFile(os.Getenv("BUCKET_URL"), "tls_keys/douban-api.pem", "config/api.pem")
	}
}
