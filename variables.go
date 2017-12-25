package main

import (
	"fmt"
	"reflect"
)

func main() {

	name := "Zach"
	course := "Go Fundamentals"
	module := 4.3

	fmt.Println("Name is", name, "and is of type", reflect.TypeOf(name))
	fmt.Println("Course is", course, "and is of type", reflect.TypeOf(course))
	fmt.Println("Module is", module, "and is of type", reflect.TypeOf(module))
}