package main

import "fmt"


func main()  {
	str := "Hello 월드!"
	for _, v := range str {
		fmt.Printf(" 타입:%T 값:%d 문자값:%c\n", v, v, v)
	}
}

// 타입:int32 값:72 문자값:H
// 타입:int32 값:101 문자값:e
// 타입:int32 값:108 문자값:l
// 타입:int32 값:108 문자값:l
// 타입:int32 값:111 문자값:o
// 타입:int32 값:32 문자값: 
// 타입:int32 값:50900 문자값:월
// 타입:int32 값:46300 문자값:드
// 타입:int32 값:33 문자값:!

// 추가 메모리 할당 없ㅇ이 문자열을 한 글자씩 순회 -> 불필요한 메모리 낭비 X