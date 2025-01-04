package main

import (
	"fmt"
)

// Функция для проверки, является ли число совершенным
func isPerfectNumber(n int) bool {
	sum := 0
	for i := 1; i < n; i++ {
		if n%i == 0 {
			sum += i
		}
	}
	return sum == n
}

func main() {
	count := 0
	num := 1
	var perfectNumbers []int

	// Ищем первые три совершенных числа
	for count < 3 {
		if isPerfectNumber(num) {
			perfectNumbers = append(perfectNumbers, num)
			count++
		}
		num++
	}

	// Выводим результат через запятую
	fmt.Println(perfectNumbers[0], ",", perfectNumbers[1], ",", perfectNumbers[2])
}
