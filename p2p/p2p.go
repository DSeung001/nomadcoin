package p2p

import (
	"github.com/gorilla/websocket"
	"github.com/nomadcoders/nomadcoin/utils"
	"net/http"
)

var upgrader = websocket.Upgrader{}

func Upgrade(rw http.ResponseWriter, r *http.Request) {
	_, err := upgrader.Upgrade(rw, r, nil)
	utils.HandleErr(err)
}
