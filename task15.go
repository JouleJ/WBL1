package main

import (
    "fmt"
    "strings"
)

/* 
    Строки в Golang являются иммутабельными.
    За счёт этого при копировании строки копия не создаётся,
    то есть строка копируется по указателю.
    Аналогично при взятии подстроки.
    Однако это означает, что пока жива подстрока, не может умиреть строка.
    Таким образом ради короткой подстроки в 100 байт мы должны держать килобайт лишних данных.
    Чтобы избежать этого, сделаем глубокое копирование по значению.
*/

var justString string

func createHugeString(size int) string {
    runes := make([]rune, 0, size)
    for i := 0; i < size; i++ {
        runes = append(runes, 'X')
    }

    return string(runes)
}

func someFunc() {
  v := createHugeString(1 << 10)
  justString = strings.Clone(v[:100])
}

func main() {
  someFunc()
  fmt.Println(justString)
}

