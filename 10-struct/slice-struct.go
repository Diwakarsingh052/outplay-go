package main

type student struct {
	name string
}

func main() {

	var s []student = []student{

		{name: "abc"},
		{name: "xyz"},
	}

	s1 := student{name: "abcde"}

	s = append(s,s1)


}
