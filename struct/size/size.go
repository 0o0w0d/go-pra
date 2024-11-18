package main

import (
	"fmt"
	"unsafe"
)

type User struct {
	Age		int32		// 4 byte
	Score	float64		// 8 byte
}


func main()  {
	user := User{18, 81.5}
	fmt.Println(unsafe.Sizeof(user))		// expect: 12 -> 16
}
