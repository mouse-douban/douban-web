package utils

import (
	"archive/zip"
	"douban-webend/config"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"os"
	"time"
)

// RegisterLogFile 注册日志文件创建任务
func RegisterLogFile() {
	t := time.NewTicker(time.Hour * 24) // 每天创建一次日志文件
	gin.DisableConsoleColor()           // 关闭后台颜色显示

	go func(ticker *time.Ticker) {

		for true {
			createLogFile()
			<-t.C
		}

	}(t)

}

func createLogFile() {

	now := time.Now().Format("2006-01-02")
	log.Printf("正在创建今日: %v 的日志文件\n", now)

	err := os.MkdirAll("logs/"+now, os.ModePerm) // 创建一次日志文件夹
	if err != nil {
		log.Println("创建失败！", err)
		return
	}

	ginLog, err := os.Create("./logs/" + now + "/gin-" + now + ".log") // 创建gin日志文件
	if err != nil {
		log.Println("创建失败！", err)
		return
	}

	loggerLog, err := os.Create("./logs/" + now + "/logger-" + now + ".log") // 创建logger日志文件
	if err != nil {
		log.Println("创建失败！", err)
		return
	}

	gin.DefaultWriter = io.MultiWriter(ginLog, os.Stdout) // 设置gin log输出
	log.SetOutput(io.MultiWriter(loggerLog, os.Stdout))   // 设置logger输出
}

func RegisterUploadLogTask(duration time.Duration) {
	t := time.NewTicker(duration)
	go func(ticker *time.Ticker) {

		for true {
			<-t.C // 并非启动就上传日志

			now := time.Now().Format("2006-01-02")
			log.Printf("正在上传今日: %v 的日志文件\n", now)

			src := "./logs/" + now // 今日日志文件夹
			target := src + "/log-" + now + ".zip"
			res, err := os.Open(target)
			if err != nil {
				res, _ = os.Create(target)
			} else {
				err = os.Remove(target)
				if err != nil {
					log.Println(err)
				}
				res, _ = os.Create(target)
			}

			compressedLog(src, now, zip.NewWriter(res)) // 压缩

			upload, err := os.Open(target)
			if err != nil {
				log.Println(err)
			}
			UploadFile(config.Config.TencentLogBucketUrl, "/log-"+now+".zip", upload) // 上传到 cos
			err = upload.Close()

		}

	}(t)
}

// compressedLog 压缩日志文件
func compressedLog(src, now string, zw *zip.Writer) {
	ginLog, err := os.Open(src + "/gin-" + now + ".log")
	if err != nil {
		createLogFile()
	}
	loggerLog, err := os.Open(src + "/logger-" + now + ".log")
	if err != nil {
		log.Printf("上传日志失败!! 原因: %v", err)
		return
	}

	defer func(ginLog *os.File) {
		err := ginLog.Close()
		if err != nil {
			log.Println("关闭ginLog文件失败！", err)
		}
	}(ginLog)
	defer func(loggerLog *os.File) {
		err := loggerLog.Close()
		if err != nil {
			log.Println("关闭logger文件失败！", err)
		}
	}(loggerLog)

	defer func() {
		// 如果没有正常关闭就写入日志
		if err := zw.Close(); err != nil {
			log.Println(err)
		}
	}()

	// 一路梭哈
	info, _ := ginLog.Stat()
	header, _ := zip.FileInfoHeader(info)
	writer, _ := zw.CreateHeader(header)
	_, _ = io.Copy(writer, ginLog)

	info, _ = loggerLog.Stat()
	header, _ = zip.FileInfoHeader(info)
	writer, _ = zw.CreateHeader(header)
	_, _ = io.Copy(writer, loggerLog)

}
