package wallet

import (
	"crypto/x509"
	"encoding/hex"
	"io/fs"
	"reflect"
	"testing"
)

const (
	testKey     string = "30770201010420da732cc84d66346a3e2d95e8e3fa450290019ce761b469e3e0257d1e6bb6b61fa00a06082a8648ce3d030107a1440342000434baec5b9b8aae693fe508e1cebf7bfa866e097ee57f777eb1a21ae87634ca126a262d5463c9f61ce7064d7e157ccb839a3ed7b516be4892a9288d44f45546fd"
	testPayload        = "00bff541d78c5a152fa7a3dbd9ac581d99bacc3c629a02dabe6347f064701964"
	testSig            = "b5c5fe738b309250817fedb45a19bbe46f772c649220252fecc26e7b5b9c5c9a779d854ab1cacc64dae873481768a66366a15e37b24b2508f6bb8814a5873173"
)

type fakeLayer struct {
	fakeHasWalletFile func() bool
}

func (f fakeLayer) hasWalletFile() bool {
	return f.fakeHasWalletFile()
}

func (fakeLayer) writeFile(name string, data []byte, perm fs.FileMode) error {
	return nil
}

func (fakeLayer) readFile(name string) ([]byte, error) {
	return x509.MarshalECPrivateKey(makeTestWallet().privateKey)
}

func TestWallet(t *testing.T) {
	t.Run("New Wallet is created", func(t *testing.T) {
		files = fakeLayer{
			// Wallet func의 hasWalletFile을 변경
			fakeHasWalletFile: func() bool {
				t.Log("I have been called")
				return false
			},
		}

		tw := Wallet()
		if reflect.TypeOf(tw) != reflect.TypeOf(&wallet{}) {
			t.Error("New Wallet should return a new wallet instance")
		}
	})

	t.Run("Wallet is restored", func(t *testing.T) {
		files = fakeLayer{
			// Wallet func의 hasWalletFile을 변경
			fakeHasWalletFile: func() bool {
				t.Log("I have been called")
				return true
			},
		}

		// w 초기화
		w = nil
		tw := Wallet()
		if reflect.TypeOf(tw) == reflect.TypeOf(&wallet{}) {
			t.Error("New Wallet should return a new wallet instance")
		}
	})
}

func makeTestWallet() *wallet {
	w := &wallet{}
	b, _ := hex.DecodeString(testKey)
	key, _ := x509.ParseECPrivateKey(b)
	w.privateKey = key
	w.Address = aFromK(key)
	return w
}

func TestSign(t *testing.T) {
	s := Sign(testPayload, makeTestWallet())
	_, err := hex.DecodeString(s)
	if err != nil {
		t.Errorf("Sign() should return a hex encoded string, got %s", s)
	}
}

func TestVerify(t *testing.T) {
	type test struct {
		input string
		ok    bool
	}

	tests := []test{
		{testPayload, true},
		{"04bff541d78c5a152fa7a3dbd9ac581d99bacc3c629a02dabe6347f064701964", false},
	}

	for _, tc := range tests {
		w := makeTestWallet()
		ok := Verify(testSig, tc.input, w.Address)
		if ok != tc.ok {
			t.Error("Verify() could not verify testSignature and testPayload")
		}
	}
}

func TestRestoreBigInts(t *testing.T) {
	_, _, err := restoreBigInts("xx")
	if err == nil {
		t.Error("restoreBigInts should return error and when payload is not hex.")
	}
}
