package main

import "fmt"

type account struct {
	balance		int
	firstName	string
	lastName	string
}

func (a1 *account) withdrawPointer(amount int) {
	a1.balance -= amount
}
// 포인터를 통해 직접 원본 객체를 수정

func (a2 account) withdrawValue(amount int) {
	a2.balance -= amount
}
// account의 복사본이 생성되어 그 복사본에서만 작업 수행
// 함수 종료 시 변경사항이 사라짐

func (a3 account) withdrawReturnValue(amount int) account {
	a3.balance -= amount
	return a3
}
// 복사본을 수정하고 그 수정된 복사본을 반환

// 값을 변경 시 해당 값을 리턴하지 않으면 원본에는 영향가지 않고 스택 프레임이 해제
// 포인터를 통해 원본 값 수정

func main()  {
	var mainA *account = &account{100, "jjo", "bae"}
	mainA.withdrawPointer(30)
	fmt.Println("A:", mainA.balance)

	mainA.withdrawValue(20)
	fmt.Println("A:", mainA.balance)

	var mainB account = mainA.withdrawReturnValue(20)
	fmt.Println("A:", mainA.balance)
	fmt.Println("B:", mainB.balance)

	mainB.withdrawPointer(30)
	fmt.Println("B:", mainB.balance)
}