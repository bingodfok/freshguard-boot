package dto

import "github.com/bingodfok/freshguard-boot/internal/system/common"

type UserProfileRep struct {
	Username string           `json:"username"`
	Avatar   string           `json:"avatar"`
	UserType common.UserScope `json:"userType"`
	UserId   string           `json:"userId"`
	Phone    string           `json:"phone"`
}
