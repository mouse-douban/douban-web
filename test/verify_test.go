package test

import (
	"douban-webend/config"
	"douban-webend/model"
	"douban-webend/utils"
	"golang.org/x/crypto/bcrypt"
	"math/rand"
	"testing"
	"time"
)

func TestSms(t *testing.T) {
	config.Init("../config/config.json")
	////err := utils.SendSMS("114514", "+8617805621625") 短信好贵啊，还是不发了
	//if err != nil {
	//	t.Error(err)
	//}
}

func TestEmail(t *testing.T) {
	config.Init("../config/config.json")
	err := utils.SendEmail("114514", "1545766400@qq.com")
	if err != nil {
		t.Error(err)
	}
}

func TestVerifyEmailOk(t *testing.T) {
	config.Init("../config/config.json")
	utils.ConnectRedis()
	rand.Seed(time.Now().Unix())
	vCode := "114514"
	utils.SendVerifyCode("email", "1545766400@qq.com", vCode)
	<-time.NewTimer(time.Second * 5).C // 等五秒
	err := utils.VerifyInputCode("1545766400@qq.com", "email", vCode)
	if err != nil {
		t.Error(err)
	}
}

func TestVerifyEmailFailed(t *testing.T) {
	config.Init("../config/config.json")
	utils.ConnectRedis()
	rand.Seed(time.Now().Unix())
	vCode := "11451"
	vCodeFailed := "114514"
	utils.SendVerifyCode("email", "1545766400@qq.com", vCode)
	<-time.NewTimer(time.Second * 3).C // 等三秒
	err := utils.VerifyInputCode("1545766400@qq.com", "email", vCodeFailed)
	if err == nil {
		t.Error(err)
	}
}

func TestVerifyPassword(t *testing.T) {
	user := model.User{PlaintPassword: "114514"}
	hash := user.EncryptPassword()
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte("114514"))
	if err != nil {
		t.Error(err)
	}
}
