package ctrl

import (
	"im/args"
	"im/service"
	"im/util"
	"net/http"
)

var contactService service.ContactService

func AddFriend(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	var arg args.ContactArg
	util.Bind(r, &arg)
	err := contactService.AddFriend(arg.Userid, arg.Dstid)
	if err != nil {
		util.RespFail(w, err.Error())
	} else {
		util.RespOk(w, nil, "好友添加成功")
	}
}

func JoinCommunity(w http.ResponseWriter, r *http.Request) {
	var arg args.ContactArg
	util.Bind(r, &arg)
}

func LoadFriend(w http.ResponseWriter, r *http.Request) {
	var arg args.ContactArg
	util.Bind(r, &arg)
	users := contactService.SearchFriend(arg.Userid)
	util.RespOkList(w, users, len(users))
}
func LoadCommunity(w http.ResponseWriter, r *http.Request) {
	var arg args.ContactArg
	util.Bind(r, &arg)
	communitys := contactService.SearchCommunity(arg.Userid)
	util.RespOkList(w, communitys, len(communitys))
}
