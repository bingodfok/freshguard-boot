package dao

import (
	"time"
	"xorm.io/xorm"
)

type Home struct {
	Belong   int64
	Name     string
	Id       int64 `xorm:"pk autoincr"`
	CreateBy int64
	CreateAt time.Time `xorm:"created"`
	UpdateAt time.Time `xorm:"updated"`
	DeleteAt time.Time `xorm:"deleted"`
}

func (h *Home) TableName() string {
	return "home"
}

func (h *Home) Insert(xorm *xorm.Engine) error {
	_, err := xorm.Insert(h)
	if err != nil {
		return err
	}
	return nil
}

func ListByHomeIds(xorm *xorm.Engine, hid []int64) ([]*Home, error) {
	hms := make([]*Home, 0)
	err := xorm.In("id", hid).Find(&hms)
	if err != nil {
		return nil, err
	}
	return hms, nil
}
