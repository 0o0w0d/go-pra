package main

import (
	"bufio"
	"fmt"
	"os"
)

func main()  {
	stdin := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("값을 입력하세요.")

		var num int
		
		_, err := fmt.Scanln(&num)	// 한 줄 입력 받기
		if err != nil {
			fmt.Println(("숫자를 입력하세요."))

			stdin.ReadString('\n')	// 키보드 버퍼 지우기
			continue				// return # 11
		}
		fmt.Printf("입력하신 숫자는 %d입니다.\n", num)
		if num % 2 == 0 {
			break					// 짝수라면 for문 종료
		}
	}
	fmt.Println("for문 종료.")
}