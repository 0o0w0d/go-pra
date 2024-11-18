package main

import "fmt"

type User struct {
    Name    string
    ID      string
    Age     int
}

type VIPUser struct {
    UserInfo    User
    VIPLevel    int
    Price       int
}

func main()  {
	user := User{"zion", "zion123", 27}
	vip := VIPUser{
		User{"jojo", "joya", 10},
		3,
		250,
	}

	fmt.Printf("유저: %s ID: %s 나이: %d\n", user.Name, user.ID, user.Age)
	fmt.Printf("VIP 유저: %s ID: %s 나이: %d\n VIP 레벨: %d VIP 가격: %d 만원",
					vip.UserInfo.Name, 
					vip.UserInfo.ID, 
					vip.UserInfo.Age, 
					vip.VIPLevel,
					vip.Price)
}