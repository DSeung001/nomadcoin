package blockchain

import (
	"errors"
	"github.com/nomadcoders/nomadcoin/utils"
	"time"
)

const (
	minderReward int = 50
)

type mempool struct {
	Txs []*Tx
}

var Mempool *mempool = &mempool{}

type Tx struct {
	ID        string   `json:"id"`
	Timestamp int      `json:"timestamp"`
	TxIns     []*TxIn  `json:"txIns"`
	TxOuts    []*TxOut `json:"txOuts"`
}

func (t *Tx) getId() {
	t.ID = utils.Hash(t)
}

// TxIn : TxID는 Output을 생성한 트랜잭션을 Index는 위치를
type TxIn struct {
	TxID  string `json:"txID"`
	Index int    `json:"index"`
	Owner string `json:"owner"`
}

type TxOut struct {
	Owner  string `json:"owner"`
	Amount int    `json:"amount"`
}

// UTxOut : 사용자가 아직 사용하지 않는 transaction output
type UTxOut struct {
	TxID   string `json:"txID"`
	Index  int    `json:"index"`
	Amount int    `json:"amount"`
}

// coinbase transaction : 블록체인 네트워크에서 발생하는 거래 내역(은행의 돈 인쇄)
func makeCoinbaseTx(address string) *Tx {
	txIns := []*TxIn{
		{"", -1, "COINBASE"},
	}
	txOuts := []*TxOut{
		{address, minderReward},
	}
	tx := Tx{
		ID:        "",
		Timestamp: int(time.Now().Unix()),
		TxIns:     txIns,
		TxOuts:    txOuts,
	}
	tx.getId()
	return &tx
}

func makeTx(from, to string, amount int) (*Tx, error) {
	if Blockchain().BalanceByAddress(from) < amount {
		return nil, errors.New("not enough 돈")
	}
	var txOuts []*TxOut
	var txIns []*TxIn
	total := 0
	UTxOuts := Blockchain().UTxOutsByAddress(from)
	for _, UTxOut := range UTxOuts {
		if total > amount {
			break
		}
		txIn := &TxIn{UTxOut.TxID, UTxOut.Index, from}
		txIns = append(txIns, txIn)
		total += UTxOut.Amount
	}
	if change := total - amount; change != 0 {
		changeTxOut := &TxOut{from, change}
		txOuts = append(txOuts, changeTxOut)
	}
	txOut := &TxOut{to, amount}
	txOuts = append(txOuts, txOut)
	tx := &Tx{
		ID:        "",
		Timestamp: int(time.Now().Unix()),
		TxIns:     txIns,
		TxOuts:    txOuts,
	}
	tx.getId()
	return tx, nil
}

func (m *mempool) AddTx(to string, amount int) error {
	tx, err := makeTx("nico", to, amount)
	if err != nil {
		return err
	}
	m.Txs = append(m.Txs, tx)
	return nil
}

func (m *mempool) TxToConfirm() []*Tx {
	coinbase := makeCoinbaseTx("nico")
	txs := m.Txs
	txs = append(txs, coinbase)
	m.Txs = nil
	return txs
}
