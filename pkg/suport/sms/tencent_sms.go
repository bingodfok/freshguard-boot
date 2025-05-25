package sms

import (
	"fmt"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	tencentsms "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/sms/v20210111"
	"strconv"
)

var (
	tencent *tencentsms.Client
)

type TencentSMS struct {
	SecretId   string
	SecretKey  string
	SdkAppId   string
	SignName   string
	TemplateId string
}

func (ts *TencentSMS) InitTencentSms() *tencentsms.Client {
	// 腾讯云
	clientProfile := profile.NewClientProfile()
	credential := common.NewCredential(ts.SecretId, ts.SecretKey)
	if client, err := tencentsms.NewClient(credential, "ap-guangzhou", clientProfile); err != nil {
		panic(err)
	} else {
		tencent = client
	}
	return tencent
}

// TencentSmsCodeSend 腾讯云短信验证码
func (ts *TencentSMS) TencentSmsCodeSend(code string, exp int, mobile ...string) error {
	request := tencentsms.NewSendSmsRequest()
	request.SmsSdkAppId = common.StringPtr(ts.SdkAppId)
	request.SignName = common.StringPtr(ts.SignName)
	request.TemplateId = common.StringPtr(ts.TemplateId)
	request.PhoneNumberSet = common.StringPtrs(mobile)
	request.TemplateParamSet = common.StringPtrs([]string{code, strconv.Itoa(exp)})
	if response, err := tencent.SendSms(request); err != nil {
		return err
	} else {
		fmt.Println(response.ToJsonString())
		return nil
	}
}
