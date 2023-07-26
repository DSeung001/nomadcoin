package p2p

import (
	"fmt"
	"github.com/gorilla/websocket"
)

var Peers map[string]*peer = make(map[string]*peer)

type peer struct {
	key     string
	address string
	port    string
	conn    *websocket.Conn
	inbox   chan []byte
}

// Receiver Method
func (p *peer) close() {
	p.conn.Close()
	delete(Peers, p.key)
}

func (p *peer) read() {
	// 에러 발생 시 메서드가 종료되니 defer 사용
	defer p.close()
	for {
		fmt.Println("read blocking...")
		_, m, err := p.conn.ReadMessage()
		if err != nil {
			break
		}
		fmt.Printf("[read] %s\n", m)
	}
}

func (p *peer) write() {
	defer p.close()
	for {
		m, ok := <-p.inbox
		if !ok {
			break
		}
		p.conn.WriteMessage(websocket.TextMessage, m)
	}
}

func initPeer(conn *websocket.Conn, address, port string) *peer {
	key := fmt.Sprintf("%s:%s", address, port)
	p := &peer{
		conn:    conn,
		inbox:   make(chan []byte),
		address: address,
		key:     key,
		port:    port,
	}
	go p.read()
	go p.write()
	Peers[key] = p
	return p
}
