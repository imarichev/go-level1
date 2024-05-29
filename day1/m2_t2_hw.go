/*
Задача № 4. Проверить, является ли четырехзначное число палиндромом
Пример:
Вход: 1221  Выход: 1221 - палиндром
Вход: 1234  Выход: 1234 - не палиндром
*/

package main

import "fmt"

func main() {
	var number int
	fmt.Print("Введите четырехзначное число: ")
	fmt.Scanf("%d\n", &number)

	if number % 10000 != number || number / 1000 == 0 {
		fmt.Println("Вы ввели не четырехзначное число!")
		return
	} 

	first := number / 1000
	second := (number % 1000) / 100
	third := (number % 100) / 10
	fourth := (number % 10)

	if first == fourth && second == third {
		fmt.Printf("%d - палиндром\n", number)
	} else {
		fmt.Printf("%d - не палиндром\n", number)
	}
}