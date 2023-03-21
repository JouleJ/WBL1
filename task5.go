package main

import (
    "fmt"
    "time"
    "sync"
)

const (
    N = 3
)

func main() {
    c := make(chan int, 1)
    t := time.NewTimer(N * time.Second)
    wg := sync.WaitGroup{}

    wg.Add(1)
    go func(c chan<- int) {
        Outer:
        for value := 0;; value++ {
            select {
            case c <- value:
                {}
            case <- t.C:
                break Outer
            }
        } 

        close(c)
        wg.Done()
    }(c)

    wg.Add(1)
    go func() {
        for value := range c {
            fmt.Printf("Recieved value: %v\n", value)
        }

        wg.Done()
    }()

    wg.Wait()
    fmt.Println("Finished")
}
