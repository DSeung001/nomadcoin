package test

import (
	"testing"
)

// *_test.go의 Test* 함수들은 go test 명령어 입력시 실행됨

func TestHello(t *testing.T) {
	want := "Hello, world"
	//fmt.Println(want)
	if got := Hello(); got != want {
		t.Errorf("Hello() = %q, want %q", got, want)
	}
}

func Hello() string {
	return "Hello, world"
}
