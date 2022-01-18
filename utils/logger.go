package utils

import (
	"archive/zip"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"os"
	"time"
)

// RegisterLogFile 注册日志文件创建任务
func RegisterLogFile() {
	t := time.NewTicker(time.Hour * 24) // 每天创建一次日志文件
	gin.DisableConsoleColor()         // 关闭后台颜色显示

	go func(ticker *time.Ticker) {
		now := time.Now().Format("2006-01-02")

		os.MkdirAll("logs/"+now, os.ModePerm) // 创建一次日志文件夹

		ginLog, _ := os.Create("./logs/"+now+"/gin-"+now+".log") // 创建gin日志文件
		loggerLog, _ := os.Create("./logs/"+now+"/logger-"+now+".log") // 创建logger日志文件

		defer ginLog.Close()
		defer loggerLog.Close()

		gin.DefaultWriter = io.MultiWriter(ginLog, os.Stdout) // 设置gin log输出
		log.SetOutput(io.MultiWriter(loggerLog, os.Stdout)) // 设置logger输出
		<- t.C
	}(t)

}

func RegisterUploadLogTask(duration time.Duration) {
	t := time.NewTicker(duration)
	go func(ticker *time.Ticker) {
		<- t.C
		now := time.Now().Format("2006-01-02")
		src := "./logs/"+now	// 今日日志文件夹
		res, err := os.Open(src + "/log-" + now + ".zip")
		if err != nil {
			res, _ = os.Create(src + "/log-"+ now + ".zip")
		}
		defer res.Close()
		compressedLog(src, now, zip.NewWriter(res)) // 压缩
		// todo upload log to cos
	}(t)
}

// compressedLog 压缩日志文件
func compressedLog(src, now string, zw *zip.Writer) {
	ginLog, _ := os.Open(src + "/gin-" + now + ".log")
	loggerLog, _ := os.Open(src + "/logger-" + now + ".log")

	defer ginLog.Close()
	defer loggerLog.Close()

	defer func() {
		// 如果没有正常关闭就写入日志
		if err := zw.Close(); err != nil {
			log.Println(err)
		}
	}()

	info, _ := ginLog.Stat()
	header, _ := zip.FileInfoHeader(info)
	writer, _ := zw.CreateHeader(header)
	_, _ = io.Copy(writer, ginLog)


	info, _ = loggerLog.Stat()
	header, _ = zip.FileInfoHeader(info)
	writer, _ = zw.CreateHeader(header)
	_, _ = io.Copy(writer, loggerLog)

}