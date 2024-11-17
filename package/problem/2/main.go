package main

import "fmt"

var (
	a = b + f()	// 9
	b = c		// 4
	c = f() 	// 4
	d = 3		// 5
)

func init()  {
	fmt.Printf("init func a:%d b:%d c:%d d:%d\n", a, b, c, d)
}

func f() int {
	d++
	fmt.Println("f() d:", d)
	return d
}
