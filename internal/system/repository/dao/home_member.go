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

func (hm *HomeMember) Insert(xorm *xorm.Engine) error {
	_, err := xorm.Insert(hm)
	if err != nil {
		return err
	}
	return nil
}
