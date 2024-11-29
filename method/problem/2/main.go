package main

import "fmt"


type ParkingLot struct {
	LotSize		int
}

func ParkCar(lot *ParkingLot, carSize int)  {
	lot.LotSize -= carSize
}


func (lot *ParkingLot) ParkCarMethod(carSize int) {
	lot.LotSize -= carSize
}


func main()  {
	lot := &ParkingLot{100}
	ParkCar(lot, 10)
	fmt.Println(lot.LotSize)

	lot.ParkCarMethod(20)
	fmt.Println(lot.LotSize)
}