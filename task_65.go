package main

import (
	"fmt"
	"math"
)

func main() {
	// Ввод координат двух точек
	var x1, y1, x2, y2 int
	fmt.Scan(&x1, &y1)
	fmt.Scan(&x2, &y2)

	// Вычисление манхэттенского расстояния
	manhattanDistance := math.Abs(float64(x2-x1)) + math.Abs(float64(y2-y1))

	// Вывод результата
	fmt.Println(manhattanDistance)
}
