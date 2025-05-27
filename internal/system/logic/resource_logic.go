package logic

import (
	"github.com/bingodfok/freshguard-boot/cmd/ctx"
	"github.com/bingodfok/freshguard-boot/pkg/utils"
	"github.com/google/uuid"
	"time"
)

// SendSmsCodeLogic 发送短信验证码
func SendSmsCodeLogic(phone string, ctx *ctx.AppContext) (string, error) {
	code, err := utils.GenNumberString(6)
	if err != nil {
		return "", err
	}
	codeKey := uuid.New().String()
	ctx.Redis.Set("auth:sms_code:"+phone, map[string]string{
		"code_key": codeKey,
		"code":     code,
	}, time.Minute*5)
	err = ctx.SmsClient.TencentSmsCodeSend(code, 5, phone)
	if err != nil {
		return "", err
	}
	return code, nil
}
