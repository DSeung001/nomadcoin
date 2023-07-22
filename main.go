package main

import (
	"fmt"
	"time"
)

func countToTen(c chan int) {
	for i := range [10]int{} {
		time.Sleep(1 * time.Second)
		fmt.Printf("sending %d\n", i)
		c <- i
	}
}

// main 은 go 루틴을 기다리지 않음
func main() {
	c := make(chan int)
	go countToTen(c)
	for {
		a := <-c
		fmt.Printf("received %d\n", a)
	}
}
