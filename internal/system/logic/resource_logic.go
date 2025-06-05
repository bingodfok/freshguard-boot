package logic

import (
	"fmt"
	"github.com/bingodfok/freshguard-boot/cmd/ctx"
	"github.com/bingodfok/freshguard-boot/pkg/utils"
	"github.com/google/uuid"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/json"
	"time"
)

// SendSmsCodeLogic 发送短信验证码
func SendSmsCodeLogic(phone string, ctx *ctx.AppContext) (string, error) {
	code, err := utils.GenNumberString(4)
	if err != nil {
		return "", err
	}
	codeKey := uuid.New().String()
	bytes, err := json.Marshal(map[string]string{
		"codeKey": codeKey,
		"code":    code,
	})
	if err != nil {
		return "", err
	}
	val := ctx.Redis.Set("auth:sms_code:"+phone, string(bytes), time.Minute*5).Val()
	fmt.Println("val:", val)
	err = ctx.SmsClient.TencentSmsCodeSend(code, 5, phone)
	if err != nil {
		return "", err
	}
	return codeKey, nil
}
