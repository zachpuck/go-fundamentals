package main

import (
    "fmt"
)

func main() {
    hikes := make(map[string]int)

    hikes["Snow Lake"] = 1
    hikes["Rattlesnake Ridge"] = 2

    winterHikes := map[string]int {
        "Gold Creek Pond": 1,
        "Iron Horse Trail": 1,
    }

    fmt.Printf("\nHikes: %v\nWinter Hikes: %v\n", 
        hikes, winterHikes)
}