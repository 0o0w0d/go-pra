package main

import (
	"fmt"
	"unsafe"
)

func main()  {
	var str string = "Hello world"
	var slice []byte = []byte(str)		// 슬라이스로 타입 변환

	slice[1] = 'a'
	
	fmt.Println(str)
	fmt.Printf("%s\n", slice)

	fmt.Printf("str:\t%p\n", unsafe.StringData(str))
	fmt.Printf("slice:\t%p\n", unsafe.SliceData(slice))

	// 문자열과 슬라이스의 실제 데이터가 저장된 메모리 주소를 직접 얻을 수 있는 함수 : unsafe.StringData / unsafe.SliceData
}