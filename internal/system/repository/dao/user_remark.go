package dao

import (
	"time"
	"xorm.io/xorm"
)

type UserRemark struct {
	Id       int64 `xorm:"pk autoincr"`
	TargetId int64
	Remark   string
	CreateBy int64
	CreateAt time.Time `xorm:"created"`
	UpdateAt time.Time `xorm:"updated"`
	DeleteAt time.Time `xorm:"deleted"`
}

func (UserRemark) TableName() string {
	return "user_remark"
}

func GetUserRemark(xorm *xorm.Engine, targetId int64, userId int64) (*UserRemark, error) {
	userRemark := &UserRemark{}
	has, err := xorm.Where("target_id = ? and create_by = ?", targetId, userId).Get(userRemark)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, nil
	}
	return userRemark, nil
}
