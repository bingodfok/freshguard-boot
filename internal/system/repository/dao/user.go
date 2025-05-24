package dao

import (
	"github.com/bingodfok/freshguard-boot/pkg/suport/mysql"
	"log"
)

var db = mysql.DB

type User struct {
	UserId   string `xorm:"unique"`
	Name     string
	Password string
	Phone    string
	Gender   string
	Avatar   string
	BaseEntity
}

func (u *User) TableName() string {
	return "user"
}

func Insert(user *User) bool {
	insert, err := db.Insert(user)
	if err != nil {
		log.Println(err)
		return false
	}
	return insert > 0
}

func UpdateById(user *User) bool {
	update, err := db.ID(user.Id).Update(user)
	if err != nil {
		log.Println(err)
		return false
	}
	return update > 0
}

func DeleteById(id int64) bool {
	deleted, err := db.ID(id).Delete()
	if err != nil {
		log.Println(err)
		return false
	}
	return deleted > 0
}

func SelectById(id int64) *User {
	user := &User{}
	err := db.ID(id).Find(user)
	if err != nil {
		log.Println(err)
		return nil
	}
	return user
}
