package handler

import (
	"github.com/bingodfok/freshguard-boot/cmd/ctx"
	"github.com/bingodfok/freshguard-boot/internal/fridge/handler/dto"
	"github.com/bingodfok/freshguard-boot/internal/fridge/logic"
	syslogic "github.com/bingodfok/freshguard-boot/internal/system/logic"
	"github.com/bingodfok/freshguard-boot/pkg/auth"
	"github.com/bingodfok/freshguard-boot/pkg/model/resp"
	"github.com/gofiber/fiber/v2"
)

// CategoryListHandler 食物分类列表
func CategoryListHandler(ctx *ctx.AppContext) func(f *fiber.Ctx) error {
	return func(f *fiber.Ctx) error {
		claims := f.Locals("auth_context").(*auth.StandardClaims)
		home, err := syslogic.GetHomeByUser(ctx, claims.Id)
		if err != nil {
			return err
		}
		categories, err := logic.GetCategoryByHomeLogic(ctx, home.Id)
		if err != nil {
			return err
		}
		var resps []dto.CategoryRsp
		for _, category := range categories {
			if category.Type == 1 { // 系统内置
				resps = append(resps, dto.CategoryRsp{
					Id:      category.Id,
					Name:    category.Name,
					ImgPath: category.ImgPath,
					Type:    category.Type,
					CanEdit: false,
				})
			}
			if category.Type == 2 { // 自定义
				user, err := syslogic.GetUserByIdLogic(ctx, category.CreateBy)
				if err != nil {
					return err
				}
				resps = append(resps, dto.CategoryRsp{
					Id:      category.Id,
					Name:    category.Name,
					ImgPath: category.ImgPath,
					Type:    category.Type,
					Admin:   user.Name,
					CanEdit: category.CreateBy == claims.Id || claims.Id == home.Belong,
				})
			}
		}
		return f.JSON(resp.Success(resps))
	}
}
