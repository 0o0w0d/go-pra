package main

import "fmt"

func main()  {
	day := "thursday"

	switch day {
	case "monday", "tuesday":
		fmt.Println("월, 화요일은 수업 가는 날!")
	case "wednesday", "thursday", "friday":
		fmt.Println("수, 목, 금은 실습하는 날!") 
	}
}