package main

import "fmt"

func main()  {
	str1 := "가나다라마"
	str2 := "abcde"

	fmt.Printf("len('가나다라마') = %d\n", len(str1))
	fmt.Printf("len('abcde') = %d\n", len(str2))

	// string 타입과 []rune 타입은 문자들의 집합을 나타냄 -> 상호 타입 변환 가능
	// 문자열의 각 문자가 rune 배열의 각 요소에 대입되어 글자 수를 알 수 있음
	runes := []rune(str1)
	fmt.Printf("len('가나다라마') = %d\n", len(runes))
}

// UTF-8에서 한글은 글자 당 3byte, 영문자는 글자 당 1byte

// 또한 string 타입 <-> []byte 타입도 상호 변환 가능