package main

import "fmt"

/*
	Go는 Class나 Object가 없고 대신 Struct가 있음
*/

type person struct {
	name string
	age  int
}

// receiver 함수
// Struct을 입력안하면 모든 변수에 유효하지만 여기서는 person만을 지정함
func (p person) sayHello() {
	fmt.Printf("Hello! My Name is %s and I'm, %d\n", p.name, p.age)
}

func (p person) sayHi() {
	fmt.Printf("안녕 내 이름은 %s이고 내 나이는 %d\n", p.name, p.age)
}

func main() {
	// 동일
	// nico := person{name : "nico", age : 12}
	nico := person{"nico", 12}
	nico.sayHello()
	nico.sayHi()
}
