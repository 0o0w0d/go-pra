package main

import "fmt"

// fmt.Scan() 함수의 입력으로는 변수의 메모리 주소를 입력으로 넘겨야 함(포인터:& 이용)

func main()  {
	var a int
	var b int

	n, err := fmt.Scan(&a, &b)	// return : 읽은 항목의 개수, error
	if err != nil {
		fmt.Println(n, err)
	} else {
		fmt.Println(n, a, b)
	}
}

// 첫 번째 입력이 잘못되면 두 번째 입력을 받지 않고 함수가 에러 반환