package main

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/user/login", login)
	// http.HandleFunc("/", index)
	// specify fileserver
	// http.Handle("/", http.FileServer(http.Dir(".")))
	// specify static fileserver
	http.Handle("/asset/", http.FileServer(http.Dir(".")))
	// user/login.shtml
	http.HandleFunc("/user/login.shtml", func(w http.ResponseWriter, r *http.Request) {
		tpl, err := template.ParseFiles("view/user/login.html")
		if err != nil {
			log.Fatal(err.Error())
		}
		tpl.ExecuteTemplate(w, "user/login.shtml", nil)
	})
	http.ListenAndServe(":8080", nil)
}

func login(w http.ResponseWriter, r *http.Request) {
	// io.WriteString(writer, "hello world")
	r.ParseForm()
	mobile := r.PostForm.Get("mobile")
	passwd := r.PostForm.Get("passwd")

	// var str string

	if mobile == "123456" && passwd == "123456" {
		// str = `{"code":0,"data":{"id":1,"token":"test"}}`
		data := make(map[string]interface{})
		data["id"] = 1
		data["token"] = "test"
		Resp(w, 0, data, "")
	} else {
		// str = `{"code":-1,"msg":"密码不正确"}`
		Resp(w, -1, nil, "密码不正确")
	}
	// w.Header().Set("Content-Type", "application/json")
	// w.WriteHeader(http.StatusOK)
	// w.Write([]byte(str))

}

// func index(w http.ResponseWriter, r *http.Request) {
// 	io.WriteString(w, "Hello world.")
// }

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
