package main

import "time"

type Courier struct {
	Name	string
}

type Product struct {
	Name	string
	Price	int
	ID		int
}

type Parcel struct {
	Pdt				*Product
	ShippedTime		time.Time
	DeliveredTime	time.Time
}

func (c *Courier) SendProduct(pr *Product) *Parcel {
	pa := &Parcel{Pdt: pr, ShippedTime: time.Now()}
	return pa
}

func (pa *Parcel) Delivered() *Product {
	pa.DeliveredTime = time.Now()
	return pa.Pdt
}


// 포인터를 반환하도록 선언된 함수(*)에서는 반드시 주소값(&)을 반환해야 함