package main

import "fmt"

func main() {
	var n, s int
	fmt.Scan(&n)
	count := 0
	for i := 1; i <= n; i++ {
		fmt.Scan(&s)
		if s == 0 {
			count++
		}
	}
	if count > 0 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}

}
