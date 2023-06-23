package main

import "fmt"

func main() {

	// 2 출력
	a := 2
	b := a
	a = 12
	fmt.Println(b)
	fmt.Println(&a, &b)

	// 둘이 동일
	c := &b
	fmt.Println(&b, &c)

	// 아래와 같이 기존 변수에는 못하고 새로 정의시 사용 가능
	//a = &b

	*c = 123
	fmt.Println(c, *c, &b, b)
}
