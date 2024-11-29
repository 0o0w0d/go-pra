package main

import "fmt"

// 빈 인터페이스의 경우 모든 타입을 인수로 쓸 수 있음

// v.(type) 구문은 interface 타입의 변수에서만 사용 가능
// 컴파일 타임에 타입이 결정. interface 타입만이 실행 시점에 실제 타입을 확인할 필요성 있음

// v.(type) 사용 상황
// 1. interface{} 파라미터를 받는 함수에서
// 2. 빈 인터페이스(empty interface)를 다루는 코드에서
// 3. 여러 타입을 처리해야 하는 제네릭한 함수에서

func PrintVal(v interface{})  {	// 빈 인터페이스를 인수로 받음
	switch t:= v.(type)	{
	case int:
		fmt.Printf("v is int %d\n", int(t))
	case float64:
		fmt.Printf("v is float64 %f\n", float64(t))
	case string:
		fmt.Printf("v is string %s\n", string(t))
	default:
		fmt.Printf("Not supported type: %T:%v\n", t, t)
	}
}

type Student struct {
	Age int
}


func main()  {
	PrintVal(10)
	PrintVal(3.14)
	PrintVal("Hi")
	PrintVal(Student{15})
}