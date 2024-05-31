/*
Написать функцию.
Входные аргументы функции: количество бутылок от 0 до 200.
Функция должна вернуть количество и слово бутыл<?> с правильным окончанием.
Пример :
In: 5
Out: 5 бутылок
    
In: 1
Out: 1 бутылка
    
In: 22
Out: 22 бутылки
*/
package main

import (
	"example.com/bottles"
	"example.com/readers"
)

func main() {
	number := readers.ReadIntMinMax(0, 200)
	bottles.Humanize(i)
}