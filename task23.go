package main

import (
    "fmt"
)

func RemoveKthElement(slice []int, k int) []int {
    n := len(slice)
    for i := k; i < n - 1; i++ {
        slice[i] = slice[i + 1]
    }
    return slice[:n - 1]
}

func main() {
    slice := []int{1, 2, 3, 4}
    slice = RemoveKthElement(slice, 1)
    fmt.Println(slice)
}
