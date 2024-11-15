package main

import (
	"fmt"
)

// func main()  {
// 	i := 1
// 	for {
// 		time.Sleep(time.Second)
// 		fmt.Println(i)
// 		i++
// 	}
// }

func main()  {
	for i, v := range []int{1,2,3,4,5,6,7,8,9,10} {
		fmt.Printf("%d: %d\n", i, v)
	}
}