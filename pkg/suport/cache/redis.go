package cache

import (
	"fmt"
	"github.com/go-redis/redis"
)

type Redis struct {
	Host     string
	Port     int
	Password string
	Database int
}

func (r *Redis) InitRedisClient() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", r.Host, r.Port),
		Password: r.Password,
		DB:       r.Database,
	})
	ping := client.Ping()
	if pong, err := ping.Result(); err != nil {
		panic(err)
	} else {
		fmt.Println(pong)
	}
	return client
}
