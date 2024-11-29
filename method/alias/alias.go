package main

import "fmt"

type myInt int

func (a myInt) add(b int) int {
	return int(a) + b
}

func main()  {
	var a myInt = 10
	fmt.Println(a.add(30))

	var b int = 20
	fmt.Println(myInt(b).add(50))		// int 타입을 myInt로 타입 변환해 메소드 사용
}