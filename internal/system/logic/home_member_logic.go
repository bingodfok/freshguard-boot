package logic

import (
	"github.com/bingodfok/freshguard-boot/cmd/ctx"
	"github.com/bingodfok/freshguard-boot/internal/system/repository/dao"
)

func GetHomeMemberByHome(ctx *ctx.AppContext, homeId int64) ([]*dao.HomeMember, error) {
	return dao.GetHomeMembersByHome(ctx.Xorm, homeId)
}
