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

// GetHomeByUser 获取用户当前所在家庭
func GetHomeByUser(ctx *ctx.AppContext, id int64) (*dao.Home, error) {
	homeMembers, err := dao.GenHomeMemberByUser(ctx.Xorm, id)
	if err != nil {
		return nil, err
	}
	var homeIds []int64
	for _, member := range homeMembers {
		homeIds = append(homeIds, member.HomeId)
	}
	homes, err := dao.ListByHomeIds(ctx.Xorm, homeIds)
	if err != nil {
		return nil, err
	}
	if len(homes) == 1 {
		return homes[0], nil
	} else {
		for _, home := range homes {
			if home.Belong != id {
				return home, nil
			}
		}
	}
	return nil, nil
}
