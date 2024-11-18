package main

import (
	"bufio"
	"fmt"
	"go-pra/NumQuest/randomseed"
	"os"
)

var stdin = bufio.NewReader(os.Stdin)

func InputValue() (int, error) {
	var n int
	_, err := fmt.Scanln(&n)
	if err != nil {
		stdin.ReadString('\n')
	}
	return n, err
}

func main() {
	r := randomseed.RandomSeed()
	var c int		// 입력 횟수 count

	for {
		fmt.Println("숫자를 입력하세요.")
		n, err := InputValue()
		
		if err != nil {
			fmt.Println("숫자만 입력")
		} else {
			c++
			
			if r > n {
				fmt.Printf("정답은 %d보다 큽니다.\n", n)
			} else if r < n {
				fmt.Printf("정답은 %d보다 작습니다.\n", n)
			} else {
				fmt.Printf("정답! 시도한 횟수: %d\n", c)
				break
			}
		}
	}
}