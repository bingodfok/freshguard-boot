package dao

import (
	"log"
	"time"
	"xorm.io/xorm"
)

type User struct {
	UserId   string `xorm:"unique"`
	Name     string
	Password string
	Phone    string
	Gender   string
	Avatar   string
	Id       int64 `xorm:"pk autoincr"`
	CreateBy int64
	CreateAt time.Time `xorm:"created"`
	UpdateAt time.Time `xorm:"updated"`
	DeleteAt time.Time `xorm:"deleted"`
}

func (u *User) TableName() string {
	return "user"
}

func Insert(xorm *xorm.Engine, user *User) (bool, error) {
	insert, err := xorm.Insert(user)
	if err != nil {
		log.Println(err)
		return false, err
	}
	return insert > 0, nil
}

func UpdateById(xorm *xorm.Engine, user *User) (bool, error) {
	update, err := xorm.ID(user.Id).Update(user)
	if err != nil {
		log.Println(err)
		return false, err
	}
	return update > 0, nil
}

func DeleteById(xorm *xorm.Engine, id int64) (bool, error) {
	deleted, err := xorm.ID(id).Delete()
	if err != nil {
		log.Println(err)
		return false, err
	}
	return deleted > 0, nil
}

func SelectById(xorm *xorm.Engine, id int64) (*User, error) {
	user := &User{}
	has, err := xorm.ID(id).Get(user)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	if !has {
		return nil, nil
	}
	return user, nil
}

func SelectByUserId(xorm *xorm.Engine, userId string) (*User, error) {
	user := &User{}
	has, err := xorm.Where("user_id = ?", userId).Get(user)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	if !has {
		return nil, nil
	}
	return user, nil
}

func SelectByPhone(xorm *xorm.Engine, phone string) (*User, error) {
	user := &User{}
	has, err := xorm.Where("phone = ?", phone).Get(user)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	if !has {
		return nil, nil
	}
	return user, err
}
