package web

type contestOnlineInfo struct {
	ContestID    int
	OnlineNum    int
	OnlineStatus bool
}

// Response 响应体
type Response struct {
	Err
	Data interface{} `json:"data"`
}

// Err 错误体
type Err struct {
	Msg     string `json:"msg"`
	ErrCode int    `json:"errCode"`
}

var (
	// Normal 正常
	Normal = Err{ErrCode: 000, Msg: "successful"}

	// ConnectionLimitError 连接超过上线
	ConnectionLimitError = Err{ErrCode: 001, Msg: "connection over max limit"}

	// RequsetParamsError 请求参数错误
	RequsetParamsError = Err{ErrCode: 002, Msg: "request params invalid "}

	// WebSocketError websocket 连接失败
	WebSocketError = Err{ErrCode: 003, Msg: "websocket error"}

	// SeverCacheError 系统缓存错误
	SeverCacheError = Err{ErrCode: 003, Msg: "server error"}

	// SeverDatabaseError 系统数据库错误
	SeverDatabaseError = Err{ErrCode: 004, Msg: "server error"}
)
