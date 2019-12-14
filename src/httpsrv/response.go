package httpsrv

import (
	"encoding/json"
	"net/http"
)

// Success 成功信息。
type Success struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
}

// Failure 失败信息。
type Failure struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func SendSuccess(w http.ResponseWriter, r *http.Request, data interface{}) {
	sendResponse(w, r, 200, Success{Code: 0, Data: data})
}

func SendFailure(w http.ResponseWriter, r *http.Request, httpcode int, failure Failure) {
	if failure.Code == 0 {
		failure.Code = httpcode
	}
	sendResponse(w, r, httpcode, failure)
}

func sendResponse(w http.ResponseWriter, r *http.Request, httpcode int, payload interface{}) {
	response, _ := json.Marshal(payload)
	responseType := r.Header.Get("Accept")

	if responseType == "" || responseType == "*/*" {
		responseType = "application/json"
	}

	w.Header().Set("Content-Type", responseType)
	w.WriteHeader(httpcode)
	w.Write(response)
}
