package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)
	sum := 0

	for n != 0 {
		sum = sum + n // сначала прибавляем, а потом вводим!
		fmt.Scan(&n)
	}

	fmt.Println(sum)
}
