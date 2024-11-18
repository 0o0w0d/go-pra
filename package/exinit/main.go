package main

import (
	"fmt"
	"go-pra/exinit/einit"
)

func main()  {
	fmt.Println("main func")
	
	einit.PrintD()
}

// 패키지의 모든 전역 변수들이 초기화 된 후, init() 함수가 호출됨.