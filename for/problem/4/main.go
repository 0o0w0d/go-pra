package main

import (
	"fmt"
	// "strings"
)

// func main()  {
// 	for i := 5; i > 0; i-- {
// 		for j := 0; j < 5; j++ {
// 			fmt.Print(strings.Repeat("*", i))
// 			break
// 		}
// 		fmt.Println()
// 	}
// }


func main()  {
	for i := 0; i < 5; i++ {
		for j := 0; j < 5-i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}