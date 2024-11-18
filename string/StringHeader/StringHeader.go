package main

import (
	"fmt"
	"unsafe"
)

type StringHeader struct {
    Data    uintptr     // 문자열의 데이터가 있는 메모리 주소를 나타내는 일종의 포인터
    Len     int         // 문자열의 길이
}

func main()  {
	str1 := "Hello World!"
	str2 := str1			// str1 변수값을 str2에 복사

	// Go에서 문자열은 내부적으로 StringHeader 구조체로 표현
	stringHeader1 := (*StringHeader)(unsafe.Pointer(&str1))
	stringHeader2 := (*StringHeader)(unsafe.Pointer(&str2))

	// 1. &str1: str1 문자열의 메모리 주소
	// 2. unsafe.Pointer(&str1): 문자열 주소를 unsafe.Pointer로 변환
	// 3. (*StringHeader)(...): unsafe.Pointer를 StringHeader 구조체 포인터로 타입 변환

	fmt.Println(stringHeader1)
	fmt.Println(stringHeader2)

	// &{4832072 12}
	// 첫 번째 숫자 4832072 : Data 필드 값 (메모리 주소)
	// 실제 "Hello World!" 문자열이 저장된 메모리 위치
	// str1과 str2가 같은 주소를 가짐 (데이터를 공유한다는 의미)

	// 두 번째 숫자 12:
	// Len 필드 값
	// "Hello World!"의 길이 (12글자)

	// str1과 str2는 같은 문자열을 가리키고 있음
	// 실제 문자열 데이터는 메모리에 한 번만 저장되어 있음
	// 두 변수가 그 데이터를 공유하고 있음

	// 문자열을 복사할 때 실제 데이터를 복사하지 않고, 같은 데이터를 참조
}