// arrays: numbered lists of a single type
// all slices are built on top of arrays
// slices are references
package main

import (
    "fmt"
)

func main() {
    //Declares a slice called myCats
    //myCats := make(<type>, <len>, <cap>)
    // square brackes indicate we are create a slice. 
    
    // defining the the type, len, and cap
    // myCats := make([]string, 5, 10)
    
    // infering the length and type
    myCats := []string{"Milo", "Eva", "Torchy"}

    fmt.Printf("Length is: %d.\nCapacity is: %d",
        len(myCats), cap(myCats))

}