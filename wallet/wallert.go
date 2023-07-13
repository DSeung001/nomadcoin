package wallet

import (
	"crypto/ecdsa"
	"crypto/x509"
	"encoding/hex"
	"fmt"
	"github.com/nomadcoders/nomadcoin/utils"
	"math/big"
)

const (
	signature     string = "2a0aa305456162e2495b370daf4d89db4307f566af68608e58aa0c111dddaca776f4395999d591ff74a844cd823977add44420dc0c567d87233f4c24526c0d60"
	privateKey    string = "30770201010420d64abd6f0843486444631b19c6a6a351ae5e2f6a872ea985561a0274c15c0c3ba00a06082a8648ce3d030107a144034200048add9b337ea457e1972432b5f0b701854d93c75b9fd29d0eb81f99e79a88616f88c1ad337fa45f598eba45ee2d5a2e5ae9ceeb222a6c32cc2d114495848cb2f6"
	hashedMessage string = "1c5863cd55b5a4413fd59f054af57ba3c75c0698b3851d70f99b8de2d5c7338f"
)

func Start() {
	// 지금은 16진수인게 확실하지만 나중에 지갑으로 만들면 파일이기에 변경될 수 있는 점을 대처하기 위함
	privBytes, err := hex.DecodeString(privateKey)
	utils.HandleErr(err)

	private, err := x509.ParseECPrivateKey([]byte(privBytes))
	utils.HandleErr(err)
	fmt.Println(private)

	sigBytes, err := hex.DecodeString(signature)
	rBytes := sigBytes[:len(sigBytes)/2]
	sBytes := sigBytes[len(sigBytes)/2:]
	fmt.Printf("%d\n\n%d\n\n%d\n\n", sigBytes, rBytes, sBytes)

	var bigR, bigS = big.Int{}, big.Int{}
	bigR.SetBytes(rBytes)
	bigS.SetBytes(sBytes)
	fmt.Println(bigR, bigS)

	hashBytes, err := hex.DecodeString(hashedMessage)
	utils.HandleErr(err)

	ok := ecdsa.Verify(&private.PublicKey, hashBytes, &bigR, &bigS)
	fmt.Println(ok)
}
