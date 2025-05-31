package http_server

import (
	"github.com/bingodfok/freshguard-boot/cmd/ctx"
	"github.com/bingodfok/freshguard-boot/internal/fridge"
	"github.com/bingodfok/freshguard-boot/internal/msg"
	"github.com/bingodfok/freshguard-boot/internal/system"
	"github.com/gofiber/fiber/v2"
)

func BuildRoute(ctx *ctx.AppContext) func(router fiber.Router) {
	return func(router fiber.Router) {
		router.Route("/msg", msg.Route(ctx))
		router.Route("/sys", system.Route(ctx))
		router.Route("/fridge", fridge.Route(ctx))
	}
}
