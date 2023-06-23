package main

import "fmt"

/*
	Go는 배열이 무한할 수 없음 => C와 유사
	그래서 Slice 가 존재하며 무한히 커질 수 잇음 => 리스트와 유사
*/

func main() {
	foods := [3]string{"potato", "pizza", "pasta"}
	for _, food := range foods {
		fmt.Println(food)
	}

	for i := 0; i < len(foods); i++ {
		fmt.Println(foods[i])
	}

	// 크기를 정하지 않으면 slice
	newFoods := []string{"potato", "pizza", "pasta"}

	// append는 새로 생성, 직접적으로 바꾸지 않음
	newFoods2 := append(newFoods, "tomato")

	fmt.Printf("%v\n", newFoods)
	fmt.Printf("%v\n", newFoods2)
}
