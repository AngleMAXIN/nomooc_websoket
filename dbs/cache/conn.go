package cache

import (
	prcf "Project/websocket/config"
	"fmt"
	"time"

	"github.com/go-redis/redis"
)

func init() {
	redisDao = GetRedisConn()
}

// GetRedisConn 获得redis conn
func GetRedisConn() *redis.Client {
	cf := prcf.ProConfig.GetRedisConfig()
	client := redis.NewClient(&redis.Options{
		Addr:         fmt.Sprintf("%s:%d", cf.Host, cf.Port),
		DialTimeout:  time.Duration(cf.DialTime) * time.Second,
		ReadTimeout:  time.Duration(cf.ReadTime) * time.Second,
		WriteTimeout: time.Duration(cf.WriteTime) * time.Second,
		Password:     "",
		DB:           cf.DB,
		PoolSize:     cf.PoolSize,
	})

	_, err := client.Ping().Result()
	if err != nil {
		panic(err)
	}

	return client
}
