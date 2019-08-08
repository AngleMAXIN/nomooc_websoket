package web

import (
	"encoding/json"
	"net/http"
)

// SendResponse 应答方法
func SendResponse(w http.ResponseWriter, err *Err, data interface{}) {

	response := Response{
		Data: data,
	}
	response.Err = *err
	resp, _ := json.Marshal(response)
	w.Write(resp)
}
