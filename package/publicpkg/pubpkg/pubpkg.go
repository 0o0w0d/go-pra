package pubpkg

import "fmt"

const (
	PI = 3.1415
	pi = 3.141516
)

var ScreenSize int = 1080
var screenHeight int

func PublicFunc()  {
	const Myconst = 100				// 함수 내부에서 선언되어 패키지 외부로 공개되지 않음
	fmt.Println("This is public function:", Myconst)
}

func privateFunc()  {
	fmt.Println("This is private function")
}

type Myint int
type myString string

type MyStruct struct {
	Age int
	name string
}

func (m MyStruct) PublicMethod()  {
	fmt.Println("This is public method")
}

func (m MyStruct) privateMethod()  {
	fmt.Println("This is private method")
}

type myPrivateStruct struct {
	Age		int
	name	string
}

func (m myPrivateStruct) PrivateMethod()  {
	fmt.Println("This is private method")
}