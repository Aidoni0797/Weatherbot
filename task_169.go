package main

import "fmt"

func main() {
	var prev, current, count int

	// Считываем первый элемент последовательности
	fmt.Scan(&prev)

	for {
		fmt.Scan(&current) // Считываем следующий элемент
		if current == 0 {
			break // Завершаем, если встречаем 0
		}

		// Проверяем, изменился ли знак (один из чисел положительное, другое отрицательное)
		if (prev > 0 && current < 0) || (prev < 0 && current > 0) {
			count++ // Увеличиваем счётчик при смене знака
		}

		prev = current // Обновляем значение предыдущего элемента
	}

	fmt.Println(count) // Выводим результат
}
