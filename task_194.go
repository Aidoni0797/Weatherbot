package main

import (
	"fmt"
)

func main() {
	var n int
	_, _ = fmt.Scan(&n)

	var cnt int
	var number int

	for i := 1; i <= n; i++ {
		_, _ = fmt.Scan(&number)

		if number == 0 {
			cnt = cnt + 1
		}

	}

	fmt.Println(cnt)
}
