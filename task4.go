package main

import (
    "fmt"
    "os"
    "os/signal"
    "sync"
)

func main() {
    signalChannel := make(chan os.Signal, 1)
    signal.Notify(signalChannel, os.Interrupt)
    c := make(chan int, 1)
    wg := sync.WaitGroup{}

    var workerCount int
    fmt.Print("Enter worker count: ")
    fmt.Scan(&workerCount)

    wg.Add(workerCount)
    for workerId := 0; workerId < workerCount; workerId++ {
        go func(id int) {
            for value := range c {
                fmt.Printf("Worker-%v received value %v\n", id, value)
            }

            wg.Done()
            fmt.Printf("Worker-%v finished\n", id)
        }(workerId)
    }

Outer:
    for value := 0;; value++{
        select {
        case <-signalChannel:
            break Outer
        case c <- value:
            {}
        }
    }

    close(c)
    wg.Wait()
    fmt.Println("Finished")
}
