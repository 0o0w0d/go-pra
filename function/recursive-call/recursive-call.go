package main

import "fmt"

func printNo(n int)  {
	// base case: 재귀 함수를 종료하는 부분
	if n == 0 {
		return
	}

	// recursive case: 자기 자신을 호출하는 부분
	fmt.Println(n)
	printNo(n-1)
	fmt.Println("After", n)
}

func main()  {
	printNo(3)
}