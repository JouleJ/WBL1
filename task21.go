package main

import (
    "fmt"
)

// Пусть у нас есть два интерфейса

type InterfaceA interface {
    Request(i int, s string)
}

type InterfaceB interface {
    Request(s string, i int)
}

// И также есть функция, которая принимает второй интерфейс
func UseInterfaceB(ib InterfaceB) {
    ib.Request("hello", 1)
    ib.Request("world", 2)
}

// Также у нас есть реализация первого интерфейса
type implA struct {}

func NewA() InterfaceA {
    return &implA{}
}

func (ia *implA) Request(i int, s string) {
    fmt.Printf("A.Request(i=%v, s=%v)\n", i, s)
}

// Теперь допустим, что мы хотим вызывать нашу функцию от второго интерфейса.
// Напрямую это сделать нельзя.
// Поэтому напишем, адаптер, переводящий вызовы одного метода в другой.

type adapterAtoB struct {
    ia InterfaceA
}

func AdaptAtoB(ia InterfaceA) InterfaceB {
    return &adapterAtoB{ia: ia}
}

func (a2b *adapterAtoB) Request(s string, i int) {
    a2b.ia.Request(i, s)
}

func main() {
    a := NewA()
    a2b := AdaptAtoB(a)
    UseInterfaceB(a2b)
}
