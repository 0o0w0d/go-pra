package main

import "fmt"

// 오차를 무시하는 방법 1
// 아주 작은 오차를 무시하는 방법으로 값을 비교

// 임의로 정해준 오차가 변수에 따라 작은 값이 아닐 수도 있는 문제 발생
// => 값의 차이가 1비트만큼 난다면 같은 값이라고 간주
// => math 패키지의 Nextafter() 함수 사용

const epsilon = 0.000001	// 매우 작은 값

func equal(a, b float64) bool {
	if a > b {
		if a - b <= epsilon {	// 작은 차이 무시
			return true
		} else {
			return false
		}
	} else {
		if b - a <= epsilon {
			return true
		} else {
			return false
		}
	}
}


func main()  {
	var a float64 = 0.1
	var b float64 = 0.2
	var c float64 = 0.3

	fmt.Printf("%0.18f + %0.18f = %0.18f\n", a, b, a + b)
	fmt.Printf("%0.18f == % 0.18f : %v\n", c, a + b, equal(a+b, c))

	// 0.100000000000000006 + 0.200000000000000011 = 0.300000000000000044
	// 0.299999999999999989 ==  0.300000000000000044 : true


	a = 0.0000000000004
	b = 0.0000000000002
	c = 0.0000000000007

	fmt.Printf("%g == %g : %v\n", c, a+b, equal(a+b, c))

	// 7e-13 == 6.000000000000001e-13 : true
}