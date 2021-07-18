package main

import (
	"fmt"
)

func main() {
	cards := []string{"C#", "JavaScript"}
	cards = append(cards, "GoLang", "C", "Python")

	fmt.Println("Old Slice")
	PrintList(cards)
	// Removing the third element of Slice
	// cards = append(cards[0:3], cards[4:]...)
	cards = Remove(cards, 3)
	// cards = MyRemove(cards, 3)
	fmt.Println("\n New Slice")
	PrintList(cards)
}

// Print a List of Slice
func PrintList(list []string) {
	for i, item := range list {
		fmt.Println(i, item)
	}
}

// Remove elements from a slice at the specific position
func Remove(slice []string, index int) []string {
	slice = append(slice[0:index], slice[index+1:]...)
	return slice
}

// func append(slice []Type, elems ...Type) []Type
// func MyRemove(slice []Type, index int) []Type {
// 	slice = append(slice[0:index], slice[index+1:]...)
// 	return slice
// }
