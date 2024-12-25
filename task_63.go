package main

import (
	"fmt"
	"math"
)

func main() {
	// Ввод катетов
	var a, b float64
	fmt.Scan(&a, &b)

	// Вычисление гипотенузы
	c := math.Sqrt(a*a + b*b)

	// Вычисление периметра
	perimeter := a + b + c

	// Вывод результата
	fmt.Println(perimeter)
}
