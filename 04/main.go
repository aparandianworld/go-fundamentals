package main

import (
	"fmt"
	"strconv"
)

type Student struct {
	name  string
	score int
}

func main() {

	students := []Student{}
	var shouldContinue bool = true

	for shouldContinue {

		fmt.Println("Please select an option (1 - 3): ")
		fmt.Println("1) Enter a name and score: ")
		fmt.Println("2) Display all student scores: ")
		fmt.Println("3) Quit")

		var option string
		fmt.Scanln(&option)

		switch option {

		case "1":
			name, score := getUserInput()
			if name == "" || score == 0 {
				fmt.Println("Invalid input. Student name or score is invalid.")
				continue
			}
			students = append(students, Student{name: name, score: score})
		case "2":
			fmt.Println("Student scores:")
			printScores(students...)
		case "3":
			shouldContinue = false
		default:
			fmt.Println("Invalid option. Please enter a number between 1 and 3.")
		}
	}

}

func getUserInput() (string, int) {
	fmt.Println("Please enter a student name and score (separated by a space): ")
	var name, rawScore string
	fmt.Scanln(&name, &rawScore)
	score, err := strconv.Atoi(rawScore)
	if err != nil {
		fmt.Println("Invalid score. Please enter a valid integer.")
		return "", 0
	}

	return name, score
}

func printScores(students ...Student) { // variadic parameter
	for _, student := range students {
		fmt.Println(student.name, student.score)
	}
	fmt.Println()
}
