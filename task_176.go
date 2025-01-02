package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)
	for i := n - 1; i > 0; i-- {
		if n%i == 0 {
			fmt.Println(i)
			break
		}

	}
}
