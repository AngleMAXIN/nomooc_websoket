package main

import (
	// "Project/websocket/config"
	"Project/websocket/web"
	"Project/websocket/webconn"
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// RegisterRouter 注册路由
func RegisterRouter() *httprouter.Router {
	router := httprouter.New()

	router.GET("/contest-listen/contest/:contestID/user/:uID", web.WsListenHandler)
	router.GET("/contest-listen/contest/:contestID/contestOnline", web.ContestOnlineNumHandler)
	return router
}

func main() {

	port := flag.String("port", "8888", "server port")
	flag.Parse()

	// config.InitConfig()

	go webconn.ListenkillSignal()

	serverAddr := fmt.Sprintf("%s:%s", "127.0.0.1", *port)
	log.Println("WebSocket Server Listen At:", serverAddr)

	r := RegisterRouter()
	log.Println(http.ListenAndServe(serverAddr, r))

}
