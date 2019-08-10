package webconn

import (
	prcf "Project/websocket/config"

	"github.com/gorilla/websocket"
)

var (
	// MaxConnNum 最大的连接数
	wsConnBucket chan notify
	killChan     chan *userOne
	cf           prcf.ServerConfig
)

type notify struct{}

type userOne struct {
	uID       int
	contestID int
}

// CustomerConnection 自定义连接结构体
type CustomerConnection struct {
	wsConn *websocket.Conn
	user   userOne
}
