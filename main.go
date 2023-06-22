package main

import "fmt"

func main() {
	x := 405940594059

	fmt.Printf("%b\n", x)
	fmt.Printf("%o\n", x)
	fmt.Printf("%x\n", x)
	fmt.Printf("%U\n", x)

	// Format 된 데이터를 반환 함
	xAsBinary := fmt.Sprintf("%b\n", x)
	fmt.Println(xAsBinary)

	// 아래 공식 문서에서 확인이 가능
	// https://pkg.go.dev/fmt
}
