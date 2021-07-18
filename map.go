package main

import "fmt"

func main() {
	contryPopulation := map[string]int{
		"Angola":       123,
		"Portugal":     234,
		"South Africa": 242,
		"Canada":       443,
	}

	fmt.Println(contryPopulation)

	// Add new element to Map
	contryPopulation["Brazil"] = 532
	fmt.Println(contryPopulation["Brazil"])

	// delete Canada from map
	delete(contryPopulation, "Canada")
	fmt.Println(contryPopulation)

	// Verify if Canada is in the map
	temp, ok := contryPopulation["Canada"]

	if !ok {
		fmt.Println("Canada isn't in the map")
	} else {
		fmt.Println("Canada is in the map with value=", temp)
	}

	// See the length of map
	fmt.Println(len(contryPopulation))

}
