package main

import "fmt"

func main() {
	name := SayName()
	fmt.Println("Hello", name)
}

func SayName() string {
	return "Derone"
}
