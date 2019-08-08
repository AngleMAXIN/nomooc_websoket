package cache

import (
	prcf "Project/websocket/config"
	"fmt"
	"log"
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
		Password:     "",    // no password set
		DB:           cf.DB, // use default DB
		PoolSize:     cf.PoolSize,
	})

	pong, err := client.Ping().Result()
	if err != nil {
		panic(err)
	}
	log.Println(pong)

	return client
}
