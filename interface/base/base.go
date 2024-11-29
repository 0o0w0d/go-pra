package main

type DuckInterface interface {
	Fly()
	Walk(distance int)	int
}

// 1. 메서드는 반드시 메서드명이 있어야 함
// 2. 매개변수와 반환이 다르더라도 이름이 같은 메서드는 있을 수 없음
// 3. 인터페이스에서는 메서드 구현을 포함하지 않음


// 아래 중 하나라도 다르면 컴파일러가 에러를 발생
// 메서드 이름
// 입력 매개변수의 타입과 개수
// 반환값의 타입과 개수
// 매개변수와 반환값의 순서

type Sample interface {
	String()	string
	String(int)	string		// error: duplicate method String 메서드명 중복
	_(x int)				// error: methods must have a unique non-blank name
}