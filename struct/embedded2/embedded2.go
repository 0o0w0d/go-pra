package main

import "fmt"

type User struct {
    Name    string
    ID      string
	Age		int
	Level	int
}

type VIPUser struct {
    User				// 필드명 생략
    Level    int
    Price    int
}


func main()  {
	user := User{"zion", "zion123", 27, 10}
	vip := VIPUser{
		User{"jojo", "joya", 10, 7},
		3,
		250,
	}

	fmt.Printf("유저: %s ID: %s 나이: %d 레벨: %d\n", user.Name, user.ID, user.Age, user.Level)
	fmt.Printf("VIP 유저: %s ID: %s 나이: %d 레벨: %d\n VIP 레벨: %d VIP 가격: %d 만원",
					vip.Name, 
					vip.ID, 
					vip.Age,
					vip.User.Level,	// 포함된 구조체명을 쓰고 접근
					vip.Level,		// VIPUser의 Level
					vip.Price)
}