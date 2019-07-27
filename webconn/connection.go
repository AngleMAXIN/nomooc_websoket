package webconn

import (
	"errors"
	"fmt"
	"sync"

	"github.com/gorilla/websocket"
)

type Connection struct {
	wsConn  *websocket.Conn
	inChan  chan []byte
	outChan chan []byte

	isClosed  bool
	mutex     sync.Mutex
	closeChan chan byte
}

func NewConnection(wsConn *websocket.Conn) (conn *Connection, err error) {
	conn = &Connection{
		wsConn:    wsConn,
		inChan:    make(chan []byte, 1000),
		outChan:   make(chan []byte, 1000),
		closeChan: make(chan byte, 1),
	}

	// 启动读协程
	go conn.readLoop()

	// 自动写协程
	go conn.writeLoop()

	return
}

// ReadMessage 读取来自Client的数据,从channel中
func (conn *Connection) ReadMessage() (data []byte, err error) {
	select {
	case data = <-conn.inChan:
	case <-conn.closeChan:
		err = errors.New("ReadMessage connection is closed")
		fmt.Println(err)
	}

	return
}

// WriteMessage 写进来自Client的数据,从channel中
func (conn *Connection) WriteMessage(data []byte) (err error) {
	select {
	case conn.outChan <- data:
	case <-conn.closeChan:
		err = errors.New("WriteMessage connection is closed")
		fmt.Println(err)

	}
	return
}

func (conn *Connection) Close() {
	//线程安全
	conn.wsConn.Close()
	fmt.Println("satrt close")
	// 这里只会执行一次
	conn.mutex.Lock()
	if !conn.isClosed {
		fmt.Println("real close")
		close(conn.closeChan)
		conn.isClosed = true
	}
	conn.mutex.Unlock()
	fmt.Println("----close over----\n")

}

func (conn *Connection) readLoop() {
	var (
		data []byte
		err  error
	)

	for {
		if _, data, err = conn.wsConn.ReadMessage(); err != nil {
			fmt.Println("readLoop", err)
			goto ERR
		}

		select {
		case conn.inChan <- data:
		case <-conn.closeChan:
			goto ERR

		}
	}
ERR:
	fmt.Println("---resdLoop")
	conn.Close()
}

func (conn *Connection) writeLoop() {
	var (
		data []byte
		err  error
	)

	for {

		select {
		case data = <-conn.outChan:
		case <-conn.closeChan:
			goto ERR
		}
		if err = conn.wsConn.WriteMessage(websocket.TextMessage, data); err != nil {
			fmt.Println("writeLoop", err)
			goto ERR
		}
	}
ERR:
	fmt.Println("---writeLoop")

	conn.Close()
}
