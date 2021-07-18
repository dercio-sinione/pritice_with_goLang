package main

import (
	"fmt"
	"training/utils"
)

func main() {
	p := utils.Person{}

	p.Bi = "006153280LA047"
	p.Firstname = "Dércio"
	p.Lastname = "Sinione Derone"
	p.Fullname = (p.Firstname + p.Lastname)
	p.Gender = "Male"
	p.Age = 21

	fmt.Printf("%+v", p)
}
