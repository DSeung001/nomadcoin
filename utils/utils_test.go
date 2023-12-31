package utils

import (
	"encoding/hex"
	"encoding/json"
	"errors"
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

func TestSplitter(t *testing.T) {
	type test struct {
		input  string
		sep    string
		index  int
		output string
	}

	tests := []test{
		{input: "0:6:0", sep: ":", index: 1, output: "6"},
		{input: "0:6:0", sep: ":", index: 10, output: ""},
		{input: "0:6:0", sep: "/", index: 0, output: "0:6:0"},
	}

	for _, tc := range tests {
		got := Splitter(tc.input, tc.sep, tc.index)
		if got != tc.output {
			t.Errorf("Expected %s and got %s", tc.output, got)
		}
	}
}

func TestHandleErr(t *testing.T) {
	oldLogFn := logFn
	defer func() {
		logFn = oldLogFn
	}()

	called := false

	// logFn 덮어씌우기
	logFn = func(v ...any) {
		called = true
	}

	err := errors.New("test")
	HandleErr(err)

	if !called {
		t.Error("HandleError should call fn")
	}
}

func TestFromBytes(t *testing.T) {
	type testStrcut struct {
		Test string
	}

	var restored testStrcut
	ts := testStrcut{"test"}
	b := ToBytes(ts)
	FromBytes(&restored, b)

	if !reflect.DeepEqual(ts, restored) {
		t.Error("FromBytes() should restore struct")
	}
}

func TestToJSON(t *testing.T) {
	type testStruct struct{ Test string }
	s := testStruct{"test"}
	b := ToJSON(s)
	k := reflect.TypeOf(b).Kind()
	if k != reflect.Slice {
		t.Errorf("Expected %v and got %v", reflect.Slice, k)
	}

	var restored testStruct
	json.Unmarshal(b, &restored)
	if !reflect.DeepEqual(s, restored) {
		t.Error("ToJson() should restore struct")
	}
}
