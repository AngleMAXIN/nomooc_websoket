package cache

import "github.com/go-redis/redis"

const (
	keyPrefix = "contest:"
	keySuffix = ":onlineNum"
)

var (
	redisDao *redis.Client
)
