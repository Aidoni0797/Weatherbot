package main

import (
	"fmt"
)

func largestMultipleOf7(a, b int) interface{} {
	// Находим большее и меньшее значение между a и b
	start := a
	end := b
	if a > b {
		start = b
		end = a
	}

	// Проверяем от большего числа к меньшему
	for x := end; x >= start; x-- {
		if x%7 == 0 {
			return x
		}
	}

	return "no"
}

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	fmt.Println(largestMultipleOf7(a, b))
}
