package main

import (
	"fmt"
)

func main() {
	var x int
	fmt.Scan(&x)

	count := 0
	for i := 1; i <= x; i++ {
		if x%i == 0 {
			count++
		}
	}

	fmt.Println(count)
}
