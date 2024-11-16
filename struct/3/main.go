package main

import (
	"fmt"
	"unsafe"
)

type Padding1 struct {
	A	int8	// 1 byte
	B	int		// 8 byte
	C 	float64	// 8
	D 	uint16	// 2
	E 	int		// 8
	F 	float32	// 4
	G 	int8	// 1
}

type Padding2 struct {
	A	int8	// 1 byte
	G 	int8	// 1
	D 	uint16	// 2
	F 	float32	// 4
	B	int		// 8 byte
	C 	float64	// 8
	E 	int		// 8
}

func main()  {
	var padding1 Padding1
	var padding2 Padding2

	fmt.Println(unsafe.Sizeof(padding1))
	fmt.Println(unsafe.Sizeof(padding2))
	fmt.Printf("Pointer size: %d bytes\n", unsafe.Sizeof(uintptr(0)))
}

// uint16는 2바이트 단위로 정렬
// int32는 4바이트 단위로 정렬
// int64는 8바이트 단위로 정렬