package main

import "fmt"


func main()  {
	var a int = 200
	var p *int			// int를 가리키는 포인터 변수 p 선언

	p = &a				// a의 메모리 주소를 변수 p의 값으로 대입(복사)

	fmt.Println("변수 a의 메모리 주소 값:", &a)
	fmt.Println("포인터 변수 p의 값:", p)
	fmt.Println("포인터 변수 p의 메모리 주소 값:", &p)
	fmt.Println("포인터 변수 p가 가리키는 값:", *p)

	fmt.Println("---*p 값 변경---")
	*p = 100
	fmt.Println("변수 a의 값:", a)
}

// 변수 p (주소: &p)  ─────> 변수 a (주소: &a)
//     │                        │
//     │                        │
//   저장값: &a             저장값: 200
//     (p)                    (*p)