package main

import (
    "fmt"
)

func main() {
    
    mySlice := []int{1,2,3,4,5,6,7,8,9,10}
    fmt.Println(mySlice[4])
    
    mySlice[1] = 0
    fmt.Println(mySlice) // Referencing a slice by its variable name references the entire slice

    sliceOfSlice := mySlice[2:5] //gives values from index 2, 3, and 4
    fmt.Println(sliceOfSlice)

    // slices are flexible
    fmt.Println("working with 'append()'")

    expandingMySlice := make([]int, 1, 4)
    fmt.Printf("Lenght is: %d, Capacity is: %d",
        len(expandingMySlice), cap(expandingMySlice))
    
    for i := 1; i < 17; i++ {
        expandingMySlice = append(expandingMySlice, i)
        fmt.Printf("\nCapacity is: %d",
            cap(expandingMySlice))
    }

    // for range loops iterate slices
    // for range returns two values - 
    // -- index
    // -- data
    // index is disregarded by assigning it to the underscore
    for _, i := range mySlice {
        fmt.Println("for range loop:", i)
    }

    // can append() slices to slices with the ellipses (as long as the types are equivalent)
    newSlice := []int{10, 20, 30}
    mySlice = append(mySlice, newSlice...) //the appends the elements the newSlice holds to the mySlice
    fmt.Println(mySlice)
}