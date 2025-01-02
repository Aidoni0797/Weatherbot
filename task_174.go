package main

import (
	"fmt"
	"math"
)

// Функция для разложения числа на простые множители
func primeFactors(N int) []int {
	var factors []int

	// Пробное деление на 2
	for N%2 == 0 {
		factors = append(factors, 2)
		N /= 2
	}

	// Пробное деление на все нечётные числа от 3 до √N
	for i := 3; i <= int(math.Sqrt(float64(N))); i += 2 {
		for N%i == 0 {
			factors = append(factors, i)
			N /= i
		}
	}

	// Если оставшееся число больше 1, то это простое число
	if N > 1 {
		factors = append(factors, N)
	}

	return factors
}

func main() {
	var N int
	fmt.Scan(&N)

	// Разлагаем число N на простые множители
	factors := primeFactors(N)

	// Выводим результат
	for i, factor := range factors {
		if i > 0 {
			fmt.Print(" ")
		}
		fmt.Print(factor)
	}
	fmt.Println()
}
