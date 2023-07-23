package main

import (
	"fmt"
	"time"
)

// 보내기 전용 채널, 채널에 보내기 전용
func send(c chan<- int) {
	for i := range [10]int{} {
		fmt.Printf(">> sending %d <<\n", i)
		c <- i
		fmt.Printf(">> sent %d <<\n", i)
	}
	// 채널 닫기
	close(c)
}

// 받기 전용 채널, 채널에서 받기 전용
func receive(c <-chan int) {
	for {
		time.Sleep(10 * time.Second)
		// ok : 채널 open/close 여부
		// 단순 채널 일 경우는 담긴 데이터를 보내야지만 데이터를 다시 받을 수 있음 <=> 버퍼 채널은 여러개 담을 수 있음
		a, ok := <-c
		if !ok {
			fmt.Println("we are done")
			break
		}
		fmt.Printf("|| received || %d\n", a)
	}
}

// main 은 go 루틴을 기다리지 않음
func main() {
	// 버퍼 채널
	c := make(chan int, 5)
	go send(c)
	receive(c)

}
