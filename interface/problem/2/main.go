package main


type OurDB struct {
	Name	string
}

func (db *OurDB) GetData() string {
	return ""
}

func (db *OurDB) WriteData(data string) {
	return 
}

func (db *OurDB) Close() error {
	return nil
}

type DB interface {
	GetData()				string
	WriteData(data string)
	Close()					error
}