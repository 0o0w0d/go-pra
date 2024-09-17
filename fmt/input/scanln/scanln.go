package main

import "fmt"

// 마지막 입력 이후 enter 키로 입력을 종료

// Scan vs Scanln
// Scan : 공백(스페이스, 탭, 줄바꿈 등)으로 구분된 입력을 읽음 => 여러 개의 값을 공백으로 구분하여 읽을 수 있으며, 입력 중 공백 문자로 분리된 값들을 계속 읽음
// Scanln : 한 줄의 입력을 읽고, 공백으로 구분된 값들을 처리 => 줄바꿈 문자가 입력의 끝을 표시하는 역할, 한 줄 전체를 읽고 입력이 끝나면 해당 줄의 값을 처리

func main()  {
	var a int
	var b int

	n, err := fmt.Scanln(&a, &b)
	if err != nil {
		fmt.Println(n, err)
	} else {
		fmt.Println(n, a, b)
	}
}