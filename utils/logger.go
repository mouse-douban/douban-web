package utils

import (
	"douban-webend/config"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"os"
	"time"
)

var (
	loggerFile     *os.File
	loggerCreateAt time.Time
	updateDuration int // 距离创建日期下一天的时间间隔 (s)
)

type GinWriter struct {
}

func (l *GinWriter) Write(p []byte) (n int, err error) {
	size := len(p)
	logGin(p[:size-1]) // 切掉 \n
	return size, nil
}

// EnableLog 开启日志
func EnableLog() {
	gin.DisableConsoleColor()
	now := time.Now().Format("2006-01-02")

	_, err := os.Stat("logs")
	if err != nil { // 检查文件是否存在
		err = os.Mkdir("logs", 0777) // 创建文件夹
		if err != nil {
			log.Fatalln("日志开启失败! 原因: ", err)
		}
	}

	loggerFile, err = os.Create("./logs/" + now + ".log")
	if err != nil {
		log.Fatalln("日志开启失败! 原因: ", err)
	}

	// 设置创建时间
	loggerCreateAt = time.Now()

	// 更新间隔
	n := time.Now()
	updateDuration = 24*60*60 - (n.Hour()*60*60 + n.Minute()*60 + n.Second())

	// 设置输出
	gin.DefaultWriter = &GinWriter{}

	log.SetOutput(io.MultiWriter(loggerFile, os.Stdout))
}

func RegisterUploadLogTask(duration time.Duration) {
	ticker := time.NewTicker(duration)
	go func() {
		for {
			<-ticker.C
			now := time.Now().Format("2006-01-02")
			LoggerInfo("正在上传日志文件", "./logs/"+now+".log")
			file, err := os.Open("./logs/" + now + ".log")
			if err != nil {
				file, err = os.Create("./logs/" + now + ".log")
				if err != nil {
					LoggerWarning("上传失败! 原因: ", err)
					continue
				}
				loggerFile = file
			}
			UploadFile(config.Config.TencentLogBucketUrl, "logs/"+now+".log", file)
		}
	}()
}

// 检查时间
func checkLoggerFile() {
	now := time.Now()
	if int(now.Sub(loggerCreateAt).Seconds()) > updateDuration {
		var err error
		loggerFile, err = os.Create("./logs/" + now.Format("2006-01-02") + ".log")
		if err != nil {
			LoggerWarning("日志更新失败! 原因: ", err)
		}
		log.SetOutput(io.MultiWriter(loggerFile, os.Stdout))
	}
}

func logGin(p []byte) {
	log.SetPrefix("")
	checkLoggerFile()
	log.Println(string(p))
}

func LoggerInfo(mess ...interface{}) {
	log.SetPrefix("[INFO] ")
	checkLoggerFile()
	log.Println(mess...)
}

func LoggerWarning(mess ...interface{}) {
	log.SetPrefix("[WARNING] ")
	checkLoggerFile()
	log.Println(mess...)
}

type Closed interface {
	Close() error
}

func LoggerError(mess string, c Closed) {
	err := c.Close()
	if err != nil {
		log.SetPrefix("[ERROR] ")
		checkLoggerFile()
		log.Println(mess, err)
	}
}

func LoggerPanic(mess ...interface{}) {
	log.SetPrefix("[PANIC] ")
	checkLoggerFile()
	log.Panicln(mess...)
}

func LoggerFatal(mess ...interface{}) {
	log.SetPrefix("[FATAL] ")
	checkLoggerFile()
	log.Fatalln(mess...)
}
