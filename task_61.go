package main

import (
	"fmt"
	"math"
)

func main() {
	var n float64 = 8
	fmt.Println(math.Max(n, 4)) // выводит на экран число 8

	fmt.Println(math.Min(n, 4)) // выводит на экран число 4
}
