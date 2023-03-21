package main

import (
    "fmt"
    "math/big"
)

func main() {
    a := big.NewInt(0)
    b := big.NewInt(1)

    for i := 0; i < 100; i++ {
        t := &big.Int{}
        t.Add(a, b)
        a.Set(b)
        b.Set(t)
    }

    a, b = b, a

    t := &big.Int{}
    fmt.Printf("a=%v, b=%v\n", a, b)

    t.Add(a, b)
    fmt.Printf("a+b=%v\n", t)

    t.Sub(a, b)
    fmt.Printf("a-b=%v\n", t)

    t.Mul(a, b)
    fmt.Printf("a*b=%v\n", t)

    t.Div(a, b)
    fmt.Printf("a/b=%v\n", t)
}
