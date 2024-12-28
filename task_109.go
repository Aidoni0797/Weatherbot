package main

import "fmt"

func main() {
	var month int
	// Считывание номера месяца
	fmt.Scan(&month)

	// Определение поры года по номеру месяца
	switch month {
	case 12, 1, 2:
		fmt.Println("Зима")
	case 3, 4, 5:
		fmt.Println("Весна")
	case 6, 7, 8:
		fmt.Println("Лето")
	case 9, 10, 11:
		fmt.Println("Осень")
	default:
		// В случае некорректного ввода (например, числа не от 1 до 12)
		fmt.Println("Некорректный месяц")
	}
}
