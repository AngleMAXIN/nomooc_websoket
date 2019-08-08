package webconn

import (
	prcf "Project/websocket/config"

	"github.com/gorilla/websocket"
)

var (
	maxConnLimit int
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

func init() {
	maxConnLimit = prcf.ProConfig.GetServerConfig().MaxConnLimit
}
