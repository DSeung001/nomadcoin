package blockchain

import (
	"crypto/sha256"
	"fmt"
	"sync"
)

type Block struct {
	Data     string
	Hash     string
	PrevHash string
}

type blockchain struct {
	blocks []*Block
}

var b *blockchain

// Once 는 오직 한번만 실행
var once sync.Once

// calculateHash : 해시 값 계산, 포인터 변수로 직접 변경
func (b *Block) calculateHash() {
	hash := sha256.Sum256([]byte(b.Data + b.PrevHash))
	b.Hash = fmt.Sprintf("%x", hash)
}

// getLastHash : 마지막 해시 값 가져오기
func getLastHash() string {
	totalBlocks := len(GetBlockchain().blocks)
	if totalBlocks == 0 {
		return ""
	}
	return GetBlockchain().blocks[totalBlocks-1].Hash
}

// createBlock : 블록 생성
func createBlock(data string) *Block {
	newBlock := Block{Data: data, Hash: "", PrevHash: getLastHash()}
	newBlock.calculateHash()
	return &newBlock
}

// AddBlock : 블록 추가
func (b *blockchain) AddBlock(data string) {
	b.blocks = append(b.blocks, createBlock(data))
}

// GetBlockchain : 블록체인 가져오기
func GetBlockchain() *blockchain {
	if b == nil {
		once.Do(func() {
			b = &blockchain{}
			b.AddBlock("Genesis")
		})
	}
	return b
}

func (b *blockchain) AllBlocks() []*Block {
	return b.blocks
}
