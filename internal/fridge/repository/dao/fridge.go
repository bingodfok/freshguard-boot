package dao

import (
	"time"
	"xorm.io/xorm"
)

type Fridge struct {
	HomeId   int64
	Name     string
	Id       int64 `xorm:"pk autoincr"`
	CreateBy int64
	CreateAt time.Time `xorm:"created"`
	UpdateAt time.Time `xorm:"updated"`
	DeleteAt time.Time `xorm:"deleted"`
}

func (fridge *Fridge) TableName() string {
	return "fridge"
}

func (fridge *Fridge) Insert(xorm *xorm.Engine) error {
	_, err := xorm.Insert(fridge)
	if err != nil {
		return err
	}
	return nil
}

func GetFridgeByHomeAndName(xorm *xorm.Engine, homeId int64, name string) (*Fridge, error) {
	friend := &Fridge{}
	has, err := xorm.Where("home_id = ? AND name = ?", homeId, name).Get(friend)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, nil
	}
	return friend, nil
}

func FridgeListByHome(xorm *xorm.Engine, homeId int64) ([]*Fridge, error) {
	friends := make([]*Fridge, 0)
	if err := xorm.Where("home_id = ?", homeId).Find(&friends); err != nil {
		return nil, err
	}
	return friends, nil
}
