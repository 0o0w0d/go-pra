package main

import "fmt"


func Divide(a, b int) (int, bool) {
	if b == 0 {
		return 0, false
	}

	return a / b, true
}

func main()  {
	// go에서 선언된 변수는 반드시 사용되어야 하며, 함수의 반환 값을 모두 받거나
	// 명시적으로 무시(_) 해야 함

	// 이렇게 설계한 이유
	// 1. 코드의 명확성을 높이기 위해
	// 2. 실수로 중요한 반환값(특히 error)을 놓치는 것을 방지하기 위해
	// 3. 미사용 변수로 인한 메모리 낭비를 방지하기 위해

	// c := Divide(9, 3) -> Error: assignment mismatch: 1 variable but Divide returns 2 values

	c, success := Divide(9, 3)
	fmt.Println(c, success)
	
	d, success := Divide(9, 0)
	fmt.Println(d, success)
}