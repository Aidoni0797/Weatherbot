package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)

	if a == b && b == c {
		fmt.Print("YES")
	} else {
		fmt.Print("NO")
	}
}
