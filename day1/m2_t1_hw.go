/*
Задача № 3. Вывести на экран в порядке возрастания три введенных числа
Пример:
Вход: 1, 9, 2
Выход: 1, 2, 9
*/

package main

import "fmt"

func main() {
	var first, second, third int

	fmt.Print("Введите три числа через пробел: ")
	fmt.Scanf("%d %d %d\n", &first, &second, &third)

	fmt.Print("Числа в порядке возрастания: ")
	if first < second && first < third{
		fmt.Printf("%d ", first)
		if second < third {
			fmt.Printf("%d %d\n", second, third)
		}else{
			fmt.Printf("%d %d\n", third, second)
		}
	}else if second < third {
		fmt.Printf("%d ", second)
		if first < third {
			fmt.Printf("%d %d\n", first, third)
		} else {
			fmt.Printf("%d %d\n", third, first)
		}
	} else {
		fmt.Printf("%d ", third)
		if first < second {
			fmt.Printf("%d %d\n", first, second)
		} else {
			fmt.Printf("%d %d\n", second, first)
		}
	}
		
}