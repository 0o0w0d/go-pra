package main

import "fmt"

// 실수끼리의 == 연산에서 예기치 못한 결과가 나올 수도 있음
// float 표현은 표현 방식으로 인한 오차 발생

// 컴퓨터의 지수부와 소수부는 10진수 기준이 아니라 2진수 기준
// 소수점 이하 숫자들을 2의 음수 승수로 표현 -> 최대한 가까운 값으로 사용
// ex) float32 0.376 -> 0.375999987125396728515625

// 이 문제의 해결을 위해 operator/float-resolve 참고

func main()  {
	var a float64 = 0.1
	var b float64 = 0.2
	var c float64 = 0.3		// 0.3의 정확한 실수 값을 2진수 체계로 표현할 수 없음

	fmt.Printf("%f + %f == %f : %v\n", a, b, c, a + b == c)
	fmt.Println(a + b)

	// 0.100000 + 0.200000 == 0.300000 : false
	// 0.30000000000000004
}