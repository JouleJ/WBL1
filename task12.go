package main

import (
    "fmt"
)

func main() {
    strs := []string{"cat", "cat", "dog", "cat", "tree"}
    unique := map[string]bool{}

    for _, str := range strs {
        unique[str] = true
    }

    fmt.Println(unique)
}
