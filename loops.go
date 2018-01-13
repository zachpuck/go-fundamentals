// for keyword
// expression
// - Blank expression = infinit loop
// can be Boolean expression
// can be a range
package main

import (
    "fmt"
    "time"
)

func main() {
    for timer := 10; timer >= 0; timer-- {
        if timer == 0 {
            fmt.Println("Liftoff!")
            break
        }
        fmt.Println("Countdown", timer)
        time.Sleep(1 * time.Second)
    }
}