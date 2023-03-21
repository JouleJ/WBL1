package main

import (
    "fmt"
    "strings"
)

func ReverseWords(s string) string {
    words := strings.Fields(s)
    n := len(words)
    i, j := 0, n - 1
    for i < j {
        words[i], words[j] = words[j], words[i]
        i++
        j--
    }

    return strings.Join(words, " ")
}

func main() {
    s := "snow dog sun"
    fmt.Printf("%v -> %v\n", s, ReverseWords(s))
}
