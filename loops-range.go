package main

import (
	"fmt"
)

func main() {
	// slice of cats
	catsInHouse := []string{"Eva", "Milo", "Mia"}
	catsOutside := []string{"Fred", "Mia", "Snowball"}

	// settings index to '_', and i to the value
	for _, i := range catsInHouse {
		fmt.Println(i)
		for _, j := range catsOutside {
			if i == j {
				fmt.Println("The cats escaped! Quick! Get", i)
			}
		}
	}
}