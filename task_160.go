package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)

	for n > 0 {
		fmt.Println(n % 2)
		n = n / 2
	}
}
