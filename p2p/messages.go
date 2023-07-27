package p2p

import (
	"encoding/json"
	"fmt"
	"github.com/nomadcoders/nomadcoin/blockchain"
	"github.com/nomadcoders/nomadcoin/utils"
)

type MessageKind int

const (
	// iota 를 사용하는 순간 1,2,3 로 차례대로 매핑
	MessageNewestBlock MessageKind = iota
	MessageAllBlocksRequest
	MessageAllBlocksResponse
)

type Message struct {
	Kind    MessageKind
	Payload []byte
}

func makeMessage(kind MessageKind, payload interface{}) []byte {
	m := Message{
		Kind:    kind,
		Payload: utils.ToJSON(payload),
	}
	return utils.ToJSON(m)
}

func sendNewestBlock(p *peer) {
	b, err := blockchain.FindBlock(blockchain.Blockchain().NewestHash)
	utils.HandleErr(err)
	m := makeMessage(MessageNewestBlock, b)
	p.inbox <- m
}

func handleMsg(m *Message, p *peer) {
	switch m.Kind {
	case MessageNewestBlock:
		var payload blockchain.Block
		utils.HandleErr(json.Unmarshal(m.Payload, &payload))
		fmt.Println(payload)
	default:
		fmt.Printf("Peer: %s, Sent a message with kind of: %d", p.key, m.Kind)
	}
}
