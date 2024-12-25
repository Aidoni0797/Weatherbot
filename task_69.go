package main

import (
	"fmt"
	"math"
)

func main() {
	// Ввод катетов
	var a, b float64
	var result float64
	fmt.Scan(&a, &b)
	result = math.Sqrt((a * a) + (b * b))
	fmt.Println(result)
}
