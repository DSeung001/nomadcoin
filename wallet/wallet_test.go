package wallet

import (
	"crypto/x509"
	"encoding/hex"
	"testing"
)

const (
	testKey     string = "30770201010420da732cc84d66346a3e2d95e8e3fa450290019ce761b469e3e0257d1e6bb6b61fa00a06082a8648ce3d030107a1440342000434baec5b9b8aae693fe508e1cebf7bfa866e097ee57f777eb1a21ae87634ca126a262d5463c9f61ce7064d7e157ccb839a3ed7b516be4892a9288d44f45546fd"
	testPayload        = "00bff541d78c5a152fa7a3dbd9ac581d99bacc3c629a02dabe6347f064701964"
	testSig            = "b5c5fe738b309250817fedb45a19bbe46f772c649220252fecc26e7b5b9c5c9a779d854ab1cacc64dae873481768a66366a15e37b24b2508f6bb8814a5873173"
)

func makeTestWallet() *wallet {
	w := &wallet{}
	b, _ := hex.DecodeString(testKey)
	key, _ := x509.ParseECPrivateKey(b)
	w.privateKey = key
	w.Address = aFromK(key)
	return w
}

func TestVerify(t *testing.T) {
	s := Sign(testPayload, makeTestWallet())
	_, err := hex.DecodeString(s)
	if err != nil {
		t.Errorf("Sign() should return a hex encoded string, got %s", s)
	}
}
