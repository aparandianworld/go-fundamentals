package main

import (
	"fmt"
	"maps"
)

func main() {

	cart := map[string]int{
		"apple":  1,
		"banana": 2,
		"orange": 3,
	}

	fmt.Println("Original cart: ", cart)
	printCart(cart)

	fmt.Println("----------------")

	// key exists in map
	if quantity, ok := cart["apple"]; ok {
		fmt.Println("Apple quantity: ", quantity)
	}

	// key does not exist in map
	if quantity, ok := cart["grapes"]; ok {
		fmt.Println("Grapes quantity: ", quantity)
	} else {
		fmt.Println("Grapes not found in cart")
	}

	fmt.Println("----------------")

	// add key to map
	cart["grapes"] = 4
	fmt.Println("Updated cart: ", cart)
	printCart(cart)

	fmt.Println("----------------")

	// remove key from map
	delete(cart, "orange")
	fmt.Println("Updated cart: ", cart)
	printCart(cart)

	fmt.Println("----------------")

	clonedCart := maps.Clone(cart)
	fmt.Println("Cloned cart: ", clonedCart)
	printCart(clonedCart)

	fmt.Println("----------------")

	copiedCart := cart

	// demonstrate maps are independent
	cart["banana"] = 10
	fmt.Println("Updated cart: ", cart)
	fmt.Println("Copied cart: ", copiedCart)
	printCart(cart)
	printCart(copiedCart)

	fmt.Println("----------------")

	fmt.Println("Cloned cart: ", clonedCart)
	printCart(clonedCart)

	fmt.Println("----------------")
}

func printCart(cart map[string]int) {
	fmt.Printf("Cart contents (%d items): \n", len(cart))
	for item, quantity := range cart {
		fmt.Printf("item: %s, quantity: %d\n", item, quantity)
	}
}
