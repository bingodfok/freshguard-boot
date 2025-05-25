package cache

import (
	"fmt"
	"github.com/go-redis/redis"
)

var Client *redis.Client

type Redis struct {
	Host     string
	Port     int
	Password string
	Database int
}

func (r *Redis) InitRedisClient() *redis.Client {
	Client = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", r.Host, r.Port),
		Password: r.Password,
		DB:       r.Database,
	})
	ping := Client.Ping()
	if pong, err := ping.Result(); err != nil {
		panic(err)
	} else {
		fmt.Println(pong)
	}
	return Client
}
