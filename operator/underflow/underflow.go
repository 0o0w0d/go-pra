package main

import "fmt"

// underflow: overflow와 반대로, 정수 타입이 표현할 수 있는 가장 작은 값에서 -1을 했을 때는 가장 큰 값으로 바뀜
// y가 정수 타입일 때 y > y - 1이 항상 true가 아닐 수도 있음


func main()  {
	
	var y int8 = -128
	fmt.Printf("\n%d > %d - 1: %v\n", y, y, y > y - 1)
	fmt.Printf("y\t= %4d, %08b\n", y, y)
	fmt.Printf("y - 1\t= %4d, %08b\n", y - 1, y - 1)
	fmt.Printf("y - 2\t= %4d, %08b\n", y - 2, y - 2)
	fmt.Printf("y - 3\t= %4d, %08b\n", y - 3, y - 3)

	// -128 > -128 - 1: false
	// y       = -128, -10000000
	// y - 1   =  127, 01111111
	// y - 2   =  126, 01111110
	// y - 3   =  125, 01111101
}