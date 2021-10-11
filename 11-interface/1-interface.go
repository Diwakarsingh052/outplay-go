package main

import "fmt"

type reader interface {
	read(b []byte) (int, error)
	//read2(b []byte) (int, error) // we need to implement all methods of interface to satisfy it
}

type file struct {
	name string
}

func (f file) read(b []byte) (int, error) {
	s := "hello all go devs"
	copy(b, s)
	return len(s), nil
}

func (f file) read2(b []byte) (int, error) {
	return 0, nil
}

type jsonObject struct {
	name string
}

func (f jsonObject) read(b []byte) (int, error) {
	s := `{name:"dev"}`
	copy(b, s)
	return len(s), nil
}

func main() {

	f := file{name: "abc.txt"}
	o := jsonObject{name: "data.json"}
	fmt.Println(fetchData(f))
	fmt.Println(fetchData(o))
	var r reader
	fmt.Println("default", r)

}

//this func can accept any type that implements reader interface
func fetchData(r reader) error {
	fmt.Printf("type %T\n", r)
	data := make([]byte, 50)
	len, err := r.read(data)

	if err != nil {
		return err
	}

	fmt.Println(string(data[:len]))
	return nil

}
