package system

import (
	"github.com/bingodfok/freshguard-boot/cmd/ctx"
	"github.com/bingodfok/freshguard-boot/internal/system/handler"
	"github.com/gofiber/fiber/v2"
)

var (
	getProfileHandler = handler.GetProfileHandler
)

func Route(ctx *ctx.AppContext) func(router fiber.Router) {
	return func(router fiber.Router) {
		router.Get("/profile", getProfileHandler)
		router.Get("/resource/sms_captcha", handler.SmsCodeHandler(ctx))
		router.Post("/login/sms_captcha", handler.PhoneCaptchaLoginHandler(ctx))
	}
}
