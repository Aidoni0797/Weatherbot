package main

import (
	"fmt"
)

func main() {
	var a, b, c int

	// Ввод трех целых чисел
	fmt.Scan(&a, &b, &c)

	// Определение количества совпадающих чисел
	if a == b && b == c {
		fmt.Println(3) // Все три числа совпадают
	} else if a == b || a == c || b == c {
		fmt.Println(2) // Совпадают любые два числа
	} else {
		fmt.Println(0) // Все числа различны
	}
}
