package main

import (
	"fmt"
)

func main() {
	var n, m int
	fmt.Scan(&n, &m) // Считываем размеры массива

	// Создаем двумерный массив
	A := make([][]int, n)
	for i := 0; i < n; i++ {
		A[i] = make([]int, m)
		for j := 0; j < m; j++ {
			fmt.Scan(&A[i][j]) // Считываем элементы строки
		}
	}

	// Инициализация переменных для поиска максимальной суммы
	maxSum := 0
	rowIndex := 0

	// Перебираем строки и находим максимальную сумму
	for i := 0; i < n; i++ {
		sum := 0
		for j := 0; j < m; j++ {
			sum += A[i][j] // Считаем сумму строки
		}
		// Если текущая сумма больше maxSum или равна maxSum, но индекс меньше
		if sum > maxSum || (sum == maxSum && i < rowIndex) {
			maxSum = sum
			rowIndex = i
		}
	}

	// Вывод результата
	fmt.Println(maxSum, rowIndex)
}
