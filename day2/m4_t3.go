/*
Написать 3 функции.
Даны координаты трех точек(x1, y1, x2, y2, x3, y3), значения(целые) которых >= 0.
Первая функция проверяет, что можно построить треугольник по заданным точкам
Вторая функция вычисляет площадь треугольника.
Третья функция должна определить, является ли треугольник прямоугольным.
*/
package main

import (
	"example.com/readers"
	"fmt"
	"errors"
	"example.com/triangles"
	"log"
)

func main() {
	fmt.Println("Введите неотрицательные координаты трех точек(x1, y1, x2, y2, x3, y3) через пробел:")
	ints := readers.ReadIntsWithCount(6)
	err := checkPositiveness(ints)
	for err != nil {
		fmt.Println("Вы ввели отрицательное значение, повторите ввод:")
		ints = readers.ReadIntsWithCount(6)
		err = checkPositiveness(ints)
	}
	a, b, c := triangles.GetSides(ints[0], ints[1], ints[2], ints[3], ints[4], ints[5])
	err = triangles.IsDotsCorrect(a, b, c)
	if err != nil {
		log.Fatal("Треугольник нельзя построить")
	}else{
		fmt.Println("Треуголник можно построить")
	}
	s := triangles.Square(a, b, c)
	fmt.Printf("Площадь треугольника: %.2f\n", s)
	if(triangles.IsRight(a, b, c)){
		fmt.Println("Треугольник прямоугольный")
	}else{
		fmt.Println("Треугольник не прямоугольный")
	}
}

func checkPositiveness(numbers []int) (error){
	for _, value := range numbers {
		if value < 0 {
			return errors.New("negative")
		}
	}
	return nil
}