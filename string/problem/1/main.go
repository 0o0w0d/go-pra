package main

import (
	"fmt"
	"unsafe"
)

func main()  {
	str1 := "학교종이 "
	str2 := "땡땡땡"
	
	fmt.Println("str1:", unsafe.StringData(str1))
	fmt.Println("str2:", unsafe.StringData(str2))
	str1 += str2

	fmt.Println(str1)		// 학교종이 땡땡땡
	fmt.Println("str1+str2:", unsafe.StringData(str1))
}