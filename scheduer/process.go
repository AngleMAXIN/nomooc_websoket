package scheduer

import "time"

// ListenContestEnd 监控竞赛考试一旦结束
func ListenContestEnd(contestID int, endTime time.Time) (err error) {
	return
}

// ClearContestAll 监控竞赛考试一旦结束，清空缓存中的数据以及关闭所有有关的长连接
func ClearContestAll(contestID int) (err error) {
	return
}
