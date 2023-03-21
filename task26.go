package main

import (
    "fmt"
    "unicode"
)

func CheckCharactersUnique(s string) bool {
    isPresent := map[rune]bool{}

    for _, r := range s {
        rToLower := unicode.ToLower(r)

        _, ok := isPresent[rToLower]
        if ok {
            return false
        }

        isPresent[rToLower] = true
    }

    return true
}

func main() {
    strs := []string{"abcd", "abCdefAaf", "aabcd"}
    for _, str := range strs {
        fmt.Printf("%v -> %v\n", str, CheckCharactersUnique(str))
    }
}
