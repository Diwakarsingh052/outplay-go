package main

import "fmt"

func main() {

	s := "hello\nWorld" +
		""
	s1 := `hello\nWorld`
	path := "C:\\Users\\Diwakar\\go\\pkg\\mod"
	pathRaw := `C:\Users\Diwakar\go\pkg\mod

	 *************************
hjkl
	` + s

	fmt.Println(s)
	fmt.Println(s1)
	fmt.Println(path)
	fmt.Println(pathRaw)
	newString := fmt.Sprintf("hello %v hye", s)
	fmt.Println(newString)

}
