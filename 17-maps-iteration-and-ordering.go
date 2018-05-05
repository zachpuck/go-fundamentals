// maps are reference types -> passed to functions by reference (like slices)
// changes made by functions visible to caller
// unsafe for concurrency - it is not defined what happens to them when you read and write at the same time
// cheap to pass - just pointers instead of actual data
// specify size for large maps - to improve performance "make(map[<keyType>]<valueType>, size)"
package main

import (
    "fmt"
)

func main() {
    
    testMap := map[string]int{"A": 1, "B": 2, "C": 3, "D": 4, "E": 5}

    // note: maps are unordered lists
    for key, value := range testMap {
        fmt.Printf("\nKey is: %v Value is: %v\n", key, value)
    }

    // print specific maps[key] value
    fmt.Println(testMap["C\n"])

    // update a maps["key"] value
    testMap["A"] = 100

    // insert new item into map (same as updating)
    testMap["F"] = 1900

    fmt.Println(testMap)

    // deleting a key:value from the map
    delete(testMap, "F")

    fmt.Println(testMap)
}