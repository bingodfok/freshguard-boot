package fridge

import (
	"github.com/bingodfok/freshguard-boot/cmd/ctx"
	"github.com/bingodfok/freshguard-boot/internal/fridge/handler"
	"github.com/gofiber/fiber/v2"
)

func Route(ctx *ctx.AppContext) func(router fiber.Router) {
	return func(router fiber.Router) {
		router.Post("/fridge/create_fridge", handler.CreateFridgeHandler(ctx))
		router.Get("/fridge/list", handler.FridgeListHandler(ctx))
		router.Post("/fridge/edit", handler.FridgeEditHandler(ctx))
	}
}
