package main

import (
	"Project/websocket/webconn"
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

var (
	userConn = sync.Map{}

	upgrader = websocket.Upgrader{
		ReadBufferSize:  16,
		WriteBufferSize: 16,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
)

func wsHandler(w http.ResponseWriter, r *http.Request) {
	var (
		wsConn *websocket.Conn
		err    error
		data   []byte
		conn   *webconn.Connection
	)
	if wsConn, err = upgrader.Upgrade(w, r, nil); err != nil {
		return
	}
	// if req, err := ioutil.ReadAll(r.Body); err != nil {
	// 	return
	// }
	fmt.Println(wsConn.LocalAddr().String())
	if conn, err = webconn.NewConnection(wsConn); err != nil {
		return
	}

	go func() {
		var err error
		for {
			if err = conn.WriteMessage([]byte("heartbeat")); err != nil {
				return
			}
			time.Sleep(time.Second * 4)
		}

	}()

	for {
		if data, err = conn.ReadMessage(); err != nil {
			goto ERR
		}

		if err = conn.WriteMessage(data); err != nil {
			goto ERR
		}
	}

ERR:
	fmt.Println("close")
	conn.Close()
}

func main() {
	http.HandleFunc("/ws", wsHandler)
	http.ListenAndServe("0.0.0.0:7777", nil)
}
