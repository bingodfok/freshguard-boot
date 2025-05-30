package handler

import (
	"github.com/bingodfok/freshguard-boot/cmd/ctx"
	"github.com/bingodfok/freshguard-boot/internal/system/handler/dto"
	"github.com/bingodfok/freshguard-boot/internal/system/logic"
	"github.com/gofiber/fiber/v2"
)

func PhoneCaptchaLoginHandler(ctx *ctx.AppContext) func(f *fiber.Ctx) error {
	return func(f *fiber.Ctx) error {
		req := &dto.PhoneLoginReq{}
		err := f.BodyParser(req)
		if err != nil {
			return fiber.ErrBadRequest
		}
		login, err := logic.PhoneCaptchaLoginLogic(ctx, req.CaptchaKey, req.Captcha, req.Phone)
		if err != nil {
			return err
		}
		return f.JSON(login)
	}
}
