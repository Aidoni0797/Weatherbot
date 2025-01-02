package main

import "fmt"

func main() {
	var num, sum, count int

	for {
		fmt.Scan(&num) // Считываем число
		if num == 0 {
			break // Завершаем, если встречаем 0
		}

		sum += num // Суммируем числа
		count++    // Увеличиваем количество чисел
	}

	// Вычисляем среднее арифметическое
	average := float64(sum) / float64(count)
	fmt.Printf("%.6f\n", average) // Выводим результат с точностью до 6 знаков
}
