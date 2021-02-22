package main

import (
	"fmt"
	"html/template"
	"im/ctrl"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/user/login", ctrl.UserLogin)
	http.HandleFunc("/user/register", ctrl.UserRegiter)
	// http.HandleFunc("/user/find", ctrl.UserFind)
	http.HandleFunc("/contact/loadcommunity", ctrl.LoadCommunity)
	http.HandleFunc("/contact/loadfriend", ctrl.LoadFriend)
	http.HandleFunc("/contact/joincommunity", ctrl.JoinCommunity)
	http.HandleFunc("/contact/createcommunity", ctrl.CreateCommunity)
	http.HandleFunc("/contact/addfriend", ctrl.AddFriend)
	http.HandleFunc("/chat", ctrl.Chat)
	http.HandleFunc("/attach/upload", ctrl.Upload)

	http.Handle("/asset/", http.FileServer(http.Dir(".")))
	http.Handle("/mnt/", http.FileServer(http.Dir(".")))
	RegisterView()
	http.ListenAndServe(":8080", nil)
}

func RegisterView() {
	tpl, err := template.ParseGlob("view/**/*")
	if err != nil {
		log.Fatal(err.Error())
	}
	// todo:此处有重复
	for _, v := range tpl.Templates() {
		tplname := v.Name()
		fmt.Println("Register template: ", tplname)
		http.HandleFunc(tplname, func(w http.ResponseWriter, r *http.Request) {
			err1 := tpl.ExecuteTemplate(w, tplname, nil)
			if err1 != nil {
				log.Fatal(err.Error())
			}
		})
	}

}
