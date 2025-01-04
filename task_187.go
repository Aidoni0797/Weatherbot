package main

import (
	"fmt"
)

func main() {
	var a, b, c, d, e int
	fmt.Scan(&a, &b, &c, &d, &e)

	count := 0
	for x := 0; x <= 1000; x++ {
		if x != e { // Проверяем, чтобы знаменатель не был равен нулю
			numerator := a*x*x*x + b*x*x + c*x + d
			denominator := x - e
			if numerator == 0 && denominator != 0 { // Условие решения уравнения
				count++
			}
		}
	}

	fmt.Println(count)
}
