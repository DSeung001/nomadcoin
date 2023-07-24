package p2p

import (
	"github.com/gorilla/websocket"
	"github.com/nomadcoders/nomadcoin/utils"
	"net/http"
)

var conns []*websocket.Conn
var upgrader = websocket.Upgrader{}

func Upgrade(rw http.ResponseWriter, r *http.Request) {
	upgrader.CheckOrigin = func(r *http.Request) bool {
		return true
	}
	conn, err := upgrader.Upgrade(rw, r, nil)
	conns = append(conns, conn)
	utils.HandleErr(err)

	// ReadMessage 는 하나만 받고 blocking
	for {
		_, p, err := conn.ReadMessage()
		if err != nil {
			conn.Close()
			break
		}

		// 본인 Connection 에는 메세지를 쏘지 않음
		for _, aConn := range conns {
			if aConn != conn {
				utils.Hash(aConn.WriteMessage(websocket.TextMessage, p))
			}
		}
	}

}
