package main

import (
    "fmt"
)

func ReverseUnicode(s string) string {
    runes := []rune(s)
    n := len(runes)
    i, j := 0, n - 1
    for i < j {
        runes[i], runes[j] = runes[j], runes[i]
        i++
        j--
    }

    return string(runes)
}

func main() {
    s := "главрыба"

    fmt.Printf("%v -> %v\n", s, ReverseUnicode(s))
}
