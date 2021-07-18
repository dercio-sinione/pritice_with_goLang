package main

import "fmt"

func main() {
	// Declaring a Slice
	langs := []string{"C"}
	// Add values into the Slice
	langs = append(langs,
		"Python",
		"Go",
		"C#",
		"JavaScript")

	for i := 0; i < len(langs); i++ {
		fmt.Println("Language", langs[i])
	}

	// Using Range in for Loop
	for i, lang := range langs {
		fmt.Println(i, lang)
	}
}
