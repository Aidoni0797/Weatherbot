package main

import (
	"fmt"
)

func main() {
	var num float64

	// Считываем вещественное число
	fmt.Scan(&num)

	// Проверяем, является ли число нулём
	if num == 0 {
		fmt.Println("Обратного числа не существует")
	} else {
		// Выводим обратное число
		fmt.Printf("%.6f\n", 1/num)
	}
}
