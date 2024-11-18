package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
)

const (
	Money = 1000
	Earn = 500
	Lose = 100
	Victory = 5000
	GameOver = 0
)

func InputValue() (int, error) {
	var n int
	
	stdin := bufio.NewReader(os.Stdin)		// 표준 입력 받기
	_, err := fmt.Scanln(&n)
	if err != nil {
		stdin.ReadString('\n')
	}
	return n, err
}


func main()  {
	m := Money

	for {
		r := rand.Intn(5) + 1
		n, err := InputValue()
		
		if err != nil {
			fmt.Println("숫자만 입력하세요")
		} else {
			fmt.Println("입력한 숫자:", n)
			
			if r == n {
				m += Earn
				fmt.Printf("정답! 현재 잔액: %d원\n", m)
			} else {
				m -= Lose
				fmt.Printf("ㅠㅠ 현재 잔액: %d원\n", m)
			}

			if m >= Victory {
				fmt.Println("승리!")
				break
			}

			if m <= GameOver {
				fmt.Println("게임 오버")
				break
			}
		}
	}
}