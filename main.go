package main

import (
	"crypto/sha256"
	"fmt"
	"strings"
)

func main() {
	// 난이도
	difficulty := 3

	// 앞에 있어야하는 문자열
	target := strings.Repeat("0", difficulty)

	// 해시키
	nonce := 1

	for {
		hash := fmt.Sprintf("%x", sha256.Sum256([]byte("hello"+fmt.Sprint(nonce))))
		fmt.Printf("Hash:%s\nTarget:%s\nNonce:%d\n\n", hash, target, nonce)
		if strings.HasPrefix(hash, target) {
			return
		} else {
			nonce++
		}

	}
}
