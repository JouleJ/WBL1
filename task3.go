package main

import (
    "fmt"
    "sync"
)

func main() {
    values := []int{2, 4, 6, 8, 10}
    sum := 0
    wg := sync.WaitGroup{}
    sumMutex := sync.Mutex{}

    wg.Add(len(values))
    for _, value := range values {
        func (value int) {
            squaredValue := value * value

            sumMutex.Lock()
            sum += squaredValue
            sumMutex.Unlock()

            wg.Done()
        }(value)
    }

    wg.Wait()
    fmt.Printf("sum = %v\n", sum)
}
