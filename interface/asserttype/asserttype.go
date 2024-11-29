package main

import "fmt"

type Reader interface {
	Read()
}

type Closer interface {
	Close()
}

type File struct {
	
}

func (f *File) Read() {
}

func ReadFile(reader Reader)  {
	// Reader 인터페이스 변수를 Closer 인터페이스로 타입 변환 -> runtime error
	
	c := reader.(Closer)
	c.Close()
}

func ReadFile1(reader Reader)  {
	// 런타임 에러 발생하지 않는 타입 변환 방법

	// 타입 변환 가능 여부를 체크해서 Close() 메서드를 호출
	//  => 첫 번째 값: 타입 변환한 결과, 두 번째 값: 변환 성공 여부
	c, ok := reader.(Closer)
	if ok {
		c.Close()
	} else {
		fmt.Println("변환 실패")
	}

	// 한 줄로 사용 가능 
	if c, ok := reader.(Closer); ok {
		c.Close()
	} else {
		fmt.Println("변환 실패1")
	}
}

func main()  {
	file := &File{}
	// ReadFile(file)

	// panic: interface conversion: *main.File is not main.Closer: missing method Close

	// goroutine 1 [running]:
	// main.ReadFile(...)
	// 		/go-pra/interface/asserttype/asserttype.go:21
	// main.main()
	// 		/go-pra/interface/asserttype/asserttype.go:27 +0x27
	// exit status 2

	// File 타입은 Close() 메서드를 포함하고 있지 않아 Closer 인터페이스로 변환할 수 없음 => 런타임 에러 발생
	ReadFile1(file)
}