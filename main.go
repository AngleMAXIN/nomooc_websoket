package main

import (
	"Project/websocket/webconn"
	"log"
	"net/http"
	"strconv"
	"sync"

	"github.com/gorilla/websocket"
	"github.com/julienschmidt/httprouter"
)

var (
	// [uid(int)][wsConn]
	userWsConn = sync.Map{}

	upgrader = websocket.Upgrader{
		ReadBufferSize:  16,
		WriteBufferSize: 16,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
)

// RegisterRouter 注册路由
func RegisterRouter() *httprouter.Router {
	router := httprouter.New()
	router.GET("/api/Listen/contest/:contestID/user/:uID", wsHandler)
	return router
}

func wsHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var (
		wsConn    *websocket.Conn
		err       error
		uID       int
		contestID int
	)

	// 取出uID
	if uID, err = strconv.Atoi(p.ByName("uID")); err != nil {
		return
	}

	// 取出contestID
	if contestID, err = strconv.Atoi(p.ByName("contestID")); err != nil {
		return
	}

	if wsConn, err = upgrader.Upgrade(w, r, nil); err != nil {
		return
	}

	_ = webconn.NewConnection(wsConn, uID, contestID)
}

func main() {
	r := RegisterRouter()
	go webconn.ListenkillSignal()
	log.Println(http.ListenAndServe("127.0.0.1:8081", r))
}
