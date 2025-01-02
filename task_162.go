package main

import (
	"fmt"
	"strconv"
)

func main() {
	var sum, n int
	var str string
	fmt.Scan(&n)

	for n > 0 {
		str = strconv.Itoa(n % 2)
		if str == "1" {
			sum = sum + 1
		}
		n = n / 2
	}
	fmt.Println(sum)
}
