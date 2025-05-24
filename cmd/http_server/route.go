package http_server

import (
	"github.com/bingodfok/freshguard-boot/internal/fridge"
	"github.com/bingodfok/freshguard-boot/internal/msg"
	"github.com/bingodfok/freshguard-boot/internal/system"
	"github.com/gofiber/fiber/v2"
)

var (
	msgRouteFunc    = msg.Route()
	sysRouteFunc    = system.Route()
	fridgeRouteFunc = fridge.Route()
)

func BuildRoute() func(router fiber.Router) {
	return func(router fiber.Router) {
		router.Route("/msg", msgRouteFunc)
		router.Route("/sys", sysRouteFunc)
		router.Route("/fridge", fridgeRouteFunc)
	}
}
