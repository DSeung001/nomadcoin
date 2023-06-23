package person

import "fmt"

/*
	대문자 = public
	소문자 = private
*/

type Person struct {
	name string
	age  int
}

// *를 사용함으로써 복사가 아닌 실제 값을 바꾸게
func (p *Person) SetDetails(name string, age int) {
	p.name = name
	p.age = age

	fmt.Println("SeeDetail's nico : ", p)
}

// 아래는 복사여도 상관이 없음
func (p Person) Name() string {
	return p.name
}
