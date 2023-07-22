package main

import (
	"fmt"
	"time"
)

func countToTen(name string) {
	for i := range [10]int{} {
		fmt.Println(name, ": ", i)
		time.Sleep(1 * time.Second)
	}
}

// main 은 go 루틴을 기다리지 않음
func main() {
	go countToTen("first")
	go countToTen("second")
	go countToTen("third")

	for {

	}
}
