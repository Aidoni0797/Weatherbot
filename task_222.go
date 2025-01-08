package main

import (
	"fmt"
)

func main() {
	// Ввод количества элементов массива
	var N int
	fmt.Scan(&N)

	// Ввод элементов массива
	arr := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scan(&arr[i])
	}

	// Инициализация переменных для поиска минимального элемента
	minIndex := 0
	for i := 1; i < N; i++ {
		if arr[i] < arr[minIndex] {
			minIndex = i
		}
	}

	// Вывод индекса минимального элемента
	fmt.Println(minIndex)
}
