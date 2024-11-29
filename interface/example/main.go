package main

import (
	"go-pra/interface/example/fedex"
	"go-pra/interface/example/koreaPost"
)

// func SendBook(name string, sender *fedex.FedexSender)  {
// 	sender.Send(name)
// }

// func main()  {
// 	// fedex 전송 객체 생성
// 	sender := &fedex.FedexSender{}
// 	SendBook("어린 왕자", sender)
// 	SendBook("그리스인 조르바", sender)

// 	// 우체국 전송 객체 생성
// 	sender2 := &koreapost.PostSender{}
// 	SendBook("어린 왕자", sender2)		// cannot use sender2 (variable of type *koreapost.PostSender) as *fedex.FedexSender value in argument to SendBook
// }

// Sender 인터페이스 생성
type Sender interface {
	Send(parcel string)
}

func SendBook(name string, sender Sender)  {
	sender.Send(name)
}

// *koreaPost.PostSender, *fedex.FedexSender 모두 send(string) 메서드를 가지고 있어
// Sender 인터페이스로 사용 가능 => SendBook() 의 인수로 사용


func main()  {
	// fedex 전송 객체 생성
	sender := &fedex.FedexSender{}
	SendBook("어린 왕자", sender)
	SendBook("그리스인 조르바", sender)

	// 우체국 전송 객체 생성
	sender2 := &koreaPost.PostSender{}
	SendBook("어린 왕자", sender2)
	SendBook("그리스인 조르바", sender2)
}