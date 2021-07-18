package main

import "fmt"

func main() {
	var (
		n1 int = 1
		n2 int = 2
	)

	fmt.Println(somar(n1, n2))
}

func somar(n1 int, n2 int) int {
	return n1 + n2
}

func subtrair(n1, n2 int) int {
	return n1 - n2
}
