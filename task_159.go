package main

import (
	"fmt"
	"strconv"
)

func main() {
	var n int
	var str, result string
	fmt.Scan(&n)
	for n > 0 { // пока цифры числа не закончатся
		digit := n % 10 // последняя цифра числа
		str = strconv.Itoa(digit)
		result = result + str
		n = n / 10 // избавляемся от последней цифры.
	}
	fmt.Println(result)

}
