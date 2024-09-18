package main

import "fmt"

// overflow: 정수가 정수 타입을 벗어난 경우 값이 비정상으로 변화하는 현상
// x가 정수 타입일 때 x < x + 1이 항상 true가 아닐 수도 있음

// 부호가 있는 정수에서 최상위 비트는 부호를 뜻하는 특수한 기능을 함

func main()  {
	var x int8 = 127

	fmt.Printf("%d < %d + 1: %v\n", x, x, x < x + 1)
	fmt.Printf("x\t= %4d, %08b\n", x, x)
	fmt.Printf("x + 1\t= %4d, %08b\n", x + 1, x + 1)
	fmt.Printf("x + 2\t= %4d, %08b\n", x + 2, x + 2)
	fmt.Printf("x + 3\t= %4d, %08b\n", x + 3, x + 3)

	// 127 < 127 + 1: false
	// x       =  127, 01111111
	// x + 1   = -128, -10000000
	// x + 2   = -127, -1111111
	// x + 3   = -126, -1111110
}