package logic

import (
	"github.com/bingodfok/freshguard-boot/cmd/ctx"
	"github.com/bingodfok/freshguard-boot/internal/system/repository/dao"
)

// GenHomeByUserLogic 根据用户信息新建一个家庭
func GenHomeByUserLogic(ctx *ctx.AppContext, user *dao.User) (*dao.Home, error) {
	home := &dao.Home{
		Belong: user.Id,
		Name:   user.Name + "的家庭",
	}
	err := home.Insert(ctx.Xorm)
	if err != nil {
		return nil, err
	}
	member := &dao.HomeMember{
		UserId: user.Id,
		HomeId: home.Id,
	}
	err = member.Insert(ctx.Xorm)
	if err != nil {
		return nil, err
	}
	return home, nil
}
