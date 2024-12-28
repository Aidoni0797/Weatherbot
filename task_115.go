package main

import (
	"fmt"
)

func maxSum(K2, K3, K5, K6 int) int {
	// Сначала создаём как можно больше чисел 256
	count256 := min(K2, min(K5, K6))
	K2 -= count256
	K5 -= count256
	K6 -= count256

	// Затем создаём как можно больше чисел 32
	count32 := min(K2, K3)

	// Считаем сумму
	totalSum := count256*256 + count32*32
	return totalSum
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	// Ввод данных
	var K2, K3, K5, K6 int
	fmt.Scan(&K2)

	fmt.Scan(&K3)

	fmt.Scan(&K5)

	fmt.Scan(&K6)

	// Вычисление и вывод результата
	result := maxSum(K2, K3, K5, K6)
	fmt.Println(result)
}
