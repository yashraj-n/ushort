package repository

import (
	"log/slog"
	"os"
	"sync"

	"github.com/redis/go-redis/v9"
)

var redisSingleton *redis.Client
var redisSingletonLock sync.Mutex

func GetRedisInstance() *redis.Client {
	if redisSingleton != nil {
		return redisSingleton
	}

	redisSingletonLock.Lock()
	defer redisSingletonLock.Unlock()
	url := os.Getenv("REDIS_URL")
	redisUrlOpts, err := redis.ParseURL(url)
	if err != nil {
		slog.Error("Failed to parse Redis URL", slog.String("url", url), slog.String("error", err.Error()))
		panic("Invalid REDIS_URL: " + err.Error())
	}

	redisSingleton = redis.NewClient(redisUrlOpts)
	return redisSingleton
}
