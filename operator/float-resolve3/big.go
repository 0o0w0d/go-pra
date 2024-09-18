package main

import (
	"fmt"
	"math/big"
)

// math/big 패키지의 Float을 이용해 정밀도를 직접 조정
// x.Cmp(y) : x, y - Float 객체
// return -1: x < y / 1: x > y / 0: x == y

func main()  {
	a, _ := new(big.Float).SetString("0.1")
	b, _ := new(big.Float).SetString("0.2")
	c, _ := new(big.Float).SetString("0.3")

	d := new(big.Float).Add(a, b)
	fmt.Println(a, b, c, d)
	fmt.Println(c.Cmp(d))		

	// 0.1 0.2 0.3 0.3
	// 0
}