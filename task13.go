package main

import (
    "fmt"
)

func main() {
    a, b := 133, 456
    fmt.Printf("a=%v, b=%v\n", a, b)

    a ^= b
    b ^= a
    a ^= b

    fmt.Printf("a=%v, b=%v\n", a, b)
}
