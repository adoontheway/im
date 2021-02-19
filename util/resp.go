package util

import (
	"encoding/json"
	"log"
	"net/http"
)

type H struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

func Resp(w http.ResponseWriter, code int, data interface{}, msg string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	// define the struct
	h := H{
		Code: code,
		Msg:  msg,
		Data: data,
	}
	// serialize the struct to json
	ret, err := json.Marshal(h)
	if err != nil {
		log.Println(err.Error())
	}
	// output
	w.Write([]byte(ret))
}

func RespFail(w http.ResponseWriter, msg string) {
	Resp(w, -1, nil, msg)
}

func RespOk(w http.ResponseWriter, data interface{}, msg string) {
	Resp(w, 0, data, "")
}
