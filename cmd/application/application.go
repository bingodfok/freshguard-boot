package application

import (
	"github.com/bingodfok/freshguard-boot/cmd/ctx"
	"github.com/bingodfok/freshguard-boot/cmd/http_server"
	"github.com/bingodfok/freshguard-boot/pkg/auth"
	"github.com/bingodfok/freshguard-boot/pkg/resource/config"
	"github.com/bingodfok/freshguard-boot/pkg/suport/cache"
	"github.com/bingodfok/freshguard-boot/pkg/suport/mysql"
	"github.com/bingodfok/freshguard-boot/pkg/suport/sms"
	"github.com/bingodfok/freshguard-boot/pkg/suport/webserver"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/spf13/viper"
	"sync"
)

func NewApplication() *ctx.AppContext {
	// 加载配置文件
	configViper, err := config.LoadConfigViper("./cmd/etc/config.yaml")
	if err != nil {
		panic(err)
	}
	return &ctx.AppContext{
		WaitGroup: &sync.WaitGroup{},
		Viper:     configViper,
	}
}

func Run(ctx *ctx.AppContext) {
	// 创建MySql连接
	xormSql := mysql.XormSql{
		Port:     ctx.Viper.GetString("mysql.port"),
		Host:     ctx.Viper.GetString("mysql.host"),
		Username: ctx.Viper.GetString("mysql.username"),
		Password: ctx.Viper.GetString("mysql.password"),
		Database: ctx.Viper.GetString("mysql.database"),
	}
	ctx.Xorm = xormSql.InitXorm()
	// 初始化redis连接
	c := cache.Redis{
		Host:     ctx.Viper.GetString("redis.host"),
		Port:     ctx.Viper.GetInt("redis.port"),
		Password: ctx.Viper.GetString("redis.password"),
		Database: ctx.Viper.GetInt("redis.database"),
	}
	ctx.Redis = c.InitRedisClient()
	// 初始化腾讯云短信发送客户端
	tencentSMS := &sms.TencentSMS{
		SecretId:   ctx.Viper.GetString("sms.tencent.secret-id"),
		SecretKey:  ctx.Viper.GetString("sms.tencent.secret-key"),
		SdkAppId:   ctx.Viper.GetString("sms.tencent.sdk-app-id"),
		SignName:   ctx.Viper.GetString("sms.tencent.sign-name"),
		TemplateId: ctx.Viper.GetString("sms.tencent.template-id"),
	}
	tencentSMS.InitTencentSms()
	ctx.SmsClient = tencentSMS
	// 启动fiber服务器
	fiberServer := webserver.FiberServer{
		AppName:           "FreshGuard",
		ContextPath:       viper.GetString("web.context-path"),
		Port:              viper.GetUint("web.port"),
		EnablePrintRoutes: true,
	}
	ctx.HttpServer = fiberServer
	fiberServer.InitFiberServer()
	compress.New()
	fiberServer.UseMiddleware(cors.New(), auth.JwtAuthMiddleware([]string{"/api/sys/sms_code"}, func() string {
		return ctx.Viper.GetString("jwt.secret")
	}))
	fiberServer.Route(http_server.BuildRoute(ctx))
	ctx.WaitGroup.Add(1)
	go func() {
		defer ctx.WaitGroup.Done()
		err := fiberServer.Run()
		if err != nil {
			panic(err)
		}
	}()
	ctx.WaitGroup.Wait()
}
