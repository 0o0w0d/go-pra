package main

import "fmt"

// Scanf : 서식에 맞춰 입력을 받음


func main()  {
	var a int
	var b int

	n, err := fmt.Scanf("%d %d\n", &a, &b)		// n m 형태로 입력을 받아야 함 (무조건 한 칸 공백)
	if err != nil {
		fmt.Println(n, err)
	} else {
		fmt.Println(n, a, b)
	}
}