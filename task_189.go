package main

import (
	"fmt"
)

func main() {
	var n int
	_, _ = fmt.Scan(&n)

	var min int
	var number int

	_, _ = fmt.Scan(&min)

	for i := 1; i < n; i++ {
		_, _ = fmt.Scan(&number)

		if number < min {
			min = number
		}
	}

	fmt.Println(min)
}
