package main

import (
    "fmt"
    "sync"
)

func main() {
    values := []int{2, 4, 6, 8, 10}
    wg := sync.WaitGroup{}

    wg.Add(len(values))
    for _, value := range values {
        go func(value int) {
            squaredValue := value * value
            fmt.Printf("%v^2 = %v\n", value, squaredValue)
            wg.Done()
        }(value)
    }

    wg.Wait()
    fmt.Println("Finished")
}
