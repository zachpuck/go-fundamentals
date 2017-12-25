package main

import (
	"fmt"
	"reflect"
)

func main() {

	name := "Zach"
	// course := "Go Fundamentals"
	module := 4.3
	ptr := &module

	fmt.Println("Name is", name, "and is of type", reflect.TypeOf(name))
	fmt.Println("Module is", module, "and is of type", reflect.TypeOf(module))
	fmt.Println("Memory address of *module* variable is", 
		ptr, "and the value of *module* is", *ptr)
}