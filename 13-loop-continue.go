package main

import (
    "time"
    "fmt"
)

func main(){
    for timer := 10; timer >= 0; timer-- {
        if timer % 2 == 0 {
            continue
        }
        fmt.Println("Countdown", timer)
        time.Sleep(1 * time.Second)
    } 
}