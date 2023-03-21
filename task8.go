package main

import (
    "fmt"
)

func SetBit(value int64, bitIndex int64) int64 {
    return value | (int64(1) << bitIndex)
}

func ClearBit(value int64, bitIndex int64) int64 {
    return value & ^(int64(1) << bitIndex)
}

func main() {
    v := int64(0)

    v = SetBit(v, 1)
    v = SetBit(v, 2)
    fmt.Printf("v = %v\n", v)

    v = SetBit(v, 0)
    v = ClearBit(v, 1)
    fmt.Printf("v = %v\n", v)
}
