package main

import "fmt"


func main()  {
	str1 := "Hello"
	str2 := "World"

	str3 := str1 + " " + str2

	fmt.Println(str3)

	str1 += " " + str2
	fmt.Println(str1)
}

// golang의 경우 문자열 연산으로 + 연산자만 가능