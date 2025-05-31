package handler

import (
	"github.com/bingodfok/freshguard-boot/cmd/ctx"
	"github.com/bingodfok/freshguard-boot/internal/system/handler/dto"
	"github.com/bingodfok/freshguard-boot/internal/system/logic"
	"github.com/bingodfok/freshguard-boot/internal/system/repository/dao"
	"github.com/bingodfok/freshguard-boot/pkg/auth"
	"github.com/bingodfok/freshguard-boot/pkg/model/resp"
	"github.com/gofiber/fiber/v2"
)

// GetHomeDetailHandler 获取用户所在家庭详情
func GetHomeDetailHandler(ctx *ctx.AppContext) func(f *fiber.Ctx) error {
	return func(f *fiber.Ctx) error {
		claims := f.Locals("auth_context").(*auth.StandardClaims)
		id := claims.Id
		home, err := logic.GetHomeByUser(ctx, id)
		if err != nil {
			return err
		}
		detailResp := &dto.HomeDetailResp{
			Name:   home.Name,
			Id:     home.Id,
			Belong: home.Belong,
		}
		members, err := logic.GetHomeMemberByHome(ctx, home.Id)
		if err != nil {
			return err
		}
		var memberResps []dto.HomeMemberResp
		for _, member := range members {
			user, err := dao.SelectById(ctx.Xorm, member.UserId)
			if err != nil {
				return err
			}
			remark := ""
			remarks, err := dao.GetUserRemark(ctx.Xorm, member.UserId, id)
			if err != nil {
				return err
			}
			if remarks != nil {
				remark = remarks.Remark
			}
			homeMemberResp := dto.HomeMemberResp{
				Id:       member.UserId,
				Name:     user.Name,
				Avatar:   user.Avatar,
				AllowDel: home.Belong == id && member.UserId != id,
				Remark:   remark,
			}
			memberResps = append(memberResps, homeMemberResp)
		}
		detailResp.Members = memberResps
		return f.JSON(resp.Success(detailResp))
	}
}
