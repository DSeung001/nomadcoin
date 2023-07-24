package p2p

import (
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/nomadcoders/nomadcoin/utils"
	"net/http"
	"strings"
)

var upgrader = websocket.Upgrader{}

// port 4000, 3000이 서로 요청을 주고 받는 걸 만듦

func Upgrade(rw http.ResponseWriter, r *http.Request) {
	// Port :3000 will upgrade the reqeust form :4000
	upgrader.CheckOrigin = func(r *http.Request) bool {
		return true
	}
	conn, err := upgrader.Upgrade(rw, r, nil)
	utils.HandleErr(err)
	openPort := r.URL.Query().Get("openPort")
	fmt.Println(r.RemoteAddr)

	result := strings.Split(r.RemoteAddr, ":")
	initPeer(conn, result[0], openPort)
}

func AddPeer(address, port, openPort string) {
	// Port :4000 is requesting an upgrade from the port :3000
	// Go에서 websocket 요청보내기
	conn, _, err := websocket.DefaultDialer.Dial(fmt.Sprintf("ws://%s:%s/ws?openPort=%s", address, port, openPort), nil)
	utils.HandleErr(err)
	initPeer(conn, address, port)
}
