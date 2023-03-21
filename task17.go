package main

import (
    "fmt"
)

func BinarySearch(values []int, key int) int {
    n := len(values)

    // Ответ находится где-то между l и r, если он вообще есть
    l, r := 0, n - 1
    for l <= r {
        m := l + (r - l) / 2 // Возмём середину
        if key < values[m] { // Если искомое число меньше, переходим в левую половину
            r = m - 1
        } else if key > values[m] { // Если искомое число больше, переходим в правую половину
            l = m + 1
        } else {
            return m // Иначе, середина это ответ
        }
    }

    return -1 // Если ничего не нашли, вернём -1
}

func main() {
    values := []int{1, 2, 3, 5, 6, 8, 10}

    fmt.Printf("BinarySearch(%v) = %v\n", 1, BinarySearch(values, 1))
    fmt.Printf("BinarySearch(%v) = %v\n", 2, BinarySearch(values, 2))
    fmt.Printf("BinarySearch(%v) = %v\n", 3, BinarySearch(values, 3))
    fmt.Printf("BinarySearch(%v) = %v\n", 4, BinarySearch(values, 4))
    fmt.Printf("BinarySearch(%v) = %v\n", 5, BinarySearch(values, 5))
    fmt.Printf("BinarySearch(%v) = %v\n", 6, BinarySearch(values, 6))
    fmt.Printf("BinarySearch(%v) = %v\n", 7, BinarySearch(values, 7))
}
