package dao

import "github.com/bingodfok/freshguard-boot/pkg/suport/mysql"

type HomeMember struct {
	HomeId int64
	UserId int64
	mysql.BaseEntity
}
