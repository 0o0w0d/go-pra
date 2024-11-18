package main

import "fmt"

type Person struct {
	Name 	string
	Age		int
	Address string
	Tall	float64
}

func main()  {
	var person Person
	person.Age = 27
	person.Tall = 160.123
	person.Name = "zion"
	person.Address = "경기도"

	fmt.Println("주소:", person.Address)
	fmt.Printf("키: %0.2fcm\n", person.Tall)
	fmt.Printf("나이: %d세\n", person.Age)
	fmt.Println("이름:", person.Name)
}