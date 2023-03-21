package main

import (
    "fmt"
    "reflect"
)

func DiscriminateType(value interface{}) reflect.Kind {
    t := reflect.TypeOf(value)
    return t.Kind()
}

func main() {
    fmt.Println(DiscriminateType(10))
    fmt.Println(DiscriminateType("hello"))
    fmt.Println(DiscriminateType(true))
    fmt.Println(DiscriminateType(make(chan int)))
}
