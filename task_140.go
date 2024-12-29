package main

import "fmt"

func main() {
	var n, s int
	fmt.Scan(&n)
	count := 0
	for i := 1; i <= n; i++ {
		fmt.Scan(&s)
		count = count + s
	}
	fmt.Println(count)
}
