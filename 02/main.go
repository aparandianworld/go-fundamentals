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

	// add key to map
	cart["grapes"] = 4
	fmt.Println("Updated cart: ", cart)
	printCart(cart)

	// remove key from map
	delete(cart, "orange")
	fmt.Println("Updated cart: ", cart)
	printCart(cart)

	clonedCart := maps.Clone(cart)
	fmt.Println("Cloned cart: ", clonedCart)
	printCart(clonedCart)

	// demonstrate maps are independent
	cart["banana"] = 10
	fmt.Println("Updated cart: ", cart)
	printCart(cart)
	fmt.Println("Cloned cart: ", clonedCart)
	printCart(clonedCart)

}

func printCart(cart map[string]int) {
	fmt.Printf("Cart contents (%d items): \n", len(cart))
	for item, quantity := range cart {
		fmt.Printf("item: %s, quantity: %d\n", item, quantity)
	}
}
