package limiter

import (
	"fmt"
	"github.com/go-redis/redis"
	"github.com/google/uuid"
	"time"
)

const keyPrefix = "limiter:slide_limiter:"

var luaScript = "local windows_start = tonumber(ARGV[3]) - tonumber(ARGV[1])\n" +
	"for _, v in pairs(KEYS) do\n" +
	"    redis.call('ZREMRANGEBYSCORE', v, 0, windows_start)\n " +
	"   local member = ARGV[3] .. ':' .. ARGV[4]\n" +
	"    local request_count = redis.call('ZCARD', v)\n" +
	"    if request_count == nil then\n" +
	"        request_count = 0\n " +
	"   end\n" +
	"    if request_count < tonumber(ARGV[2]) then\n" +
	"        -- 使用时间戳作为score，member为时间戳:唯一标识符\n" +
	"        redis.call('ZADD', v, ARGV[3], member)\n " +
	"   else\n" +
	"        -- 表示请求被拒绝\n" +
	"        return '0'\n" +
	"    end\n" +
	"end\n" +
	"-- 表示请求被允许\n" +
	"return '1'"

type RedisSlideLimiter struct {
	client *redis.Client // Redis 客户端
	window time.Duration // 时间窗口
	limit  int           // 访问次数
}

func NewRedisSlideLimiter(client *redis.Client, window time.Duration, limit int) *RedisSlideLimiter {
	return &RedisSlideLimiter{
		client: client,
		window: window,
		limit:  limit,
	}
}

func (sl *RedisSlideLimiter) Limit(keys []string) bool {
	var newKeys []string
	for _, key := range keys {
		newKeys = append(newKeys, keyPrefix+key)
	}
	eval := sl.client.Eval(luaScript, newKeys, sl.window.Milliseconds(), sl.limit, time.Now().UnixMilli(), uuid.New().String())
	if err := eval.Err(); err != nil {
		fmt.Println(err)
		return false
	} else {
		val := eval.Val().(string)
		if val == "1" {
			return true
		}
		return false
	}
}
