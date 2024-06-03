/*
Задача №1
Написать функцию, которая расшифрует строку.
code = "220411112603141304"
Каждые две цифры - это либо буква латинского алфавита в нижнем регистре либо пробел.
Отчет с 00 -> 'a' и до 25 -> 'z', 26 -> ' '(пробел).
Вход: строка из цифр. Выход: Текст.
Проверка работы функции выполняется через вторую строку.

codeToString(code) -> "???????'
*/

package main

import (
	"errors"
	"fmt"
	"log"
	"strconv"

	"example.com/readers"
)

func main() {
	fmt.Println("Введите закодированную строку:")
	code := readers.ReadString()
	answer, err := decode(code)
	if err != nil {
		log.Fatal("Вы ввели некорректную строку!", err)
	}
	fmt.Println(answer)
}

func decode(code string) (string, error) {

	var symbol, answer string
	key := "abcdefghijklmnopqrstuvwxyz "

	for _, letter := range code {
		symbol += string(letter)
		if len(symbol) < 2 {
			continue
		}
		number, err := strconv.ParseInt(symbol, 10, 0)
		if err != nil {
			return "", err
		}
		if number < 0 || number > 26 {
			return "", errors.New(fmt.Sprintf("Index out of range: %d", number))
		}
		answer += string(key[number])
		symbol = ""
	}

	return answer, nil
}
