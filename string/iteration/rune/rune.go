package main

import "fmt"

func main()  {
	str := "Hello 월드!"
	arr := []rune(str)

	for i := 0; i <len(arr); i++ {
		fmt.Printf(" 타입:%T 값:%d 문자값:%c\n", arr[i], arr[i], arr[i])
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

// []rune로 변환되는 과정에서 별도의 배열을 할당하므로 불필요한 메모리 사용