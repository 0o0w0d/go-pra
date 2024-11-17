package main

import "fmt"

type Data struct {
	value 	int
	data	[200]int
}

// 포인터를 사용하지 않는 경우
func ChangeData1(arg Data)  {
	// 여기서 arg는 data의 복사본
	// 원본과 똑같은 새로운 데이터
	
	// 이 arg는 함수 안에서만 존재 => 함수가 끝나면 메모리에서 삭제
	arg.value = 999		// 복사본만 수정하는 거라 원본에 영향 없음
	arg.data[100] = 999

	// 이 함수 호출 시 data 변숫값이 모두 복사되어 구조체 크기만큼 복사됨 => 성능 문제 발생 가능성
}

// 포인터를 사용하는 경우
func ChangeData2(arg *Data)  {
	// 여기서 arg는 data의 위치
	// 원본의 메모리 값을 사용

	// 두 방식 모두 동작
	// (*arg).value = 888
	// (*arg).data[100] = 888

	arg.value = 888
	arg.data[100] = 888

	// 이 함수 호출 시 data의 메모리 주소만큼 복사
}

func main()  {
	var data1 Data

	ChangeData1(data1)
	fmt.Printf("value = %d\n", data1.value)
	fmt.Printf("data[100] = %d\n", data1.data[100])

	ChangeData2(&data1)
	fmt.Printf("value = %d\n", data1.value)
	fmt.Printf("data[100] = %d\n", data1.data[100])
}