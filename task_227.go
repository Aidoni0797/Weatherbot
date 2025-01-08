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

	// Инициализация переменных для индексов минимального и максимального элементов
	minIndex, maxIndex := 0, 0
	for i := 1; i < N; i++ {
		if arr[i] < arr[minIndex] {
			minIndex = i
		}
		if arr[i] >= arr[maxIndex] {
			maxIndex = i
		}
	}

	// Меняем местами первый минимальный и последний максимальный элементы
	arr[minIndex], arr[maxIndex] = arr[maxIndex], arr[minIndex]

	// Вывод измененного массива
	for i := 0; i < N; i++ {
		fmt.Print(arr[i], " ")
	}
}
