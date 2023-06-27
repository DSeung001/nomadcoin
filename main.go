package main

import (
	"crypto/sha256"
	"fmt"
)

type block struct {
	data     string
	hash     string
	prevHash string
}

/*
	Go의 String은 immutable
*/

func main() {
	genesisBlock := block{"Genesis Block", "", ""}
	hash := sha256.Sum256([]byte(genesisBlock.data + genesisBlock.prevHash))

	// 16진수로
	hexHash := fmt.Sprintf("%x", hash)
	genesisBlock.hash = hexHash
	fmt.Println(genesisBlock)
}
