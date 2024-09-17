package main

import (
	"bufio" // io를 담당하는 패키지
	"fmt"
	"os" // 표준 입출력 등을 가지고 있는 패키지
)

func main()  {
	stdin := bufio.NewReader(os.Stdin)	// 표준 입력을 읽는 객체

	var a int
	var b int

	n, err := fmt.Scanln(&a, &b)
	if err != nil {
		fmt.Println(err)
		stdin.ReadString('\n')	// 에러 발생 시 출력 후 표준 입력 스트림 지우기
	} else {
		fmt.Println(n, a, b)
	}
	
	if err != nil {
		fmt.Println("다시 입력하세요: ")
		n, err = fmt.Scanln(&a, &b)		// 에러일 경우 다시 입력 받기
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(n, a, b)
		}
	}
}