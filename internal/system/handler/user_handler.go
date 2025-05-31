package handler

import (
	"github.com/bingodfok/freshguard-boot/cmd/ctx"
	"github.com/bingodfok/freshguard-boot/internal/system/common"
	"github.com/bingodfok/freshguard-boot/internal/system/handler/dto"
	"github.com/bingodfok/freshguard-boot/internal/system/logic"
	"github.com/bingodfok/freshguard-boot/pkg/auth"
	"github.com/bingodfok/freshguard-boot/pkg/model/resp"
	"github.com/gofiber/fiber/v2"
)

func GetUserProfileHandler(ctx *ctx.AppContext) func(c *fiber.Ctx) error {
	return func(f *fiber.Ctx) error {
		locals := f.Locals("auth_context")
		claims := locals.(*auth.StandardClaims)
		user, err := logic.GetUserByIdLogic(ctx, claims.Id)
		if err != nil {
			return err
		}
		if user == nil {
			// 没查询到当前用户登录过期
			return f.Status(fiber.StatusUnauthorized).JSON(resp.EmptyDataResult(resp.UnauthorizedCode))
		}
		profileRep := &dto.UserProfileRep{
			Username: user.Name,
			Avatar:   user.Avatar,
			Phone:    user.Phone,
			UserId:   user.UserId,
			UserType: common.Formal,
		}
		return f.Status(fiber.StatusOK).JSON(resp.Success(profileRep))
	}
}
