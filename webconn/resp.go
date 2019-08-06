package webconn

import "encoding/json"

// SendResponse 应答方法
func SendResponse(errno int, msg string, data interface{}) (resp []byte, err error) {
	var (
		response Response
	)
	response = Response{
		Data:  data,
		Msg:   msg,
		Errno: errno,
	}
	resp, _ = json.Marshal(response)
	return
}
