package main

import (
	"fmt"
)

func main() {
	var x float64
	fmt.Scan(&x)
	// Вычисляем дробную часть
	fractionalPart := x - float64(int(x))
	fmt.Println(fractionalPart)
}
