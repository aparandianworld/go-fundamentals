package main

import "fmt"

type MyStruct struct {
	name string
	id   int
}

func main() {

	s1 := MyStruct{
		name: "s1",
		id:   1,
	}

	s2 := s1 // struct are copied by value

	fmt.Println(s1)
	fmt.Println(s2)

	// modify s1
	s1.name = "s1 modified"

	fmt.Println(s1)
	fmt.Println(s2)

}
