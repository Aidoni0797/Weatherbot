package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)
	sum := 0    // сумматор
	for n > 0 { // пока цифры числа не закончатся
		digit := n % 10 // последняя цифра числа
		if digit == 4 {
			sum = sum + 1
		}
		n = n / 10 // избавляемся от последней цифры.
	}

	fmt.Println(sum)
}
