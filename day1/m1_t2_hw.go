/*
Задача № 2. Получить реверсную запись трехзначного числа
Пример: 
вход: 346, выход: 643
вход: 100, выход: 001
*/

package main

import "fmt"

func main() {
	var number int
	fmt.Print("Введите трехзначное число: ")
	fmt.Scanf("%d\n", &number)

	if number % 1000 != number || number / 100 == 0 {
		fmt.Println("Вы ввели не трехзначное число!")
		return
	}

	first := number / 100
	second := (number % 100) / 10
	third := number % 10

	fmt.Printf("Реверсивная запись: %d%d%d\n", third, second, first)
}