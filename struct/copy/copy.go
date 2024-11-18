package main

import "fmt"


type Student struct {
	Age		int
	No		int
	Score	float64
}

func PrintStudent(s Student) {
	fmt.Printf("나이: %d 번호: %d 점수: %0.2f\n", s.Age, s.No, s.Score)
}


func main()  {
	student1 := Student{13, 20, 81.5}
	student2 := student1
	
	PrintStudent(student2)

	student1.Age = 17
	PrintStudent(student2)
}