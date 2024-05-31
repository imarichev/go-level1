package readers

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
)

func ReadIntMinMax(min int, max int) (int) {
	fmt.Printf("Введите целое число между %d и %d включительно: ", min, max)
	number := ReadInt()
	for number < min || number > max {
		fmt.Print("Вы вышли за пределы диапазона, попробуйте ещё раз: ")
		number = ReadInt()
	}
	return number
}

func ReadInt () (int) {
	
	data := ReadString()
	number, err := strconv.ParseInt(data, 10, 0)
	for err != nil {
		fmt.Print("Кажется вы ввели не число, попробуйте еще раз: ")
		data = ReadString()
		number, err = strconv.ParseInt(data, 10, 0)
	}
	return int(number)
}

func ReadString () (string) {
	reader := bufio.NewReader(os.Stdin)
	data, err := reader.ReadString('\n')
	for err != nil {
		fmt.Print("Что-то пошло не так, попробуйте еще раз: ")
		data, err = reader.ReadString('\n')
	}
	data = strings.TrimSpace(data)
	return data
}

func ReadInts () ([]int) {
	var result []int
	str := ReadString()
	parts := strings.Split(str, " ");
	for _, value := range parts {
		number, err := strconv.ParseInt(value, 10, 0)
		if err != nil {
			continue
		}
		result = append(result, int(number))
	}
	return result
}

func ReadIntsWithCount(count int) ([]int) {
	ints := ReadInts();
	for len(ints) != count{
		fmt.Println("Вы ввели некорректные данные, проверьте количество и корректность чисел и попробуйте еще раз:")
		ints = ReadInts()
	}
	return ints
}