package main

import (
	"Project/websocket/webconn"
	"flag"
	"fmt"
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
	// router.GET("/keep-listening/contest/:contestID/user/:uID", wsHandler)
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
	port := flag.String("port", "8888", "server port")
	flag.Parse()

	r := RegisterRouter()
	go webconn.ListenkillSignal()

	serverAddr := fmt.Sprintf("%s:%s", "127.0.0.1", *port)
	log.Println("WebSocket Server Listen At:", serverAddr)

	// http.Handle("/", r)
	// http.ListenAndServe(serverAddr, nil)

	log.Println(http.ListenAndServe(serverAddr, r))

}
