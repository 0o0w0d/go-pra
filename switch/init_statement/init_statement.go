package main

import "fmt"

func getMyAge() int {
	return 28
}


func main()  {
	// switch age := getMyAge(); age {
	// case 10:
	// 	fmt.Println("Teenage")
	// case 33:
	// 	fmt.Println("Pair 3")
	// default:
	// 	fmt.Println("My age is", age)
	// }
	switch age := getMyAge(); true {
	case age < 10:
		fmt.Println("Child")
	case age < 20:
		fmt.Println("Teenager")
	case age < 30:
		fmt.Println("20s")
	default:
		fmt.Println("My age is", age)
	}

	// fmt.Println("age is", age)	// switch문의 초기문으로 age 변수를 선언했기 때문에, switch문 종료 후 age 변수에 접근 불가능
}