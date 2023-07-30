package utils

import (
	"encoding/hex"
	"fmt"
	"testing"
)

// go test ./... -v <= 전체 파일 테스팅

func TestHash(t *testing.T) {
	hash := "e005c1d727f7776a57a661d61a182816d8953c0432780beeae35e337830b1746"
	s := struct{ Test string }{Test: "test"}

	t.Run("Hash is always smae", func(t *testing.T) {
		x := Hash(s)

		if x != hash {
			t.Errorf("Exptcted %s, got %s", hash, x)
		}
	})

	t.Run("Hash is hex encoded", func(t *testing.T) {
		x := Hash(s)
		_, err := hex.DecodeString(x)
		if err != nil {
			t.Errorf("Hash should be hex encoded")
		}
	})
}

// Example을 시작어로 하면 GODOC에서 인식해서 예제로 출력
func ExampleHash() {
	s := struct{ Test string }{Test: "test"}
	x := Hash(s)
	fmt.Println(x)
	// Output: e005c1d727f7776a57a661d61a182816d8953c0432780beeae35e337830b1746
}
