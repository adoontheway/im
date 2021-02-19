package ctrl

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"

	"github.com/gorilla/websocket"
	"gopkg.in/fatih/set.v0"
)

type Node struct {
	Conn *websocket.Conn
	// 并行转串行
	DataQueue chan []byte
	GroupSets set.Interface
}

var clientMap map[int64]*Node = make(map[int64]*Node)
var rwlocker sync.RWMutex

func Chat(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	id := query.Get("id")
	token := query.Get("token")
	userId, _ := strconv.ParseInt(id, 10, 64)
	isvalid := checkToken(userId, token)
	conn, err := (&websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return isvalid
		},
	}).Upgrade(w, r, nil)
	if err != nil {
		log.Println(err.Error())
		return
	}
	node := &Node{
		Conn:      conn,
		DataQueue: make(chan []byte, 50),
		GroupSets: set.New(set.ThreadSafe),
	}
	rwlocker.Lock()
	clientMap[userId] = node
	rwlocker.Unlock()
	go sendproc(node)
	go recvproc(node)
	sendMsg(userId, []byte("hello there!~"))
}

func checkToken(userid int64, token string) bool {
	//从数据库里面查询并比对
	user := userService.Find(userid)
	return user.Token == token
}

func sendproc(node *Node) {
	for {
		select {
		case data := <-node.DataQueue:
			err := node.Conn.WriteMessage(websocket.TextMessage, data)
			if err != nil {
				log.Println(err.Error())
				return
			}
		}
	}
}

func recvproc(node *Node) {
	for {
		_, data, err := node.Conn.ReadMessage()
		if err != nil {
			log.Println(err.Error())
			return
		}
		//todo 对data进一步处理
		fmt.Printf("Receiv <= %s \n", data)
	}
}

func sendMsg(userId int64, msg []byte) {
	//
	rwlocker.RLock()
	node, ok := clientMap[userId]
	rwlocker.RUnlock()
	if ok {
		node.DataQueue <- msg
	}
}
