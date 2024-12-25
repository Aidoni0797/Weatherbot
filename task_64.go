package main

import (
	"fmt"
	"math"
)

func main() {
	// Ввод координат двух точек
	var x1, y1, x2, y2 int
	fmt.Scan(&x1, &y1, &x2, &y2)

	// Вычисление расстояния между точками
	distance := math.Sqrt(math.Pow(float64(x2-x1), 2) + math.Pow(float64(y2-y1), 2))

	// Вывод результата
	fmt.Println(distance)
}
