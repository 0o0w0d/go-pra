package main

import "fmt"

// 내부의 for문을 함수로 따로 빼내 구현
func check45(a int) (int, bool) {
	for b := 1; b <= 9; b++ {
		if a * b == 45 {
			return b, true
		}
	}
	return 0, false
}

func main()  {
	a := 1
	b := 0

	for ; a <= 9; a++{
		var found bool
		if b, found = check45(a); found {
			break
		}
	}
	fmt.Printf("%d x %d = %d\n", a, b, a * b)
}