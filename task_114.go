package main

import (
	"fmt"
	"math"
)

func main() {
	var k, m int
	// Ввод чисел
	fmt.Scan(&k, &m)

	// Вычисляем количество дней, округляя вверх
	days := math.Ceil(float64(k) / float64(m))

	// Выводим результат
	fmt.Println(int(days))
}
