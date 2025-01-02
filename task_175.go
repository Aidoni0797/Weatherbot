package main

import (
	"fmt"
)

func main() {
	for n := 1; n <= 10; n++ {
		if n%2 == 0 {
			continue // переход к следующей итерации цикла
		}
		if n == 7 {
			break // прерывание цикла
		}

		fmt.Println(n)
	}
}
