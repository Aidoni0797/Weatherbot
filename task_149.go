package main

import (
	"fmt"
)

func main() {
	var n int

	fmt.Scan(&n)

	iDONi := 1

	for i := 1; i <= n; i++ {
		if i%2 == 0 {
			iDONi = iDONi * i
		}

	}

	fmt.Println(iDONi)
}
