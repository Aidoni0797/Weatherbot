package main

import (
	"fmt"
)

func main() {
	var n, m int
	fmt.Scan(&n)
	m = n
	sum := 0    // сумматор
	for n > 0 { // пока цифры числа не закончатся
		digit := n % 10 // последняя цифра числа
		sum = sum + digit
		n = n / 10 // избавляемся от последней цифры.
	}
	if m%sum == 0 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}

}
