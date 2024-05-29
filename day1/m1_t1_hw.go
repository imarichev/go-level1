/*
Задача №1
Вход:
    расстояние(50 - 10000 км),
    расход в литрах (5-25 литров) на 100 км и
    стоимость бензина(константа) = 48 руб

Выход: стоимость поездки в рублях
*/
package main

import "fmt"

func main() {
    var distance, gasConsumption float64

    const gasCost = 48

    fmt.Print("Введите растояние в диапазоне 50 - 10 000 (км): ")
    fmt.Scanf("%f\n", &distance)
    if distance < 50 || distance > 10000 {
        fmt.Println("Вы вышли за пределы допустимого диапазона!")
        return
    }

    fmt.Print("Введите расход на 100км диапазоне 5 - 25 (литров): ")
    fmt.Scanf("%f\n", &gasConsumption)
    if gasConsumption < 5 || gasConsumption > 25 {
        fmt.Println("Вы вышли за пределы допустимого диапазона!")
        return
    }

    fmt.Printf("Стоимость вашей поездки: %.2f\n", (distance / 100) * gasConsumption * gasCost )

}
