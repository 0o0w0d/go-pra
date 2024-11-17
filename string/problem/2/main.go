package main

import (
	"fmt"
	"strings"
)

func ToUpper(str string) string {
	var builder strings.Builder
	for _, v := range str {
		if v >= 'a' && v <= 'z' {					// 소문자인지 확인
			builder.WriteRune('A' + (v - 'a'))		// 대문자로 변환
		} else {
			builder.WriteRune(v)					// 소문자가 아니면 그대로 추가
		}
	}
	return builder.String()
}

func main()  {
	str := "hello WOrld!"
	fmt.Println(ToUpper(str))
}