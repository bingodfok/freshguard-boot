package dto

import "github.com/bingodfok/freshguard-boot/internal/system/common"

type PwdLoginReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type PhoneLoginReq struct {
	Phone   string `json:"phone"`
	Captcha string `json:"captcha"`
}

type LoginRep struct {
	UserId   string           `json:"user_id"`
	Token    string           `json:"access_token"`
	Username string           `json:"username"`
	Avatar   string           `json:"avatar"`
	UserType common.UserScope `json:"user_type"`
}
