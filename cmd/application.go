package main

import (
	"github.com/bingodfok/freshguard-boot/cmd/http_server"
	"github.com/bingodfok/freshguard-boot/pkg/auth"
	"github.com/bingodfok/freshguard-boot/pkg/resource/config"
	"github.com/bingodfok/freshguard-boot/pkg/suport/cache"
	"github.com/bingodfok/freshguard-boot/pkg/suport/mysql"
	"github.com/bingodfok/freshguard-boot/pkg/suport/sms"
	"github.com/bingodfok/freshguard-boot/pkg/suport/webserver"
	"github.com/go-redis/redis"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/spf13/viper"
	"sync"
	"xorm.io/xorm"
)

type Application struct {
	httpServer webserver.FiberServer
	waitGroup  *sync.WaitGroup
	viper      *viper.Viper
	xorm       *xorm.Engine
	redis      *redis.Client
	smsClient  *sms.TencentSMS
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
	// 初始化redis连接
	c := cache.Redis{
		Host:     app.viper.GetString("redis.host"),
		Port:     app.viper.GetInt("redis.port"),
		Password: app.viper.GetString("redis.password"),
		Database: app.viper.GetInt("redis.database"),
	}
	app.redis = c.InitRedisClient()
	// 启动fiber服务器
	fiberServer := webserver.FiberServer{
		AppName:           "FreshGuard",
		ContextPath:       viper.GetString("web.context-path"),
		Port:              viper.GetUint("web.port"),
		EnablePrintRoutes: true,
	}
	app.httpServer = fiberServer
	fiberServer.InitFiberServer()
	compress.New()
	fiberServer.UseMiddleware(cors.New(), auth.JwtAuthMiddleware([]string{"/"}, func() string {
		return app.viper.GetString("jwt.secret")
	}))
	fiberServer.Route(http_server.BuildRoute())
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
