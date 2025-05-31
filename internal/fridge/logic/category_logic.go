package logic

import (
	"github.com/bingodfok/freshguard-boot/cmd/ctx"
	"github.com/bingodfok/freshguard-boot/internal/fridge/repository/dao"
)

func GetCategoryByHomeLogic(ctx *ctx.AppContext, home int64) ([]*dao.Category, error) {
	categories, err := dao.CategoryListByHome(ctx.Xorm, home)
	if err != nil {
		return nil, err
	}
	// 获取系统内置分类
	listByType, err := dao.CategoryListByType(ctx.Xorm, 1)
	categories = append(categories, listByType...)
	return categories, nil
}
