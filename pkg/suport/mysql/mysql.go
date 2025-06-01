package mysql

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"time"
	"xorm.io/xorm"
)

var DB *xorm.Engine

type XormSql struct {
	Host     string
	Port     string
	Username string
	Password string
	Database string
}

func (sql *XormSql) InitXorm() *xorm.Engine {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		sql.Username,
		sql.Password,
		sql.Host,
		sql.Port,
		sql.Database,
	)
	engine, err := xorm.NewEngine("mysql", dsn)
	if err != nil {
		panic(err)
	}
	engine.DatabaseTZ = time.UTC
	engine.TZLocation = time.Local
	engine.SetMaxIdleConns(10)
	engine.SetMaxOpenConns(100)
	engine.ShowSQL(true)
	err = engine.Ping()
	if err != nil {
		panic(err)
	}
	DB = engine
	return engine
}
