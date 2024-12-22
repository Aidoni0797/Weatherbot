package main

import "fmt"

func main() {
	// Ввод данных
	var totalMinutes int
	fmt.Scan(&totalMinutes)

	// Перевод минут в часы и оставшиеся минуты
	hours := totalMinutes / 60
	minutes := totalMinutes % 60

	// Вывод результата
	fmt.Printf("%d часов %d минут\n", hours, minutes)
}
