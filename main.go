package main

import "fmt"

// Go는 여러개 값을 반환 가능
func plus(a int, b int, name string) (int, string) {
	return a + b, name
}

func plus2(a ...int) int {
	// total := 0 과 동일
	var total int

	// underscore는 생략
	for _, value := range a {
		total += value
	}
	return total
}

func main() {
	result, name := plus(2, 2, "nico")
	fmt.Println(result, name)
	result2 := plus2(1, 2, 3, 4, 5, 6, 7, 8, 9)
	fmt.Println(result2)

	name3 := "Nicolas ! ! ! ! ! ! Is yout name"
	for index, latter := range name3 {
		// 기본형인 byteㅊ가 나오기에 formatting 이 필요
		//fmt.Println(index, " : ", latter)
		fmt.Println(index, " : ", string(latter))
	}

}
