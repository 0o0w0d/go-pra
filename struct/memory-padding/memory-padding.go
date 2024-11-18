package main

import (
	"fmt"
	"unsafe"
)



type User1 struct {
	A	int8	// 1 byte
	B	int		// 8 byte
	C 	int8
	D 	int
	E 	int8
}

// 메모리 낭비를 줄이기 위해서 8byte보다 작은 필드는 8byte 크기(단위)를 고려해 몰아서 배치


type User2 struct {
	A	int8	// 1 byte
	C 	int8
	E 	int8
	B	int		// 8 byte
	D 	int
}

func main()  {
	user1 := User1{1, 2, 3, 4, 5}
	fmt.Println("user1 memory size:", unsafe.Sizeof(user1))		// expect: 19 byte -> result: 40 byte
	
	user2 := User2{1, 2, 3, 4, 5}
	fmt.Println("user2 memory size:", unsafe.Sizeof(user2))		// 24 byte
}

// 메모리 용량이 충분한 경우 패딩으로 인한 메모리 낭비를 신경쓰지 않아도 되지만, 
// 임베디드 하드웨어와 같이 메모리 공간이 매우 작은 경우 고려하는 것이 좋음