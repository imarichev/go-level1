package invoice

import (
	"fmt"
	"strings"
	"unicode"
	"unicode/utf8"

	"example.com/readers"
)

type Invoice struct {
	PackageName   string
	Quantity      int
	CustomerName  string
	ContactNumber string
	Address
}

type Address struct {
	PostalIndex string
	City        string
	Street      string
	Building    int
	Appartment  int
}

func (invoice *Invoice) ReadInvoice(fields ...string) {
	if len(fields) == 0 {
		fields = getFields()
	}
	for _, field := range fields {
		switch field {
		case "PackageName":
			fmt.Println("Введите наименование товара: ")
			invoice.PackageName = readers.ReadString()
		case "Quantity":
			fmt.Println("Введите количество: ")
			invoice.Quantity = readers.ReadInt()
		case "CustomerName":
			fmt.Println("Введите имя покупателя: ")
			invoice.CustomerName = readers.ReadString()
		case "ContactNumber":
			fmt.Println("Введите номер телефона: ")
			invoice.ContactNumber = readers.ReadString()
		case "PostalIndex":
			fmt.Println("Введите почтовый индекс: ")
			invoice.PostalIndex = readers.ReadString()
		case "City":
			fmt.Println("Введите город: ")
			invoice.City = readers.ReadString()
		case "Street":
			fmt.Println("Введите улицу: ")
			invoice.Street = readers.ReadString()
		case "Building":
			fmt.Println("Введите номер дома: ")
			invoice.Building = readers.ReadInt()
		case "Appartment":
			fmt.Println("Введите номер квартиры: ")
			invoice.Appartment = readers.ReadInt()
		}
	}
}

func (invoice Invoice) CheckInvoice() ([]string, error) {
	var errInfo string
	errFields := make([]string, 8)
	fields := getFields()
	for _, field := range fields {
		switch field {
		case "PackageName":
			if utf8.RuneCountInString(invoice.PackageName) < 1 || utf8.RuneCountInString(invoice.PackageName) > 100 {
				errInfo += "\tдлина названия товара должна быть от 1 до 100 символов\n"
				errFields = append(errFields, field)
			}
		case "Quantity":
			if invoice.Quantity <= 0 {
				errInfo += "\tколичество товара должно быть положительным\n"
				errFields = append(errFields, field)
			}
		case "CustomerName":
			if len(invoice.CustomerName) == 0 {
				errInfo += "\tимя не может быть пустым\n"
				errFields = append(errFields, field)
			}
			nameFields := strings.Fields(invoice.CustomerName)
			for _, nameField := range nameFields {
				for _, r := range nameField {
					if !unicode.IsLetter(r) {
						errInfo += "\tимя покупателя может содержать только буквы\n"
						errFields = append(errFields, field)
						break
					}
				}
			}
		case "ContactNumber":
			correct := true
			if len(invoice.ContactNumber) != 10 {
				correct = false
			}
			for _, r := range invoice.ContactNumber {
				if !strings.Contains("0123456789", string(r)) {
					correct = false
				}
			}
			if !correct {
				errInfo += "\tномер телефона должен быть длинной 10 символов и содержать только цифры\n"
				errFields = append(errFields, field)
			}
		case "PostalIndex":
			correct := true
			if len(invoice.PostalIndex) != 6 {
				correct = false
			}
			for _, r := range invoice.PostalIndex {
				if !strings.Contains("0123456789", string(r)) {
					correct = false
				}
			}
			if !correct {
				errInfo += "\tпочтовый индекс должен быть длинной 6 символов и содержать только цифры\n"
				errFields = append(errFields, field)
			}
		case "City":
			if len(invoice.City) == 0 {
				errInfo += "\tгород не может быть пустым\n"
				errFields = append(errFields, field)
			}
		case "Street":
			if len(invoice.Street) == 0 {
				errInfo += "\tулица не может быть пустой\n"
				errFields = append(errFields, field)
			}
		case "Building":
			if invoice.Building <= 0 {
				errInfo += "\tномер дома должен быть положительным числом\n"
				errFields = append(errFields, field)
			}
		case "Appartment":
			if invoice.Appartment <= 0 {
				errInfo += "\tномер квартиры должен быть положительным числом\n"
				errFields = append(errFields, field)
			}
		}
	}
	if len(errInfo) > 0 {
		return errFields, fmt.Errorf(errInfo)
	} else {
		return nil, nil
	}
}

func (invoice Invoice) Print() {
	fmt.Println()
	fmt.Println("НАКЛАДНАЯ:")
	fmt.Print("============================================\n\n")
	fmt.Printf("Наименование товара : %s\n", invoice.PackageName)
	fmt.Printf("Количество товара: %d\n", invoice.Quantity)
	fmt.Printf("Имя покупателя: %s\n", invoice.CustomerName)
	fmt.Printf("Контактный номер: %s\n", invoice.ContactNumber)
	fmt.Println("Адрес:")
	fmt.Printf("\tПочтовый индекс: %s\n", invoice.PostalIndex)
	fmt.Printf("\tГород: %s\n", invoice.City)
	fmt.Printf("\tУлица: %s\n", invoice.Street)
	fmt.Printf("\tНомер дома: %d\n", invoice.Building)
	fmt.Printf("\tНомер квартиры: %d\n", invoice.Appartment)
	fmt.Print("\n============================================\n")
}

func getFields() []string {
	return []string{
		"PackageName",
		"Quantity",
		"CustomerName",
		"ContactNumber",
		"PostalIndex",
		"City",
		"Street",
		"Building",
		"Appartment",
	}
}
