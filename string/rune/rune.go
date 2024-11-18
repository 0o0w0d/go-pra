package main

import "fmt"

// type rune int32 -> rune 타입과 int32는 같은 타입
// rune은 단일 Unicode 문자를 표현하는 타입

// rune은 작은따옴표 ' ' 를 사용
func main()  {
	var char rune = '한'

	fmt.Printf("%T\n", char)	// char 타입 출력	-> int32
	fmt.Println(char)			// -> 54620
	fmt.Printf("%c\n", char)	// 문자 출력		-> 한

}