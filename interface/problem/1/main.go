package main

type ReadWriter interface {
	Read()
	Write()
}

type File struct {
	
}

func (f *File) Read() {
}

func ReadWrite(rw ReadWriter)  {
	rw.Read()
	rw.Write()
}

func main()  {
	f := &File{}
	ReadWrite(f)		
	// File 타입 구조체에는 Write() 메서드가 없어 ReadWriter 인터페이스로 사용 불가능
}