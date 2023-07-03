package blockchain

import (
	"github.com/nomadcoders/nomadcoin/utils"
	"time"
)

const (
	minderReward int = 50
)

type Tx struct {
	Id        string   `json:"id"`
	Timestamp int      `json:"timestamp"`
	TxIns     []*TxIn  `json:"txIns"`
	TxOuts    []*TxOut `json:"txOuts"`
}

func (t *Tx) getId() {
	t.Id = utils.Hash(t)

}

type TxIn struct {
	Owner  string
	Amount int
}

type TxOut struct {
	Owner  string
	Amount int
}

// coinbase transaction : 블록체인 네트워크에서 발생하는 거래 내역(은행의 돈 인쇄)
func makeCoinbaseTx(address string) *Tx {
	txIns := []*TxIn{
		{"COINBASE", minderReward},
	}
	txOuts := []*TxOut{
		{address, minderReward},
	}
	tx := Tx{
		Id:        "",
		Timestamp: int(time.Now().Unix()),
		TxIns:     txIns,
		TxOuts:    txOuts,
	}
	tx.getId()
	return &tx
}
