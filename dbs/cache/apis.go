package cache

import (
	"fmt"
	"strconv"
	"time"
)

// IncrOnlineNum 增加一个线人员数量
func IncrOnlineNum(contestID int) (err error) {
	var (
		num int64
		key string
	)
	key = fmt.Sprintf("%s%d%s", keyPrefix, contestID, keySuffix)
	num, err = redisDao.Incr(fmt.Sprintf("%s%d%s", keyPrefix, contestID, keySuffix)).Result()
	if err == nil {
		if num > 1000 {
			_ = redisDao.Set(key, num-1, time.Hour*24).Err()
		}
	}
	return
}

// DecrOnlineNum 减少一个在线数量
func DecrOnlineNum(contestID int) (err error) {
	var (
		num int64
		key string
	)

	num, err = redisDao.Decr(key).Result()
	if err == nil {
		if num == 0 {
			_ = redisDao.Set(key, 0, time.Hour*24).Err()
		}
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
