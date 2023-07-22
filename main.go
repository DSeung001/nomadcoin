package main

import (
	"fmt"
	"time"
)

// 보내기 전용 채널, 채널에 보내기 전용
func countToTen(c chan<- int) {
	for i := range [10]int{} {
		time.Sleep(1 * time.Second)
		fmt.Printf("sending %d\n", i)
		c <- i
	}
	// 채널 닫기
	close(c)
}

// 받기 전용 채널, 채널에서 받기 전용
func receive(c <-chan int) {
	for {
		// ok : 채널 open/close 여부
		a, ok := <-c
		if !ok {
			fmt.Println("we are done")
			break
		}
		fmt.Printf("received %d\n", a)
	}
}

// main 은 go 루틴을 기다리지 않음
func main() {
	c := make(chan int)
	go countToTen(c)
	receive(c)

}
