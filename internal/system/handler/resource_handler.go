package handler

import (
	"github.com/bingodfok/freshguard-boot/cmd/ctx"
	"github.com/bingodfok/freshguard-boot/internal/system/logic"
	"github.com/bingodfok/freshguard-boot/pkg/model/resp"
	"github.com/bingodfok/freshguard-boot/pkg/utils"
	"github.com/gofiber/fiber/v2"
)

// SmsCodeHandler 发送短信验证码
func SmsCodeHandler(ctx *ctx.AppContext) func(f *fiber.Ctx) error {
	return func(f *fiber.Ctx) error {
		phone := f.Query("phone")
		if phone == "" {
			return f.JSON(resp.CodeMsgResult(resp.BadRequestCode, "手机号不能为空"))
		} else if !utils.PhoneMatch(phone) {
			return f.JSON(resp.CodeMsgResult(resp.BadRequestCode, "手机号格式错误"))
		}
		codeKey, err := logic.SendSmsCodeLogic(phone, ctx)
		if err != nil {
			return err
		}
		return f.JSON(resp.Success(map[string]string{
			"codeKey": codeKey,
		}))
	}
}
