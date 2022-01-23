package test

import (
	"douban-webend/config"
	"douban-webend/utils"
	"log"
	"testing"
	"time"
)

func TestConnect(t *testing.T) {
	config.Init("../config/config.json") // 运行 test 的根目录不在项目的根目录那
	utils.ConnectRedis()
}

func TestSetterAndGetter(t *testing.T) {
	config.Init("../config/config.json") // 运行 test 的根目录不在项目的根目录那
	utils.ConnectRedis()
	err := utils.RedisSetString("hello", "world", time.Second*10)
	if err != nil {
		t.Error(err)
	}
	log.Println("Set success")
	ret, err := utils.RedisGetString("hello")
	if err != nil {
		t.Error(err)
	}

	log.Printf("Get %v", ret)
}
