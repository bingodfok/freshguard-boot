package logic

import (
	"fmt"
	"github.com/bingodfok/freshguard-boot/cmd/ctx"
	"github.com/bingodfok/freshguard-boot/internal/system/common"
	"github.com/bingodfok/freshguard-boot/internal/system/handler/dto"
	"github.com/bingodfok/freshguard-boot/pkg/auth"
	"github.com/bingodfok/freshguard-boot/pkg/model/errors"
	"github.com/bingodfok/freshguard-boot/pkg/model/resp"
	"github.com/golang-jwt/jwt/v5"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/json"
	"time"
)

func PwdLoginLogic(name string, password string) (*dto.LoginRep, error) {
	return nil, nil
}

func PhoneCaptchaLoginLogic(ctx *ctx.AppContext, codeKey string, code string, phone string) (*dto.LoginRep, error) {
	val := ctx.Redis.Get("auth:sms_code:" + phone).Val()
	fmt.Println("val:", val)
	if val == "" {
		return nil, errors.NewBizErrorCode(resp.CaptchaExpCode)
	}
	ctx.Redis.Del("auth:sms_code:" + phone)
	codeResult := map[string]string{}
	err := json.Unmarshal([]byte(val), &codeResult)
	if err != nil {
		return nil, err
	}
	if codeResult["code"] == code && codeResult["codeKey"] == codeKey {
		// 清除验证码缓存
		ctx.Redis.Del("auth:sms_code:" + phone)
		user, err := GetUserByPhoneLogic(ctx, phone)
		if err != nil {
			return nil, err
		}
		fmt.Println(user)
		if user == nil {
			// 新建用户并且生成token
			user, err = GenUserByPhoneLogic(ctx, phone)
			if err != nil {
				return nil, err
			}
		}
		jwtAuth := auth.JwtAuth{
			SigningKey: ctx.Viper.GetString("jwt.secret"),
		}
		exp := time.Now().Add(time.Second * ctx.Viper.GetDuration("jwt.exp"))
		claims := auth.StandardClaims{
			Id:       user.Id,
			UserId:   user.UserId,
			Avatar:   user.Avatar,
			UserName: user.Name,
			RegisteredClaims: jwt.RegisteredClaims{
				NotBefore: jwt.NewNumericDate(time.Now()),
				ExpiresAt: jwt.NewNumericDate(exp),
				IssuedAt:  jwt.NewNumericDate(time.Now()),
				Issuer:    ctx.Viper.GetString("jwt.issuer"),
			},
		}
		token, err := jwtAuth.GenToken(claims)
		if err != nil {
			return nil, err
		}
		return &dto.LoginRep{
			Token:    token,
			UserId:   user.UserId,
			Username: user.Name,
			UserType: common.Formal,
			Avatar:   user.Avatar,
			Expires:  exp.UnixMilli(),
		}, nil
	} else {
		return nil, errors.NewBizErrorCode(resp.CaptchaErrorCode)
	}
}
