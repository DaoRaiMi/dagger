package redis

import (
	"github.com/daoraimi/dagger/config"
	goRedis "github.com/go-redis/redis"
)

var client *goRedis.Client
var Nil = goRedis.Nil

func R() *goRedis.Client {
	if client == nil {
		initRedis()
	}
	return client
}

func initRedis() {
	client = goRedis.NewClient(&goRedis.Options{
		Addr:     config.GetString("redis.addr"),
		Password: config.GetString("redis.password"),
		DB:       config.GetInt("redis.db"),
	})
	if _, err := client.Ping().Result(); err != nil {
		panic(err)
	}
}
