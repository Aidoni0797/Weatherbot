package main

import (
	"fmt"
	"strconv"
)

func main() {
	var n int
	var str, result string
	fmt.Scan(&n)

	for n > 0 {
		str = strconv.Itoa(n % 2)
		result = result + str
		n = n / 2
	}
	fmt.Println(result)
}
