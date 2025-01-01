package main

import (
	"fmt"
)

func main() {
	var a, b int

	fmt.Scan(&a, &b)

	iDONi := 1

	for i := a; i <= b; i++ {
		if i%10 == 7 {
			iDONi = iDONi * i
		}

	}

	fmt.Println(iDONi)
}
