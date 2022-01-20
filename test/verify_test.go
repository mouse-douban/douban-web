package test

import (
	"douban-webend/config"
	"douban-webend/utils"
	"math/rand"
	"testing"
	"time"
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

func TestVerifyEmailOk(t *testing.T) {
	config.Init("../config/config.json")
	utils.ConnectRedis()
	rand.Seed(time.Now().Unix())
	ruid := rand.Uint64()
	vCode := "114514"
	utils.SendVerifyCode(ruid, "email", "1545766400@qq.com", vCode)
	<-time.NewTimer(time.Second * 5).C // 等五秒
	ok, err := utils.VerifyInputCode(ruid, "email", vCode)
	if err != nil {
		panic(err)
	}
	if !ok {
		panic("验证码错误")
	}
}

func TestVerifyEmailFailed(t *testing.T) {
	config.Init("../config/config.json")
	utils.ConnectRedis()
	rand.Seed(time.Now().Unix())
	ruid := rand.Uint64()
	vCode := "114514"
	vCodeFailed := "11451"
	utils.SendVerifyCode(ruid, "email", "1545766400@qq.com", vCode)
	<-time.NewTimer(time.Second * 5).C // 等五秒
	ok, err := utils.VerifyInputCode(ruid, "email", vCodeFailed)
	if err == nil || ok {
		panic("有很大问题")
	}
}
