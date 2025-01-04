package main

import (
	"fmt"
)

func main() {

	var cnt int
	var number, result int

	fmt.Scan(&number)
	fmt.Scan(&cnt)

	result = 1

	for i := 1; i <= cnt; i++ {
		result = result * number
	}

	fmt.Println(result)
}
