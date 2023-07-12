package wallet

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"encoding/hex"
	"fmt"
	"github.com/nomadcoders/nomadcoin/utils"
)

const (
	signature     string = "2a0aa305456162e2495b370daf4d89db4307f566af68608e58aa0c111dddaca776f4395999d591ff74a844cd823977add44420dc0c567d87233f4c24526c0d60"
	privateKey    string = "30770201010420d64abd6f0843486444631b19c6a6a351ae5e2f6a872ea985561a0274c15c0c3ba00a06082a8648ce3d030107a144034200048add9b337ea457e1972432b5f0b701854d93c75b9fd29d0eb81f99e79a88616f88c1ad337fa45f598eba45ee2d5a2e5ae9ceeb222a6c32cc2d114495848cb2f6"
	hashedMessage string = "1c5863cd55b5a4413fd59f054af57ba3c75c0698b3851d70f99b8de2d5c7338f"
)

// 일반적으로 아래 내용은 다 다른대서 오지만 공부를 위해 모아둔 것
// r, s : 지갑의 서명값
// privateKey : 파일
// 해시메시지 : 트랜잭션

func Start() {

	// 키 생성
	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	keyAsBytes, err := x509.MarshalECPrivateKey(privateKey)
	utils.HandleErr(err)
	fmt.Println("Private Key", privateKey.D)
	fmt.Println("Public Key, x, y", privateKey.X, privateKey.Y)
	fmt.Printf("keyAsBytes(prviate key) : %x\n", keyAsBytes)

	// 해싱
	message := "i love you"
	hashedMessage := utils.Hash(message)
	fmt.Println(hashedMessage)
	hashAsBytes, err := hex.DecodeString(hashedMessage)
	utils.HandleErr(err)

	// 서명
	r, s, err := ecdsa.Sign(rand.Reader, privateKey, hashAsBytes)
	utils.HandleErr(err)
	fmt.Printf("R:%d\nS:%d\n", r, s)
	fmt.Println(r.Bytes(), s.Bytes())
	fmt.Println(len(r.Bytes()), len(s.Bytes()))
	signature := append(r.Bytes(), s.Bytes()...)
	fmt.Printf("%x\n", signature)

	// 인증
	ok := ecdsa.Verify(&privateKey.PublicKey, hashAsBytes, r, s)
	fmt.Println(ok)
}
