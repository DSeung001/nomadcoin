package blockchain

import (
	"crypto/sha256"
	"fmt"
	"sync"
)

type block struct {
	data     string
	hash     string
	prevHash string
}

type blockchain struct {
	blocks []*block
}

var b *blockchain

// Once 는 오직 한번만 실행
var once sync.Once

// calculateHash : 해시 값 계산, 포인터 변수로 직접 변경
func (b *block) calculateHash() {
	hash := sha256.Sum256([]byte(b.data + b.prevHash))
	b.hash = fmt.Sprintf("%x", hash)
}

// getLastHash : 마지막 해시 값 가져오기
func getLastHash() string {
	totalBlocks := len(GetBlockchain().blocks)
	if totalBlocks == 0 {
		return ""
	}
	return GetBlockchain().blocks[totalBlocks-1].hash
}

// createBlock : 블록 생성
func createBlock(data string) *block {
	newBlock := block{data: data, hash: "", prevHash: getLastHash()}
	newBlock.calculateHash()
	return &newBlock
}

// GetBlockchain : 블록체인 가져오기
func GetBlockchain() *blockchain {
	if b == nil {
		once.Do(func() {
			b = &blockchain{}
			b.blocks = append(b.blocks, createBlock("Genesis Block"))
		})
	}
	return b
}
