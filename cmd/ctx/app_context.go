package ctx

import (
	"github.com/bingodfok/freshguard-boot/pkg/suport/sms"
	"github.com/bingodfok/freshguard-boot/pkg/suport/webserver"
	"github.com/go-redis/redis"
	"github.com/spf13/viper"
	"sync"
	"xorm.io/xorm"
)

type AppContext struct {
	HttpServer webserver.FiberServer
	WaitGroup  *sync.WaitGroup
	Viper      *viper.Viper
	Xorm       *xorm.Engine
	Redis      *redis.Client
	SmsClient  *sms.TencentSMS
}
