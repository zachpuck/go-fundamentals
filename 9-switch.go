//switch
// default statement can go anywhere
// switch and case types must match
// implicit breaks
// requires explicit fallthrough
// case statements can have multiple comma-seperated values

package main

import (
    "fmt"
    "math/rand"
    "time"
)

func main() {
    switch "kubernetes" {
    case "go":
        fmt.Println("\nHere are some recommended" +
            " go courses...")
    case "kubernetes":
        fmt.Println("\nHere are some recommended" +
            " kubernetes courses...")
        fallthrough
    case "docker":
        fmt.Println("\nHere are some recommended" +
            " docker courses...")
    default:
        fmt.Println("\nNot sure where to start?" +
            " check out javascript")
    }

    randomSwitch()
}

func randomSwitch() {

    switch tmpNum := random(); tmpNum {
    case 0, 2, 4, 6, 8:
        fmt.Println("We got an even number -", tmpNum)
    case 1, 3, 5, 7, 9:
        fmt.Println("we got an odd number -", tmpNum)
    }
}

func random() int {
    rand.Seed(time.Now().Unix())
    return rand.Intn(10)
}