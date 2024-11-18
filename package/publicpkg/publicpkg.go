package main

import (
	"fmt"
	"go-pra/package/publicpkg/pubpkg"
)


func main()  {
	fmt.Println("PI:", pubpkg.PI)
	pubpkg.PublicFunc()

	var myint pubpkg.Myint = 10
	fmt.Println("myint:", myint)

	var mystruct = pubpkg.MyStruct{Age: 18}
	fmt.Println("mystruct:", mystruct)
}