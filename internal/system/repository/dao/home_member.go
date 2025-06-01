package dao

import (
	"time"
	"xorm.io/xorm"
)

type HomeMember struct {
	HomeId   int64
	UserId   int64
	Id       int64 `xorm:"pk autoincr"`
	CreateBy int64
	CreateAt time.Time `xorm:"created"`
	UpdateAt time.Time `xorm:"updated"`
	DeleteAt time.Time `xorm:"deleted"`
}

func (hm *HomeMember) TableName() string {
	return "home_member"
}

func (hm *HomeMember) Insert(xorm *xorm.Engine) error {
	_, err := xorm.Insert(hm)
	if err != nil {
		return err
	}
	return nil
}

func GetHomeMemberByUser(xorm *xorm.Engine, id int64) ([]*HomeMember, error) {
	hms := make([]*HomeMember, 0)
	err := xorm.Where("user_id = ?", id).Find(&hms)
	if err != nil {
		return nil, err
	}
	return hms, nil
}

func GetHomeMembersByHome(xorm *xorm.Engine, id int64) ([]*HomeMember, error) {
	hm := make([]*HomeMember, 0)
	err := xorm.Where("home_id=?", id).Find(&hm)
	if err != nil {
		return nil, err
	}
	return hm, nil
}
