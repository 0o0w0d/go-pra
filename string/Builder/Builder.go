package main

import (
	"fmt"
	"strings"
)

func ToUpper1(str string) string {
	var rst string
	for _, c := range str {
		if c >= 'a' && c <= 'z' {
			rst += string('A' + (c - 'a'))	// 한 연산 사용
		} else {
			rst += string(c)
		}
	}
	return rst
}

func ToUpper2(str string) string {
	var builder strings.Builder
	for _, c := range str {
		if c >= 'a' && c <= 'z' {
			builder.WriteRune('A' + (c - 'a'))	// strings.Builder 사용
		} else {
			builder.WriteRune(c)
		}
	}
	return builder.String()
}

func main()  {
	var str string = "Hello World"

	fmt.Println(ToUpper1(str))
	fmt.Println(ToUpper2(str))
}

// ToUpper1의 문제점:
// 문자열 연산(rst += string(c))을 사용할 때마다 새로운 문자열 생성
// 메모리 할당이 반복적으로 발생하여 성능 저하
// 문자열이 길어질수록 성능 차이가 더 커잠

// ToUpper2의 장점:
// strings.Builder는 내부적으로 바이트 슬라이스를 사용해 메모리를 효율적으로 관리
// 문자열 연산마다 새로운 메모리를 할당하지 않고, 필요할 때 버퍼를 늘림
// 최종적으로 한 번만 문자열로 변환하므로 성능이 더 좋음