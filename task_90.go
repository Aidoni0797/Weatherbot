package main

import (
	"fmt"
)

func main() {
	var year int
	fmt.Scan(&year)

	if (year%4 == 0 && year%100 != 0) || (year%400 == 0) {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
