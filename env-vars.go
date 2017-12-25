package main

import (
	"fmt"
	"os"
)

func main() {
	for _, env := range os.Environ() {
		fmt.Println(env)
	}

	name := os.Getenv("USERNAME")
	fmt.Println("\nThe current username is", name)
}