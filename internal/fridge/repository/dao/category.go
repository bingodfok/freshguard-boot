package dao

import (
	"time"
	"xorm.io/xorm"
)

type Category struct {
	Name     string
	ImgPath  string
	HomeId   int64
	Type     int
	Id       int64 `xorm:"pk autoincr"`
	CreateBy int64
	CreateAt time.Time `xorm:"created"`
	UpdateAt time.Time `xorm:"updated"`
	DeleteAt time.Time `xorm:"deleted"`
}

func (c *Category) TableName() string {
	return "category"
}

func CategoryListByHome(xorm *xorm.Engine, homeId int64) ([]*Category, error) {
	var categories []*Category
	err := xorm.Where("home_id = ?", homeId).Find(&categories)
	if err != nil {
		return nil, err
	}
	return categories, nil
}

func CategoryListByType(xorm *xorm.Engine, tp int64) ([]*Category, error) {
	var categories []*Category
	err := xorm.Where("type = ?", tp).Find(&categories)
	if err != nil {
		return nil, err
	}
	return categories, nil
}
