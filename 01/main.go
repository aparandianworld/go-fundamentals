package main

import "fmt"

func main() {

	originalArray := [3]int{1, 2, 3}
	copiedArray := originalArray

	fmt.Printf("Original array: %v Original array memory address: %p\n", originalArray, &originalArray)
	fmt.Printf("Copied array: %v Copied array memory address: %p\n", copiedArray, &copiedArray)

	originalArray[0] = 10
	fmt.Printf("Original array: %v Original array memory address: %p\n", originalArray, &originalArray)
	fmt.Printf("Copied array: %v Copied array memory address: %p\n", copiedArray, &copiedArray)

	fmt.Println("---")

	originalSlice := []int{1, 2, 3}
	copiedSlice := originalSlice

	fmt.Printf("Original slice: %v Original slice memory address: %p\n", originalSlice, &originalSlice)
	fmt.Printf("Copied slice: %v Copied slice memory address: %p\n", copiedSlice, &copiedSlice)

	originalSlice[0] = 10
	fmt.Printf("Original slice: %v Original slice memory address: %p\n", originalSlice, &originalSlice)
	fmt.Printf("Copied slice: %v Copied slice memory address: %p\n", copiedSlice, &copiedSlice)

}
