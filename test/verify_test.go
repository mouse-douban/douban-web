package test

import (
	"douban-webend/config"
	"douban-webend/utils"
	"testing"
)

func TestSms(t *testing.T) {
	config.Init("../config/config.json")
	////err := utils.SendSMS("114514", "+8617805621625") 短信好贵啊，还是不发了
	//if err != nil {
	//	panic(err)
	//}
}

func TestEmail(t *testing.T) {
	config.Init("../config/config.json")
	err := utils.SendEmail("114514", "1545766400@qq.com")
	if err != nil {
		panic(err)
	}
}
