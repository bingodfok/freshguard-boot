package handler

import (
	"github.com/bingodfok/freshguard-boot/cmd/ctx"
	"github.com/bingodfok/freshguard-boot/internal/fridge/handler/dto"
	"github.com/bingodfok/freshguard-boot/internal/fridge/logic"
	syslogic "github.com/bingodfok/freshguard-boot/internal/system/logic"
	"github.com/bingodfok/freshguard-boot/pkg/auth"
	"github.com/bingodfok/freshguard-boot/pkg/model/resp"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

// CreateFridgeHandler 创建冰箱
func CreateFridgeHandler(ctx *ctx.AppContext) func(f *fiber.Ctx) error {
	return func(f *fiber.Ctx) error {
		req := &dto.CreateFridgeReq{}
		if err := f.BodyParser(req); err != nil {
			return fiber.ErrBadRequest
		}
		if len(req.Name) == 0 {
			return f.Status(fiber.StatusBadRequest).JSON(resp.CodeMsgResult(resp.BadRequestCode, "冰箱名不能为空"))
		}
		claims := f.Locals("auth_context").(*auth.StandardClaims)
		_, err := logic.CreateFridgeByUserLogic(ctx, req.Name, claims.Id)
		if err != nil {
			return err
		}
		return f.JSON(resp.EmptyDataResult(resp.SuccessCode))
	}
}

// FridgeListHandler 所在家庭冰箱列表
func FridgeListHandler(ctx *ctx.AppContext) func(*fiber.Ctx) error {
	return func(f *fiber.Ctx) error {
		claims := f.Locals("auth_context").(*auth.StandardClaims)
		home, err := syslogic.GetHomeByUser(ctx, claims.Id)
		if err != nil {
			return err
		}
		fridges, err := logic.FridgeListLogic(ctx, claims.Id)
		if err != nil {
			return err
		}
		fridgeReps := make([]dto.FridgeRep, 0)
		for _, fridge := range fridges {
			createBy, err := syslogic.GetUserByIdLogic(ctx, fridge.CreateBy)
			if err != nil {
				return err
			}
			dtoRep := dto.FridgeRep{
				Name:    fridge.Name,
				Id:      fridge.Id,
				Admin:   createBy.Name,
				CanEdit: claims.Id == home.Belong || claims.Id == createBy.Id,
			}
			fridgeReps = append(fridgeReps, dtoRep)
		}
		return f.JSON(resp.Success(fridgeReps))
	}
}

// FridgeEditHandler 编辑冰箱
func FridgeEditHandler(ctx *ctx.AppContext) func(*fiber.Ctx) error {
	return func(f *fiber.Ctx) error {
		req := &dto.EditFridgeReq{}
		if err := f.BodyParser(req); err != nil {
			return fiber.ErrBadRequest
		}
		if len(req.Name) == 0 {
			return f.Status(fiber.StatusBadRequest).JSON(resp.CodeMsgResult(resp.BadRequestCode, "冰箱名不能为空"))
		}
		if req.Id == 0 {
			return f.Status(fiber.StatusBadRequest).JSON(resp.CodeMsgResult(resp.BadRequestCode, "ID不能为空"))
		}
		claims := f.Locals("auth_context").(*auth.StandardClaims)
		err := logic.FridgeEditLogic(ctx, req, claims.Id)
		if err != nil {
			return err
		}
		return f.JSON(resp.EmptyDataResult(resp.SuccessCode))
	}
}

// FridgeDelHandler 删除冰箱
func FridgeDelHandler(ctx *ctx.AppContext) func(*fiber.Ctx) error {
	return func(f *fiber.Ctx) error {
		fridgeIdStr := f.Query("id")
		if fridgeIdStr == "" {
			return f.Status(fiber.StatusBadRequest).JSON(resp.CodeMsgResult(resp.BadRequestCode, "ID不能为空"))
		}
		fridgeId, err := strconv.ParseInt(fridgeIdStr, 10, 64)
		if err != nil {
			return err
		}
		claims := f.Locals("auth_context").(*auth.StandardClaims)
		err = logic.FridgeDelLogic(ctx, fridgeId, claims.Id)
		if err != nil {
			return err
		}
		return f.JSON(resp.EmptyDataResult(resp.SuccessCode))
	}
}
