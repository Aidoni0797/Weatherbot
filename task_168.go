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

		if current > prev {
			count++ // Увеличиваем счётчик, если текущий элемент больше предыдущего
		}

		prev = current // Обновляем значение предыдущего элемента
	}

	fmt.Println(count) // Выводим результат
}
