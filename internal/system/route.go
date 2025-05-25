package system

import (
	"github.com/bingodfok/freshguard-boot/internal/system/handler"
	"github.com/gofiber/fiber/v2"
)

var (
	getProfileHandler = handler.GetProfileHandler
)

func Route() func(router fiber.Router) {
	return func(router fiber.Router) {
		router.Get("/profile", getProfileHandler)
	}
}
