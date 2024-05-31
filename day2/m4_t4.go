/*
Задача №4. Шахматная доска
Вход: размер шахматной доски, от 0 до 20
Выход: вывести на экран эту доску, заполняя поля нулями и единицами

Пример:
Вход: 5
Выход:
    0 1 0 1 0
    1 0 1 0 1
    0 1 0 1 0
    1 0 1 0 1
    0 1 0 1 0
*/
package main

import (
    "example.com/readers"
    "fmt"
)

func main() {
    fmt.Println("Укажите размер доски")
    size := readers.ReadIntMinMax(0, 20)
    printDesk(size)
}

func printDesk(size int){
    for i:=0; i < size; i++ {
        for j:=0; j< size; j++ {
            fmt.Printf("%d ",(i + j) % 2)
        }
        fmt.Println()
    }
}