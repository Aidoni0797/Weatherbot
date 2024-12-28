package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b, c float64
	// Считывание входных данных
	fmt.Scan(&a, &b, &c)

	// Вычисление дискриминанта
	d := b*b - 4*a*c

	// Проверка количества корней в зависимости от дискриминанта
	if d > 0 {
		// Два корня
		x1 := (-b - math.Sqrt(d)) / (2 * a)
		x2 := (-b + math.Sqrt(d)) / (2 * a)
		// Выводим сначала меньший, затем больший корень
		if x1 > x2 {
			x1, x2 = x2, x1
		}
		fmt.Println(x1)
		fmt.Println(x2)
	} else if d == 0 {
		// Один корень
		x := -b / (2 * a)
		fmt.Println(x)
	} else {
		// Нет действительных корней
		// Ничего не выводим
	}
}
