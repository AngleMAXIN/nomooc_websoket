package cache

import (
	"fmt"
	"strconv"
	"time"

	"github.com/go-redis/redis"
)

const (
	keyPrefix = "contest:"
	keySuffix = ":onlineNum"
)

var (
	redisDao *redis.Client
)

func init() {
	redisDao = GetRedisConn()
}

// GetRedisConn 获得redis conn
func GetRedisConn() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:         "localhost:6379",
		DialTimeout:  10 * time.Second,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		Password:     "", // no password set
		DB:           2,  // use default DB
		PoolSize:     5,
	})
	return client
}

// IncrOnlineNum 增加一个线人员数量
func IncrOnlineNum(contestID int) (err error) {

	err = redisDao.Incr(fmt.Sprintf("%s%d%s", keyPrefix, contestID, keySuffix)).Err()
	if err != nil {
		return
	}
	return
}

// DecrOnlineNum 减少一个在线数量
func DecrOnlineNum(contestID int) (err error) {

	err = redisDao.Decr(fmt.Sprintf("%s%d%s", keyPrefix, contestID, keySuffix)).Err()
	if err != nil {
		return
	}
	return
}

// GetOnlineNum 获取当前在线人数
func GetOnlineNum(contestID int) (num int, err error) {
	var resVal string
	resVal, err = redisDao.Get(fmt.Sprintf("%s%d%s", keyPrefix, contestID, keySuffix)).Result()
	if err != nil {
		return 0, err
	}
	num, _ = strconv.Atoi(resVal)
	return
}

// Output: PONG <nil>
// intCmd := client.Incr("key")
// _, err := intCmd.Result()
// if err != nil {
// 	return
// }
// fmt.Println(client.PoolStats().TotalConns, client.PoolStats().StaleConns)
// fmt.Println(intCmd.String())
// fmt.Println(intCmd.Val())
// fmt.Println(intCmd.Args())
// client.Close()
// fmt.Println(intCmd.Name())
// wg := sync.WaitGroup{}
// wg.Add(3)

// for i := 0; i < 3; i++ {
// 	go func() {
// 		defer wg.Done()

// 		for j := 0; j < 100; j++ {
// 			client.Incr("key").Err()
// 		}

// 		fmt.Printf("PoolStats, TotalConns: %d, FreeConns: %d\n", client.PoolStats().TotalConns, client.PoolStats().StaleConns)
// 	}()
// }

// wg.Wait()

// }
