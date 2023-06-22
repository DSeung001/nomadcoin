package main

import "fmt"

// 전역변수에는 숏코드 사용을 못함
var globalName string = "nico"

func main() {
	// var name string = "nico"
	name := "nico"

	fmt.Println(name)
	fmt.Println(globalName)
}
