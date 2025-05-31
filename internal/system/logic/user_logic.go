package logic

import (
	"github.com/bingodfok/freshguard-boot/cmd/ctx"
	"github.com/bingodfok/freshguard-boot/internal/system/repository/dao"
	"github.com/bingodfok/freshguard-boot/pkg/utils"
)

func GetUserByPhoneLogic(ctx *ctx.AppContext, phone string) (*dao.User, error) {
	return dao.SelectByPhone(ctx.Xorm, phone)
}

func GetUserByIdLogic(ctx *ctx.AppContext, id int64) (*dao.User, error) {
	user, err := dao.SelectById(ctx.Xorm, id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func GenUserByPhoneLogic(ctx *ctx.AppContext, phone string) (*dao.User, error) {
	userId, err := GenUserIdLogic(ctx)
	if err != nil {
		return nil, err
	}
	user := &dao.User{
		Phone:  phone,
		UserId: userId,
		Gender: "N",
		Name:   "User_" + utils.GenerateUserName(),
		Avatar: "https://file.darrenyou.cn/images/bear.jpeg",
	}
	_, err = dao.Insert(ctx.Xorm, user)
	if err != nil {
		return nil, err
	}
	// 新建该用户的家庭，并将用户添加到该家庭
	_, err = GenHomeByUserLogic(ctx, user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func GenUserIdLogic(ctx *ctx.AppContext) (string, error) {
	id, err := utils.GenerateNumericUserID(10)
	if err != nil {
		return "", err
	}
	user, err := dao.SelectByUserId(ctx.Xorm, id)
	if err != nil {
		return "", err
	}
	if user == nil {
		return id, nil
	}
	return GenUserIdLogic(ctx)
}
