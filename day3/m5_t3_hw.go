/*
Задача №2
Вход:
Пользователь должен ввести правильный пароль, состоящий из:
цифр,
букв латинского алфавита(строчные и прописные) и
специальных символов  special = "_!@#$%^&"

Всего 4 набора различных символов.
В пароле обязательно должен быть хотя бы один символ из каждого набора.
Длина пароля от 8(мин) до 15(макс) символов.
Максимальное количество попыток ввода неправильного пароля - 5.
Каждый раз выводим номер попытки.
*Желательно выводить пояснение, почему пароль не принят и что нужно исправить.

digits = "0123456789"
lowercase = "abcdefghiklmnopqrstvxyz"
uppercase = "ABCDEFGHIKLMNOPQRSTVXYZ"
special = "_!@#$%^&"

Выход:
Написать, что ввели правильный пароль.

Пример:
хороший пароль -> o58anuahaunH!
хороший пароль -> aaaAAA111!!!
плохой пароль -> saucacAusacu8
*/

package main

import (
	"fmt"
	"log"
	"strings"

	"example.com/readers"
)

func main() {
	var count int
	fmt.Println("Введите пароль")
	fmt.Println("Пароль должен содержать: ")
	fmt.Println("\tХотя бы одну цифру")
	fmt.Println("\tХотя бы одну латинскую букву в нижнем регистре")
	fmt.Println("\tХотя бы одну латинскую букву в верхнем регистре")
	fmt.Println("\tХотя бы один специальный сивмол _!@#$%^&")
	fmt.Println("Общая длина пароля должны быть от 8 до 15 символов")
	fmt.Print(">>> ")

	password := readers.ReadString()
	err := checkPass(password)
	for err != nil && count < 4 {
		count++
		fmt.Println("Вы не выполнили следующие условия:")
		fmt.Print(err)
		fmt.Printf("Повторит ввод, попыток осталось: %d\n", 5-count)
		fmt.Print(">>> ")
		password = readers.ReadString()
		err = checkPass(password)
	}
	if count > 4 {
		log.Fatal("Условия не выполнены, и у вас больше не осталось попыток!")
	} else {
		fmt.Println("Пароль принят, спасибо!")
	}
}

func checkPass(password string) error {

	var errorString string
	reasons := map[string]string{
		"digits":    "В ващем пароле нет цифр",
		"lowercase": "В вашем пароле нет букв в нижнем регистре",
		"uppercase": "В вашем пароле нет букв в верхнем регистре",
		"special":   "В вашем пароле нет специальных символов",
		"length":    "Ваш пароль должен быть от 8 до 15 символов",
	}

	checks := map[string]string{
		"digits":    "0123456789",
		"lowercase": "abcdefghiklmnopqrstvxyz",
		"uppercase": "ABCDEFGHIKLMNOPQRSTVXYZ",
		"special":   "_!@#$%^&",
	}

	results := make(map[string]int)

	for checkName, symbols := range checks {
		if strings.ContainsAny(password, symbols) {
			results[checkName]++
		}
	}

	if len(password) > 7 && len(password) < 16 {
		results["length"]++
	}

	for checkName, reason := range reasons {
		_, ok := results[checkName]
		if !ok {
			errorString += "\t" + reason + "\n"
		}
	}

	if len(errorString) > 0 {
		return fmt.Errorf(errorString)
	}

	return nil

}
