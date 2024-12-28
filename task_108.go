package main

import (
	"fmt"
)

func main() {
	var month int
	// Считывание номера месяца
	fmt.Scan(&month)

	// Определяем количество дней в месяце
	var days int
	switch month {
	case 1, 3, 5, 7, 8, 10, 12:
		days = 31
	case 4, 6, 9, 11:
		days = 30
	case 2:
		days = 29 // Для задачи считаем, что в феврале всегда 29 дней
	default:
		// В случае неправильного ввода, но по условиям задачи, такого не должно быть
		days = 0
	}

	// Выводим количество дней
	fmt.Println(days)
}
