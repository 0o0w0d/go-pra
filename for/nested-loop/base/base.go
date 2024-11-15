package main

import "fmt"

// func main()  {
// 	for i := 0; i < 3; i++ {
// 		for j := 0; j < 5; j++ {
// 			fmt.Print("*")
// 		}
// 		fmt.Println()
// 	}
// }

// func main()  {
// 	for i := 0; i < 5; i++ {
// 		for j :=0; j < i+1; j++ {
// 			fmt.Print("*")

// 		}
// 		fmt.Println()
// 	}
// }


func main()  {
	dan := 2
	b := 1
	for {
		
		// 안쪽 for문
		for {
			fmt.Printf("%d x %d = %d\n", dan, b, dan*b)
			b++
			if b == 10 {
				break
			}
		}
		
		// 바깥쪽 for문
		b = 1
		dan++
		if dan == 10 {
			break
		}
	}
	fmt.Println("for문 종료")
}