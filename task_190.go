package main

import (
	"fmt"
)

func main() {
	var n int
	_, _ = fmt.Scan(&n)

	var min, max int
	var number int

	_, _ = fmt.Scan(&min)

	for i := 1; i < n; i++ {
		_, _ = fmt.Scan(&number)

		if number < min {
			min = number
		}
		if number > max {
			max = number
		}
	}

	fmt.Println(max - min)
}
