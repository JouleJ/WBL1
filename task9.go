package main

import (
    "fmt"
    "sync"
)

func Pipeline(inputChan <-chan int, outputChan chan<- int, wg *sync.WaitGroup) {
    for value := range inputChan {
        outputChan <- 2 * value
    }

    close(outputChan)
    wg.Done()
}

func Reader(inputChan <-chan int, wg *sync.WaitGroup) {
    for value := range inputChan {
        fmt.Printf("Read value %v\n", value)
    }

    wg.Done()
}

func main() {
    wg := sync.WaitGroup{}
    ch1, ch2 := make(chan int), make(chan int)

    wg.Add(1)
    go Pipeline(ch1, ch2, &wg)

    wg.Add(1)
    go Reader(ch2, &wg)

    for i := 0; i < 5; i++ {
        ch1 <- i
    }

    close(ch1)
    wg.Wait()

    fmt.Println("Main finished")
}
