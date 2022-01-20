package utils

import (
	"douban-webend/config"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	sms "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/sms/v20210111"
	"log"
	"math/rand"
	"net/http"
	"net/smtp"
	"strconv"
	"strings"
	"time"
)

const (
	EmailTemplate = `
		<!DOCTYPE html>
		<html lang="en">
		<head>
			<meta charset="iso-8859-15">
			<title>豆瓣API</title>
		</head>
		<body>
			【泡泡的泡的个人站】：验证码为 {0} 您正在使用邮箱进行验证，请勿泄露哦！
		</body>
		</html>`
	Subject  = "验证码"
	StmpHost = "smtp.qq.com"
	StmpPort = "587"
	SmsHost  = "sms.tencentcloudapi.com"
	ApRegion = "ap-guangzhou"
)

type VerifyInfo struct {
	EmailCode string `json:"email_code,omitempty"`
	SmsCode   string `json:"sms_code,omitempty"`
}

// verifyMap
// 集群
var verifyMap = make(map[uint64]map[string]string)

func VerifyInputCode(uid uint64, cType, code string) (bool, error) {
	// 先在自己内存里找一找
	if got, ok := verifyMap[uid][cType]; ok && got == code {
		return true, nil
	}

	if got, ok := verifyMap[uid][cType]; ok && got != code {
		return false, ServerError{
			HttpStatus: http.StatusBadRequest,
			Status:     40001,
			Info:       "invalid verify code",
			Detail:     cType + "验证码错误",
		}
	}
	var info VerifyInfo
	err := RedisGetStruct(strconv.Itoa(int(uid)), &info)
	if err != nil {
		return false, ServerInternalError
	}
	switch cType {
	case "email":
		if info.EmailCode == code {
			return true, nil
		}
		return false, ServerError{
			HttpStatus: http.StatusBadRequest,
			Status:     40001,
			Info:       "invalid verify code",
			Detail:     "email验证码错误",
		}
	case "sms":
		if info.SmsCode == code {
			return true, nil
		}
		return false, ServerError{
			HttpStatus: http.StatusBadRequest,
			Status:     40001,
			Info:       "invalid verify code",
			Detail:     "sms验证码错误",
		}
	}
	return false, ServerInternalError
}

// SendRandomVerifyCode 发送随机验证码
func SendRandomVerifyCode(uid uint64, vType string, target string) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	var sum = 0
	sum += 1 << 14 // 保证为5位数
	for i := 1; i <= 15; i++ {
		sum += r.Intn(2) << i
	}
	var vCode = strconv.Itoa(sum)
	SendVerifyCode(uid, vType, target, vCode)
}

// SendVerifyCode
// 异步发送，以免卡 主goroutine
func SendVerifyCode(uid uint64, vType, target, vCode string) {
	switch vType {
	case "email":
		verifyMap[uid] = map[string]string{"email": vCode}
		go func() {
			err := RedisSetStruct(strconv.Itoa(int(uid)), VerifyInfo{
				EmailCode: vCode,
			}, time.Minute*2) // 存进 redis 两分钟后过期

			if err != nil {
				log.Panicln(err)
				return
			}

			err = SendEmail(vCode, target)
			if err != nil {
				log.Panicln(err)
			}

		}()

		// 两分钟后删掉
		go func() {
			<-time.NewTimer(time.Minute * 2).C
			delete(verifyMap, uid)
		}()
	case "sms":
		verifyMap[uid] = map[string]string{"sms": vCode}
		go func() {
			err := RedisSetStruct(strconv.Itoa(int(uid)), VerifyInfo{
				SmsCode: vCode,
			}, time.Minute*2) // 存进 redis 两分钟后过期

			if err != nil {
				log.Panicln(err)
				return
			}

			err = SendSMS(vCode, target)
			if err != nil {
				log.Panicln(err)
			}
		}()

		// 两分钟后删掉
		go func() {
			<-time.NewTimer(time.Minute * 2).C
			delete(verifyMap, uid)
		}()
	}
}

// SendSMS
// 接入 sms.tencentcloudapi.com 发送短信
// 详细文档: https://cloud.tencent.com/document/api/382/55981
// 返回：
//      - nil				发送成功
//		- ServerError		发送失败
func SendSMS(verifyCode string, phoneNumbers ...string) error {

	// TODO 移到 controller 层
	//r := regexp.MustCompile("\\d")
	//alls := r.FindAllStringSubmatch(verifyCode, -1)
	//if len(alls) != len(verifyCode) {
	//	return ServerInternalError
	//}
	//
	//for _, number := range phoneNumbers {
	//	if ok, _ := regexp.MatchString("^\\+861[3-9][0-9]\\d{8}$", number); !ok { // todo 支持国际电话
	//		queryParamError := QueryParamError.Copy()
	//		queryParamError.Detail = "电话格式错误"
	//		return queryParamError
	//	}
	//}

	credential := common.NewCredential(
		config.Config.TencentSecretId,
		config.Config.TencentSecretKey,
	)
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = SmsHost
	client, _ := sms.NewClient(credential, ApRegion, cpf)

	request := sms.NewSendSmsRequest()

	request.PhoneNumberSet = common.StringPtrs(phoneNumbers)
	request.SmsSdkAppId = common.StringPtr(config.Config.TencentSmsSdkAppId)
	request.SignName = common.StringPtr(config.Config.TencentSignName)
	request.TemplateId = common.StringPtr(config.Config.TencentTemplateId)
	request.TemplateParamSet = common.StringPtrs([]string{verifyCode})

	response, err := client.SendSms(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		log.Printf("An API error has returned: %s", err)
		return ServerInternalError
	}
	if err != nil {
		return ServerInternalError
	}
	log.Printf("%s", response.ToJsonString())
	return nil
}

// SendEmail
// 发送邮件
// 返回：
//      - nil						发送成功
//		- ServerInternalError		发送失败
func SendEmail(verifyCode string, email ...string) error {
	for _, addr := range email {
		body := strings.Replace(EmailTemplate, "{0}", verifyCode, 1)
		msg := []byte("To: " + addr + "\r\nFrom: " + config.Config.EmailAuthSender + "<" + config.Config.EmailAuthAccount + ">" + "\r\nSubject: " + Subject + "\r\n" + "Content-Type: text/" + "html" + "; charset=UTF-8" + "\r\n\r\n" + body + "\r\n")
		auth := smtp.PlainAuth("", config.Config.EmailAuthAccount, config.Config.EmailAuthPassword, StmpHost)
		err := smtp.SendMail(StmpHost+":"+StmpPort, auth, config.Config.EmailAuthAccount, []string{addr}, msg)
		if err != nil {
			return ServerInternalError
		}
	}
	return nil
}
