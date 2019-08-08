package web

import (
	cache "Project/websocket/dbs/cache"
	"Project/websocket/webconn"
	"net/http"
	"strconv"

	"github.com/gorilla/websocket"
	"github.com/julienschmidt/httprouter"
)

var (
	upgrader = websocket.Upgrader{
		ReadBufferSize:  16,
		WriteBufferSize: 16,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
)

// WsListenHandler 建立websocket连接
func WsListenHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var (
		wsConn    *websocket.Conn
		errInfo   *Err
		err       error
		uID       int
		contestID int
	)

	// 取出uID
	if uID, err = strconv.Atoi(p.ByName("uID")); err != nil {
		errInfo = &RequsetParamsError
		goto ERR
	}

	// 取出contestID
	if contestID, err = strconv.Atoi(p.ByName("contestID")); err != nil {
		errInfo = &RequsetParamsError
		goto ERR
	}

	if wsConn, err = upgrader.Upgrade(w, r, nil); err != nil {
		errInfo = &WebSocketError
		goto ERR
	}

	if err = webconn.StartConn(wsConn, uID, contestID); err != nil {
		errInfo = &ConnectionLimitError
		goto ERR
	}

ERR:
	SendResponse(w, errInfo, nil)

}

// ContestOnlineNumHandler 获得竞赛实时在线人数
func ContestOnlineNumHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var (
		num       int
		contestID int
		err       error
		status    = true
	)
	if contestID, err = strconv.Atoi(p.ByName("contestID")); err != nil {
		SendResponse(w, &RequsetParamsError, nil)
		return
	}
	if num, err = cache.GetOnlineNum(contestID); err != nil {
		num = 0
		status = false
	}
	respBody := contestOnlineInfo{
		ContestID:    contestID,
		OnlineNum:    num,
		OnlineStatus: status,
	}
	SendResponse(w, &Normal, respBody)
	return
}
