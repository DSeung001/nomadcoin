package wallet

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"github.com/nomadcoders/nomadcoin/utils"
)

func Start() {

	// 키 생성
	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	utils.HandleErr(err)
	fmt.Println("Private Key", privateKey.D)
	fmt.Println("Public Key, x, y", privateKey.X, privateKey.Y)

	// 해싱
	message := "i love you"
	hashedMessage := utils.Hash(message)
	hashAsBytes, err := hex.DecodeString(hashedMessage)
	utils.HandleErr(err)

	// 서명
	r, s, err := ecdsa.Sign(rand.Reader, privateKey, hashAsBytes)
	utils.HandleErr(err)
	fmt.Printf("R:%d\nS:%d", r, s)
}
