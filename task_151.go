package main

import (
	"fmt"
)

func main() {
	var n int

	fmt.Scan(&n)

	iDONi := 0

	for n%3 == 0 {
		iDONi = iDONi + 1
		n = n / 3
	}

	fmt.Println(iDONi)
}
