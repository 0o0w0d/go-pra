package main

import "fmt"

// 인터페이스 변수의 기본값은 유효하지 않은 메모리 주소를 나타내는 nil

type Attacker interface {
	Attack()
}

func main()  {
	var att Attacker
	
	fmt.Println(att)		// <nil>

	att.Attack()			// att 값이 유효하지 않은 주솟값인 nil이기 때문에 실행 중에 에러 발생

	// panic: runtime error: invalid memory address or nil pointer dereference
	// [signal SIGSEGV: segmentation violation code=0x1 addr=0x0 pc=0x480507]

	// goroutine 1 [running]:
	// main.main()
	//         /go-pra/interface/nil/nil.go:16 +0x47
	// exit status 2
}