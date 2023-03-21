package main

import (
    "fmt"
    "sync"
)

type Counter struct {
    mutex sync.Mutex
    counter int
}

func NewCounter() *Counter {
    return &Counter{mutex: sync.Mutex{}, counter: 0}
}

func (c *Counter) Increment() {
    c.mutex.Lock()
    defer c.mutex.Unlock()

    c.counter++
}

func (c *Counter) Get() int {
    c.mutex.Lock()
    defer c.mutex.Unlock()

    result := c.counter
    return result
}

func main() {
    c := NewCounter()
    wg := sync.WaitGroup{}

    wg.Add(1)
    go func() {
        for i := 0; i < 10; i++ {
            c.Increment()
        }

        wg.Done()
    }()

    wg.Add(1)
    go func() {
        for i := 0; i < 5; i++ {
            c.Increment()
        }

        wg.Done()
    }()

    wg.Add(1)
    go func() {
        for i := 0; i < 5; i++ {
            c.Increment()
        }

        wg.Done()
    }()

    wg.Wait()
    fmt.Printf("Main finished, c.Get() = %v\n", c.Get())
}
