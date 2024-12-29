package main

import "fmt"

func main() {
	var n, s int
	fmt.Scan(&n)
	count := 0
	for i := 1; i <= n; i++ {
		fmt.Scan(&s)
		if s%2 == 0 && s%3 != 0 {
			count = count + s
		}
	}
	fmt.Println(count)
}
