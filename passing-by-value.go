package main

import (
	"fmt"
)

func main() {
	name := "Zach"
	course := "Go Fundamentals"
	
	fmt.Println("\nHi", name, "you're currently watching",
	course)

	changeCourse(&course)

	fmt.Println("\nYou are now watching course", course)
}

func changeCourse(course *string) string {
	*course = "Node.js Testing Strategies"

	fmt.Println("Trying to chnage your course to", *course)
	
	return *course
}