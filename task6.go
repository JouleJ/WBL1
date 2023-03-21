package main

import (
    "context"
    "fmt"
    "sync"
    "time"
)

func main() {
    wg := sync.WaitGroup{}

    ch1 := make(chan int, 1)
    wg.Add(1)
    go func() {
        for {
            v, ok := <-ch1
            if !ok {
                break
            }

            fmt.Printf("First goroutine recevied value %v\n", v)
        }

        wg.Done()
        fmt.Println("First goroutine finished")
    }()

    ch2 := make(chan int, 1)
    done := make(chan struct{}, 0)
    wg.Add(1)
    go func() {
        Outer:
        for {
            select {
            case value := <-ch2:
                fmt.Printf("Second goroutine recevied value %v\n", value)
            case <-done:
                break Outer
            }
        }

        wg.Done()
        fmt.Println("Second goroutine finished")
    }()

    ch3 := make(chan int, 1)
    ctx, cancel := context.WithCancel(context.Background())
    go func(ctx context.Context) {
        Outer:
        for {
            select {
            case value := <-ch3:
                fmt.Printf("Third goroutine recevied value %v\n", value)
            case <-ctx.Done():
                break Outer
            }
        }

        wg.Done()
        fmt.Println("Third goroutine finished")
    }(ctx)

    ch1 <- 0
    ch1 <- 1
    ch2 <- 2
    ch3 <- 3
    ch2 <- 4
    ch1 <- 5
    ch3 <- 6

    time.Sleep(time.Second)

    close(ch1)
    done <- struct{}{}
    cancel()

    wg.Wait()
    fmt.Println("Main finished")
}
