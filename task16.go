package main 

import (
    "fmt"
)

func QuickSort(values []int) {
    n := len(values)
    if n < 2 { // Если элементов меньше двух, делать нечего
        return
    }

    minValue := values[0]
    maxValue := values[0]

    for _, value := range values {
        if value < minValue {
            minValue = value
        }

        if value > maxValue {
            maxValue = value
        }
    }

    if minValue == maxValue { // Если всем элементы равны, делать нечего
        return
    }

    pivot := (minValue + maxValue) / 2 // Выбираем разделитель (pivot)

    // Делим массив на две части
    i, j := 0, n - 1
    for i <= j {
        if values[i] <= pivot {
            i++
        } else if values[j] > pivot {
            j--
        } else {
            values[i], values[j] = values[j], values[i]
            i++
            j--
        }
    }

    // Запускаемся рекурсивно
    QuickSort(values[:i])
    QuickSort(values[j + 1:])
}

func main() {
    a := []int{3, 3, 1, 2, 2, 2, 5, 4, 7, 6}
    QuickSort(a)

    fmt.Println(a)
}
