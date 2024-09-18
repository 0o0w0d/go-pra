package main

import (
	"fmt"
	"math"
)

// 오차를 무시하는 방법 2

// func Nextafter(x, y float64) (r float64)
// x에서 y를 향해서 1비트만 조정한 값 반환
// x < y 라면 x에서 1비트만큼 증가시킨 값
// x > y 라면 x에서 1비트만큼 감소시킨 값

// 오차가 작을 뿐, 정확한 계산은 아님


func equal(a, b float64) bool {
	return math.Nextafter(a, b) == b
}


func main()  {
	var a float64 = 0.1
	var b float64 = 0.2
	var c float64 = 0.3

	fmt.Printf("%0.18f + %0.18f = %0.18f\n", a, b, a + b)
	fmt.Printf("%0.18f == %0.18f : %v\n", c, a + b, equal(a+b, c))

	// 0.100000000000000006 + 0.200000000000000011 = 0.300000000000000044
	// 0.299999999999999989 == 0.300000000000000044 : true


	// 매우 작은 값으로 변경

	a = 0.0000000000004
	b = 0.0000000000002
	c = 0.0000000000007

	fmt.Printf("%g == %g : %v\n", c, a+b, equal(a+b, c))

	// 7e-13 == 6.000000000000001e-13 : false
}