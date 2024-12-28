package main

import (
	"fmt"
)

func main() {
	var a, b float64
	var operation string

	// Считывание входных данных

	fmt.Scan(&a, &b)

	fmt.Scan(&operation)

	// Обработка операций
	switch operation {
	case "+":
		fmt.Println(a + b)
	case "-":
		fmt.Println(a - b)
	case "*":
		fmt.Println(a * b)
	case "/":
		if b == 0 {
			fmt.Println("На ноль делить нельзя!")
		} else {
			fmt.Println(a / b)
		}
	default:
		fmt.Println("Неверная операция")
	}
}
