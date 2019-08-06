package webconn

import (
	"github.com/gorilla/websocket"
)

const (
	maxConnLimit = 500
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

// Response 响应体
type Response struct {
	Errno int         `json:"errno"`
	Msg   string      `json:"msg"`
	Data  interface{} `json:"data"`
}


