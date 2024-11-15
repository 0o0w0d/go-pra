package main

import "fmt"

func ChangeArray(arr [5]int)  {
	arr[3] = 3000
}

func main()  {
	a := [5]int{1, 2, 3, 4, 5}

	ChangeArray(a)

	fmt.Println(a[3])			// 4
	// ChangeArray() 함수의 인수로 a 값이 복사되기 때문에, arr와 a는 서로 다른 메모리 주소를 참조
}