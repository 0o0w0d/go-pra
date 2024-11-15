package main

import "fmt"


func main()  {
	for a := 1; a <= 9; a += 2{
		fmt.Printf("%d * %d = %d\n", a, a, a*a)
	}
}