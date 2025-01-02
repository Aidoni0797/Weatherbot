package main

import (
	"fmt"
	"strconv"
)

func main() {
	var num, sum int
	var str string
	fmt.Scan(&num)
	for i := 1; i <= num; i++ {
		str = strconv.Itoa(i)
		for j := 0; j < len(str); j++ {
			if str[j] == '7' {
				sum = sum + 1
			}
		}
	}

	fmt.Println(sum)
}
