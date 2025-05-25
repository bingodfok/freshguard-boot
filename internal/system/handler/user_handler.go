package handler

import (
	"github.com/gofiber/fiber/v2"
)

func GetProfileHandler(ctx *fiber.Ctx) error {
	err := ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"code": fiber.StatusOK,
	})
	if err != nil {
		return err
	}
	return nil
}
