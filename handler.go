package main

import (
	cache "Project/websocket/dbs/cache"
	"Project/websocket/webconn"
	"net/http"
	"strconv"

	"github.com/gorilla/websocket"
	"github.com/julienschmidt/httprouter"
)

var (
	// [uid(int)][wsConn]

	upgrader = websocket.Upgrader{
		ReadBufferSize:  16,
		WriteBufferSize: 16,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
)

func wsListenHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var (
		wsConn    *websocket.Conn
		err       error
		uID       int
		contestID int
	)

	// 取出uID
	if uID, err = strconv.Atoi(p.ByName("uID")); err != nil {
		goto ERR
	}

	// 取出contestID
	if contestID, err = strconv.Atoi(p.ByName("contestID")); err != nil {
		goto ERR
	}

	if wsConn, err = upgrader.Upgrade(w, r, nil); err != nil {
		goto ERR
	}
	if err = webconn.StartConn(wsConn, uID, contestID); err != nil {
		goto ERR
	}

	return
ERR:
	resp, _ := webconn.SendResponse(0, err.Error(), "error..")
	w.Write(resp)

}

func contestOnlineNumHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var (
		num       int
		contestID int
		err       error
	)
	if contestID, err = strconv.Atoi(p.ByName("contestID")); err != nil {
		return
	}
	if num, err = cache.GetOnlineNum(contestID); err != nil {
		return
	}
	respBody := contestOnlineInfo{
		contestID: contestID,
		OnlineNum: num,
	}
	resp, _ := webconn.SendResponse(100, "successful", respBody)
	w.Write(resp)
}
