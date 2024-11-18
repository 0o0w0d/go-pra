package randomseed

import (
	"math/rand"
)

func RandomSeed() int {		// 함수가 값을 반환할 때는 반드시 함수 선언부에 반환 타입을 명시해야 함
	// rand.Seed(time.Now().UnixNano())
	// go 1.20 부터는 자동으로 안전한 random source 사용
	r := rand.Intn(100)		// 100-1까지의 자연수 중에 하나
	// fmt.Println(r)
	return r
}