package utils

import (
	"encoding/hex"
	"fmt"
	"reflect"
	"testing"
)

// go test ./... -v <= 전체 파일 테스팅
// go test -v -coverprofile cover.out ./... <= 테스팅 결고 파일로 생성
// go tool cover -html="cover.out" 로 파일이 아닌 html 확인이 가능

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

func TestToBytes(t *testing.T) {
	s := "test"
	b := ToBytes(s)
	k := reflect.TypeOf(b).Kind()
	if k != reflect.Slice {
		t.Errorf("ToBytes should return a slice of bytes %s", k)
	}
}
