package main

import "fmt"

type Stringer interface {
	String()	string
}

type Student struct {
	Name	string
	Age		int
}

func (s Student) String() string {
	return fmt.Sprintf("age: %d, Name: %s", s.Age, s.Name)
}

func main()  {
	student := Student{"쪼쪼", 10}	// Student 타입
	var stringer Stringer			// Stringer 타입
	
	stringer = student				// stringer 값으로 student 대입

	fmt.Printf("%s\n", stringer.String())	// stringer의 String() 메서드 호출
}

// fmt.Sprintf -> 문자열 반환
// fmt.Printf -> 화면 출력