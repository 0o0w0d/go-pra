package main

import "fmt"


type ColorType int   	// 별칭 ColorType을 선언하고 const 열거값 정의
const (
	Red ColorType = iota
	Blue
	Green
	Yellow
)


func colorToString(color ColorType) string {
	switch color {
	case Red:
		return "Red"
	case Blue:
		return "Blue"
	case Green:
		return "Green"
	case Yellow:
		return "Yellow"
	default:
		return "Undefined"
	}
}

func getMyColor() ColorType {
	return Blue
}

func main()  {
	fmt.Println("My favorite color is", colorToString(getMyColor()))	
}