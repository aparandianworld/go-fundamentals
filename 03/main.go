package main

import "fmt"

type MyStruct struct {
	name string
	id   int
}

type Score struct {
	name  string
	count int
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

	students := []string{
		"Aaron Parandian",
		"John Doe",
	}

	scores := []Score{
		{name: "Aaron Parandian", count: 5},
		{name: "John Doe", count: 10},
	}

	fmt.Println("Student scores: ")
	fmt.Println(students[0], scores[0].count)
	fmt.Println(students[1], scores[1].count)

}
