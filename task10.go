package main

import (
    "fmt"
    "sort"
)

func main() {
    values := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
    groups := map[int][]float64{}

    for _, value := range values {
        key := int(0.1 * value) * 10
        groups[key] = append(groups[key], value)
    }

    keys := []int{}
    for key := range groups {
        keys = append(keys, key)
    }

    sort.Ints(keys)
    for _, key := range keys {
        fmt.Printf("%v: %v\n", key, groups[key])
    }
}
