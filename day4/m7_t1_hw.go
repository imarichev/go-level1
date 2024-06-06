/*
Сформировать данные для отправки заказа из
магазина по накладной и вывести на экран:
1) Наименование товара (минимум 1, максимум 100)
2) Количество (только числа)
3) ФИО покупателя (только буквы)
4) Контактный телефон (10 цифр)
5) Адрес(индекс(ровно 6 цифр), город, улица, дом, квартира)

Эти данные не могут быть пустыми.
Проверить правильность заполнения полей.

реализовать несколько методов у типа "Накладная"

createReader == NewReader
*/

package main

import (
	"fmt"

	"example.com/invoice"
)

func main() {
	invoice := new(invoice.Invoice)
	invoice.ReadInvoice()
	errFields, err := invoice.CheckInvoice()
	for err != nil {
		fmt.Println("Накладная заполнена некорректно, исправьте следующие поля:")
		fmt.Print(err)
		invoice.ReadInvoice(errFields...)
		errFields, err = invoice.CheckInvoice()
	}

	invoice.Print()
}
