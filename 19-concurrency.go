// creating multiple processes that execute independently
// concurrency is about dealing with lots of things at once. Parallelism is about doing lots of things at once.
// goroutines vs OS Threads
// - ligher weight
// - go manages goroutines (simpler for programmers)
// - less switching
// - faster start-up times
// - safe communication
// actor model(go routines) - communicating sequential processes(CSP)
// channel(saff communication between goroutines) - one goroutine puts data on the channel and another goroutine picks it off
package main

import (
    "fmt"
    "time"
    "sync"
    "runtime"
)

func main() {
    
    // to run parallel
    runtime.GOMAXPROCS(2)

    // using waitGroup to allow the goroutines to complete before the main function complets
    var waitGrp sync.WaitGroup
    waitGrp.Add(2)

    go func() { // adding go to the function causes it to become a goroutine
        defer waitGrp.Done()
        
        time.Sleep(5 * time.Second)
        fmt.Println("Hello")
    }() // <-- self executing

    go func() {
        defer waitGrp.Done()

        fmt.Println("Bots")
    }()

    waitGrp.Wait()
}