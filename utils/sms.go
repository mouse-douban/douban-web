package utils

import (
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	sms "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/sms/v20210111"
	"log"
	"regexp"
)

// SendSMS
// 接入 sms.tencentcloudapi.com 发送短信
// 详细文档: https://cloud.tencent.com/document/api/382/55981
// 返回：
//      - nil				发送成功
//		- ServerError		发送失败
func SendSMS(verifyCode string, phoneNumbers ...string) error {

	r := regexp.MustCompile("\\d")
	alls := r.FindAllStringSubmatch(verifyCode, -1)
	if len(alls) != len(verifyCode) {
		return ServerInternalError
	}

	for _, number := range phoneNumbers {
		if ok, _ := regexp.MatchString("^\\+861[3-9][0-9]\\d{8}$", number); !ok { // todo 支持国际电话
			queryParamError := QueryParamError.Copy()
			queryParamError.Detail = "电话格式错误"
			return queryParamError
		}
	}

	credential := common.NewCredential(
		"SecretId",
		"SecretKey",
	)
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "sms.tencentcloudapi.com"
	client, _ := sms.NewClient(credential, "ap-guangzhou", cpf)

	request := sms.NewSendSmsRequest()

	request.PhoneNumberSet = common.StringPtrs(phoneNumbers)
	request.SmsSdkAppId = common.StringPtr("1400623650")
	request.SignName = common.StringPtr("泡泡的泡个人网")
	request.TemplateId = common.StringPtr("1281130")
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
