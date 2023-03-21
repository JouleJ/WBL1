package main

import (
    "fmt"
    "time"
)

func Sleep(duration time.Duration) {
    t := time.NewTimer(duration)
    <-t.C
}

func main() {
    Sleep(2 * time.Second)
    fmt.Println("FINISHED")
}
