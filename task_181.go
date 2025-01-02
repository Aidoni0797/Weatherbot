package main

import "fmt"

func main() {
	var a, b, c, d int
	fmt.Scan(&a, &b, &c, &d)

	var solutions []int

	// Проходим по всем целым числам от 0 до 1000
	for x := 0; x <= 1000; x++ {
		// Проверяем, является ли x корнем уравнения
		if a*x*x*x+b*x*x+c*x+d == 0 {
			solutions = append(solutions, x)
		}
	}

	// Выводим решения, если они есть
	if len(solutions) > 0 {
		for _, solution := range solutions {
			fmt.Print(solution, " ")
		}
	}
}
