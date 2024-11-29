package main

import "fmt"

type Stringer interface {
	String()	string
}

type Student struct {
	Age int
}

func (s *Student) String() string {
	return fmt.Sprintf("Student Age: %d\n", s.Age)
}

// Student 타입은 String() 메서드를 가지고 있어, PrintAge() 함수의 stringer 인터페이스 인수로 사용 가능

func PrintAge(stringer Stringer)  {
	s := stringer.(*Student)		// *Student 타입으로 타입 변환
	fmt.Printf("Age: %d\n", s.Age)
}

type Student1 struct {
}

type Actor struct {
}

func (a *Actor) String() string {
	return "Actor"
}

func ConvertType(stringer Stringer)  {
	student := stringer.(*Student)
	fmt.Println(student)
}

func main() {
	s := &Student{15}	// *Student 타입 변수 선언 및 초기화
	
	PrintAge(s)			// 변수 s를 인터페이스 인수로 함수 호출

	
	// Student1 타입은 String() 메서드를 포함하지 않아 
	// Stringer 인터페이스에서 Student1로 타입 변환 불가능
	// s1 := &Student1{}
	// PrintAge(s1)

	// var s1 Stringer
	// s1.(*Student1)		// compile type error
	// concretetype/concretetype.go:37:2: impossible type assertion: s1.(*Student1)
    //     *Student1 does not implement Stringer (missing method String)


	actor := &Actor{}
	ConvertType(actor)

	// panic: interface conversion: main.Stringer is *main.Actor, not *main.Student

	// goroutine 1 [running]:
	// main.ConvertType({0x4b87a8?, 0x55bf50?})
	// 		/go-pra/interface/concretetype/concretetype.go:35 +0x85
	// main.main()
	// 		/go-pra/interface/concretetype/concretetype.go:56 +0x4a
	// exit status 2

	// *Actor 타입이 Stringer 인터페이스는 구현하고 있지만 
	// *Student 타입으로는 안전하게 변환될 수 없어서 발생하는 런타임 에러 발생
}