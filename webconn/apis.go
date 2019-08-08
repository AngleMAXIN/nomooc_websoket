package webconn

import (
	cache "Project/websocket/dbs/cache"
	db "Project/websocket/dbs/database"
	"errors"
	"log"
	"time"

	"fmt"

	"github.com/gorilla/websocket"
)

var (
	// MaxConnNum 最大的连接数
	wsConnBucket chan notify
	killChan     chan *userOne
)

// InitMaxConnNum 初始化最大连接数
func InitMaxConnNum(num int) (err error) {
	if num > maxConnLimit {
		err = fmt.Errorf("wsConnBucket number not more than %d", maxConnLimit)
		return
	}
	wsConnBucket = make(chan notify, num)
	return
}

// ListenkillSignal 监听conn断开信号，并更新用户状态
func ListenkillSignal() {
	killChan = make(chan *userOne, 5)

	for v := range killChan {
		go db.MarkUserStatus(v.uID, v.contestID)
		go cache.DecrOnlineNum(v.contestID)

		<-wsConnBucket
	}
}

// StartConn 封装一层实现，做并发连接限制
func StartConn(wsConn *websocket.Conn, uID, contestID int) (err error) {

	ticker := time.NewTicker(3 * time.Second)

	select {
	case wsConnBucket <- notify{}:
		go newConnection(wsConn, uID, contestID)
		// 增加一个在线人数
		go cache.IncrOnlineNum(contestID)
	case <-ticker.C:
		return errors.New("connection has reached the limit, please reconnect")
	}
	return

}

// NewConnection 创建一个Conn
func newConnection(wsConn *websocket.Conn, uID, contestID int) (conn *CustomerConnection) {

	conn = &CustomerConnection{
		wsConn: wsConn,
		user: userOne{
			uID:       uID,
			contestID: contestID,
		},
	}
	log.Printf("Connection start listen: user %d, contest %d\n", uID, contestID)
	// 启动读协程
	go conn.ListenLoop()

	return
}

// ListenLoop 监听连接
func (conn *CustomerConnection) ListenLoop() {
	var (
		err error
	)

	for {
		if _, _, err = conn.wsConn.ReadMessage(); err != nil {
			// 用户突然断开链接
			if websocket.IsCloseError(err, websocket.CloseGoingAway) {
				log.Printf("Connection Close: user %d, contest %d\n", conn.user.uID, conn.user.contestID)

				killChan <- &conn.user
			}
			conn.wsConn.Close()
			return
		}
	}
}
