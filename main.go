package main

import (
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

	router.GET("/api/Listen/contest/:contestID/user/:uID", wsListenHandler)
	router.GET("/api/contest/:contestID/contestOnline", contestOnlineNumHandler)
	return router
}

func main() {
	port := flag.String("port", "8888", "server port")
	flag.Parse()

	r := RegisterRouter()

	serverAddr := fmt.Sprintf("%s:%s", "127.0.0.1", *port)
	log.Println("WebSocket Server Listen At:", serverAddr)

	// http.Handle("/", r)
	// http.ListenAndServe(serverAddr, nil)
	go webconn.ListenkillSignal()

	log.Println(http.ListenAndServe(serverAddr, r))

}
