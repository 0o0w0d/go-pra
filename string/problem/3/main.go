package main

import (
	"fmt"
	"unsafe"
)

type User struct {
	FirstName	string
	LastName	string
	Age			int
}

// 문자열 구조체 => 16 byte
// type StringHeader struct {
//     Data    uintptr     // 8 byte
//     Len     int         // 8 byte
// }

func main()  {
	var user User
	fmt.Println(unsafe.Sizeof(user))
}