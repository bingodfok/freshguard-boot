package dao

import "github.com/bingodfok/freshguard-boot/pkg/suport/mysql"

type Category struct {
	Name   string
	Icon   string
	HomeId int64
	mysql.BaseEntity
}

func (Category) TableName() string {
	return "category"
}
