package system

import (
	"github.com/bingodfok/freshguard-boot/cmd/ctx"
	"github.com/bingodfok/freshguard-boot/internal/system/handler"
	"github.com/gofiber/fiber/v2"
)

func Route(ctx *ctx.AppContext) func(router fiber.Router) {
	return func(router fiber.Router) {
		router.Get("/user/profile", handler.GetUserProfileHandler(ctx))
		router.Get("/resource/sms_captcha", handler.SmsCodeHandler(ctx))
		router.Post("/login/sms_captcha", handler.PhoneCaptchaLoginHandler(ctx))
		router.Get("/home/detail", handler.GetHomeDetailHandler(ctx))
	}
}
