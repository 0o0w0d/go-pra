package main

import "fmt"

// \n : 줄바꿈
// \t : 탭삽입
// \\ : \ 자체를 출력
// \" : " 자체를 출력

func main()  {
	str := "Hello\tGo\t\tWorld\n\"Go\" is Awesome!\n"
	
	fmt.Print(str)			// Hello   Go              World
							// "Go" is Awesome!
	fmt.Printf("%s", str)	// Hello   Go              World
							// "Go" is Awesome!
	fmt.Printf("%q", str)	// "Hello\tGo\t\tWorld\n\"Go\" is Awesome!\n"
}