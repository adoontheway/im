package ctrl

import (
	"fmt"
	"im/model"
	"im/service"
	"im/util"
	"math/rand"
	"net/http"
)

var userService service.UserService

func UserLogin(w http.ResponseWriter, r *http.Request) {
	// io.WriteString(writer, "hello world")
	r.ParseForm()
	mobile := r.PostForm.Get("mobile")
	passwd := r.PostForm.Get("passwd")

	// var str string
	user, err := userService.Login(mobile, passwd)
	if err != nil {
		util.RespFail(w, err.Error())
	} else {
		util.RespOk(w, user, "")
	}

}

func UserRegiter(w http.ResponseWriter, r *http.Request) {
	// io.WriteString(writer, "hello world")
	r.ParseForm()
	mobile := r.PostForm.Get("mobile")
	plainpwd := r.PostForm.Get("passwd")
	nickname := fmt.Sprintf("user%06d", rand.Int31n(100000))
	avatar := ""
	sex := model.SEX_UNKNOW

	user, err := userService.Register(mobile, plainpwd, nickname, avatar, sex)
	if err != nil {
		util.RespFail(w, err.Error())
	} else {
		util.RespOk(w, user, "")
	}
}
