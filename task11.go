package main

import (
    "fmt"
)

type Set[T comparable] map[T]bool

func Intersect[T comparable](firstSet Set[T], secondSet Set[T]) Set[T] {
    result := Set[T]{}

    for key := range firstSet {
        _, ok := secondSet[key]
        if ok {
            result[key] = true
        }
    }

    return result
}

func main() {
    s1 := Set[string]{"a": true, "b": true, "c": true}
    s2 := Set[string]{"z": true, "b": true, "a": true, "d": true}

    s := Intersect(s1, s2)
    fmt.Println(s)
}
