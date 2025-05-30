package dao

import (
	"time"
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
