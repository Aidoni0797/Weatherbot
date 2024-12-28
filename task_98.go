package main

import (
	"fmt"
)

func main() {
	var x, b1, b2, d1, d2 int

	// Ввод данных
	fmt.Scan(&x)

	b1 = -3
	b2 = 1

	d1 = 5
	d2 = 9

	// Проверка принадлежности точки x одному из отрезков
	if (x >= b1 && x <= b2) || (x >= d1 && x <= d2) {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
