package main

import "fmt"

// go 언어에서는 break를 사용하지 않아도 case 하나를 실행 후 자동으로 switch문을 빠져나감
// 하나의 case 문 실행 후 다음 case 문까지 같이 실행하고 싶을 경우 fallthrough 키워드 사용

func main()  {
	a := 3

	switch a {
	case 1:
		fmt.Println("a == 1")
	case 2:
		fmt.Println("a == 2")
	case 3:
		fmt.Println("a == 3")
		fallthrough
	case 4:
		fmt.Println("a == 4")
	case 5:
		fmt.Println("a == 5")
	default:
		fmt.Println("a > 5")
	}
}