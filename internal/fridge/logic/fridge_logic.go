package logic

import (
	"github.com/bingodfok/freshguard-boot/cmd/ctx"
	"github.com/bingodfok/freshguard-boot/internal/fridge/repository/dao"
	"github.com/bingodfok/freshguard-boot/internal/system/logic"
	"github.com/bingodfok/freshguard-boot/pkg/model/errors"
	"github.com/bingodfok/freshguard-boot/pkg/model/resp"
)

// CreateFridgeByUserLogic 创建冰箱
func CreateFridgeByUserLogic(ctx *ctx.AppContext, name string, userId int64) (*dao.Fridge, error) {
	home, err := logic.GetHomeByUser(ctx, userId)
	if err != nil {
		return nil, err
	}
	fridged, err := dao.GetFridgeByHomeAndName(ctx.Xorm, home.Id, name)
	if err != nil {
		return nil, err
	}
	if fridged != nil {
		return nil, errors.NewBizErrorCode(resp.FridgeExistCode)
	}
	fridge := &dao.Fridge{
		Name:     name,
		HomeId:   home.Id,
		CreateBy: userId,
	}
	if fridge.Insert(ctx.Xorm) != nil {
		return nil, err
	}
	return fridge, nil
}

func FridgeListLogic(ctx *ctx.AppContext, userId int64) ([]*dao.Fridge, error) {
	home, err := logic.GetHomeByUser(ctx, userId)
	if err != nil {
		return nil, err
	}
	fridges, err := dao.FridgeListByHome(ctx.Xorm, home.Id)
	if err != nil {
		return nil, err
	}
	return fridges, nil
}
