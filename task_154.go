package main

import "fmt"

func main() {
	var N int
	fmt.Scan(&N) // Ввод числа N

	k := 0
	powerOfTwo := 1 // Начинаем с 2^0 = 1

	for powerOfTwo < N {
		powerOfTwo *= 2 // Умножаем текущее значение на 2
		k++
	}

	fmt.Println(k) // Выводим минимальное значение k
}
