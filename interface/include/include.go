package main

// Read, Close 메소드를 포함한 Reader 인터페이스
type Reader interface {
	Read() 	(n int, err error)
	Close()	error
}

// Write, Close 메소드를 포함한 Writer 인터페이스
type Writer interface {
	Write()	(n int, err error)
	Close()	error
}


// Reader, Writer 인터페이스의 메서드 집합을 모두 포함한 ReadWriter 인터페이스
// Read, Write, Close 메서드를 포함
type ReadWriter interface {
	Reader
	Writer
}

// 인터페이스 선언 시 메서드 이름이 중복되면 안되지만, 
// Close() error와 같이 시그니처(input, output이)가 동일한 메서드 형식일 경우 통합되어 
// 하나의 Close() error 메서드가 ReadWriter 인터페이스에 포함됨

// Reader 인터페이스 구현: Read()와 Close() 구현 필요
// Writer 인터페이스 구현: Write()와 Close() 구현 필요
// ReadWriter 인터페이스 구현: Read(), Write(), Close() 모두 구현 필요