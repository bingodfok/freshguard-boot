package main

import (
	"github.com/bingodfok/freshguard-boot/pkg/resource/config"
	"github.com/bingodfok/freshguard-boot/pkg/suport/mysql"
	"github.com/bingodfok/freshguard-boot/pkg/suport/webserver"
	"github.com/spf13/viper"
	"sync"
	"xorm.io/xorm"
)

type Application struct {
	httpServer webserver.FiberServer
	waitGroup  *sync.WaitGroup
	viper      *viper.Viper
	xorm       *xorm.Engine
}

func NewApplication() *Application {
	// 加载配置文件
	configViper, err := config.LoadConfigViper("./cmd/etc/config.yaml")
	if err != nil {
		panic(err)
	}
	return &Application{
		waitGroup: &sync.WaitGroup{},
		viper:     configViper,
	}
}

func (app *Application) Run() {
	// 创建MySql连接
	xormSql := mysql.XormSql{
		Port:     app.viper.GetString("mysql.port"),
		Host:     app.viper.GetString("mysql.host"),
		Username: app.viper.GetString("mysql.username"),
		Password: app.viper.GetString("mysql.password"),
		Database: app.viper.GetString("mysql.database"),
	}
	app.xorm = xormSql.InitXorm()
	// 启动fiber服务器
	fiberServer := webserver.FiberServer{
		AppName:           "FreshGuard",
		ContextPath:       "/freshGuard/api",
		Port:              8080,
		EnablePrintRoutes: true,
	}
	app.httpServer = fiberServer
	fiberServer.InitFiberServer()
	app.waitGroup.Add(1)
	go func() {
		defer app.waitGroup.Done()
		err := fiberServer.Run()
		if err != nil {
			panic(err)
		}
	}()
	app.waitGroup.Wait()
}
