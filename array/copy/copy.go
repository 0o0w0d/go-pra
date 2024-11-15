package main

import "fmt"


func main()  {
	a := [5]int{1, 2, 3, 4, 5}
	b := [5]int{55, 44, 33, 22, 11}

	for i, v := range a {
		fmt.Printf("a[%d] = %d\n", i, v)
	}

	fmt.Println()

	for i, v := range b {
		fmt.Printf("b[%d] = %d\n", i, v)
	}

	b = a		// a 배열을 b 변수에 복사

	fmt.Println("---copy after---")
	for i, v := range b {
		fmt.Printf("b[%d] = %d\n", i, v)
	}

	// 얕은 복사(값만 복사)로 a 배열은 변화하지 않음(메모리 주소 값이 a 배열과 b 배열이 다름)
	b[1] = 100

	for i, v := range a {
		fmt.Printf("a[%d] = %d\n", i, v)
	}
}