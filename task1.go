package main

import (
    "fmt"
)

type Human struct {
    Name string
}

func (h *Human) Greet() {
    fmt.Printf("Hello, I am %v\n", h.Name)
}

type Action struct {
    Human
}

func main() {
    h := &Human{Name: "Oleg"}
    h.Greet()

    a := &Action{Human: Human{"Olga"}}
    a.Greet()
}
