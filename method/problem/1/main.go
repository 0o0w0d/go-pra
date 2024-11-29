package main

import "fmt"

type Cart struct {
	productList	string
}

func (c *Cart) AddProduct(product string)  {
	if c.productList != "" {
		c.productList += ", "
	}
	c.productList += product
}

func (c *Cart) Clear() {
	c.productList = ""
}

func (c Cart) GetProductList() string {
	return c.productList
}

func main()  {
	c := &Cart{}
	c.AddProduct("Apple")
	c.AddProduct("Banana")

	fmt.Println(c.GetProductList())

	c.Clear()
	c.AddProduct("watermelon")
	fmt.Println(c.GetProductList())
}

// 구조체의 상태를 변경하는 메서드 -> 포인터 리시버 사용 (AddProduct, Clear)
// 단순히 값을 읽는 메서드 -> 값 리시버 사용해도 됨 (GetProductList)