package main

import "fmt"

// 선언하고 사용하지 않는 변수가 있을 경우 컴파일에러 발생
// range 사용 시, index를 사용하지 않는다면 _ 로 무효화 필요
func main()  {
	var t [5]float64 = [5]float64{24.0, 25.1, 23.7, 29.3, 21.9}

	// for i, v := range t {
	// 	fmt.Println(i, v)
	// }

	for _, v := range t {
		fmt.Println(v)
	}
}