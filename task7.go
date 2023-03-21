package main

import (
    "fmt"
    "sync"
)

/* Существует также готовая релизация sync.Map (которая использует более хитрый алгоритм).
   Но в задаче просят именно реализовать, а не использовать готовое. */

type ConcurrentMap struct {
    Mutex sync.Mutex
    Map map[int]int
}

func NewConcurrentMap() *ConcurrentMap {
    return &ConcurrentMap{Mutex: sync.Mutex{}, Map: map[int]int{}}
}

func (cm *ConcurrentMap) Write(key, value int) {
    cm.Mutex.Lock()
    defer cm.Mutex.Unlock()

    cm.Map[key] = value
}

func (cm *ConcurrentMap) Read(key int) (int, bool) {
    cm.Mutex.Lock()
    defer cm.Mutex.Unlock()

    value, ok := cm.Map[key]
    return value, ok
}

func main() {
    cm := NewConcurrentMap()

    cm.Write(0, 1)

    val, ok := cm.Read(0)
    if ok {
        fmt.Printf("cm[0] = %v\n", val)
    } else {
        fmt.Printf("cm[0] not found\n")
    }
}
