package main


type Stringer interface {
	String()
}

type Reader interface {
	Read()
}


func CheckAndRun(stringer Stringer)  {
	// r := stringer.(Reader)
	// r.Read()

	if r, ok := stringer.(Reader); ok {
		r.Read()
	}
}

func main()  {

}