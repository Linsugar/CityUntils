package Untils

import (
	"context"
	"github.com/go-redis/redis/v8"
	"time"
)

var redisClient = redis.NewClient(&redis.Options{
	Addr: "10.10.8.222:26379",
})

type RedisInterface interface {
	RedisGet(key string) (data any)
	RedisSet(key string, value any, t time.Duration) (data any)
}

type RedisConf struct {
}

func (r *RedisConf) RedisGet(key string) (data any) {
	str := redisClient.Get(context.Background(), key).String()
	Info.Println(str)
	return str
}

func (r *RedisConf) RedisSet(key string, value any, t time.Duration) (data any) {
	result, err := redisClient.Set(context.Background(), key, value, t).Result()
	if err != nil {
		Error.Println(err.Error())
		return nil
	}
	Info.Println(result)
	return result
}
