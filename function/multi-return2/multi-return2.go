package main

import "fmt"

// 반환할 변수에 이름을 지정할 경우 모든 반환 변수에 이름을 지정해야 함
func Divide(a, b int) (result int, success bool) {	// 반환할 함수명 명시
	if b == 0 {
		result = 0
		success = false

		return  // 명시적으로 반환할 값을 지정하지 않은 return 문
	}

	return a / b, true
}

func main()  {
	c, success := Divide(9, 3)
	fmt.Println(c, success)
	
	d, success := Divide(9, 0)
	fmt.Println(d, success)
}