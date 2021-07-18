package main

import (
	"fmt"
	"strings"
)

func main() {
	var (
		name   = "Sinione"
		number = 2
	)

	number = number << 3 // 2^3 * 2^3
	fmt.Println(number)
	name = strings.ToUpper(name)
	fmt.Println(name)
}
