package main

import "fmt"

func main()  {
	a := 1
	b := 1
	found := false		// boolean 변수를 flag로 설정

	for ; a <= 9; a++ {
		for b = 1; b <= 9; b++ {
			if a * b == 45 {
				found = true		// 찾았으면 break
				break
			}
		}
		if found {				// 바깥쪽 for문에서도 break
			break
		}
	}
	fmt.Printf("%d x %d = %d\n", a, b, a*b)
}