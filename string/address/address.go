package main

import (
	"fmt"
	"unsafe"
)

func main()  {
	var str string = "Hello"
	addr1 := unsafe.StringData(str)

	str += " World"
	addr2 := unsafe.StringData(str)

	str += " Welcome!"
	addr3 := unsafe.StringData(str)

	fmt.Println(str)
	fmt.Printf("addr1:\t%p\n", addr1)
	fmt.Printf("addr2:\t%p\n", addr2)
	fmt.Printf("addr3:\t%p\n", addr3)
}

// Hello World Welcome!
// addr1:  0x49cd7a
// addr2:  0xc0000120e0
// addr3:  0xc0000180d8

// Golang 은 기존 문자열 메모리 공간을 건드리지 않고, 새로운 메모리 공간을 만들어서 두 문자열을 합산
// => string 합 연산을 빈번하게 할 경우 메모리 낭비 => strings 패키지의 Builder 이용해 낭비를 줄일 수 있음!