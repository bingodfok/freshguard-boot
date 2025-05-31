package dto

import (
	"github.com/bingodfok/freshguard-boot/internal/system/common"
)

type PwdLoginReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type PhoneLoginReq struct {
	Phone      string `json:"phone"`
	Captcha    string `json:"captcha"`
	CaptchaKey string `json:"key"`
}

type LoginRep struct {
	UserId   string           `json:"userId"`
	HomeId   int64            `json:"homeId"`
	Token    string           `json:"accessToken"`
	Username string           `json:"username"`
	Avatar   string           `json:"avatar"`
	UserType common.UserScope `json:"userType"`
	Expires  int64            `json:"expires"`
}
