package einit

import "fmt"

var (
	a = c + b		// a 값은 c와 b가 초기화된 후 초기화
	b = f()
	c = f()
	d = 3
)

func init()  {
	d++
	fmt.Println("init func", d)
}

func f() int {
	d++
	fmt.Println("f() d:", d)
	return d
}

func PrintD()  {
	fmt.Println("d:", d)
	fmt.Println("a:", a)
	fmt.Println("b:", b)
	fmt.Println("c:", c)
}